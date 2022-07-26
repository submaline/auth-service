package auth_service

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/bufbuild/connect-go"
	authv1 "github.com/submaline/auth-service/gen/auth/v1"
	supervisorv1 "github.com/submaline/auth-service/gen/supervisor/v1"
	"github.com/submaline/auth-service/gen/supervisor/v1/supervisorv1connect"
	typesv1 "github.com/submaline/auth-service/gen/types/v1"
	"go.uber.org/zap"
	"os"
	"time"
)

type AuthService struct {
	AuthClient       *auth.Client
	SupervisorClient *supervisorv1connect.SupervisorServiceClient
	Logger           *zap.Logger
}

func (as *AuthService) LoginWithEmail(
	_ context.Context, request *connect.Request[authv1.LoginWithEmailRequest]) (
	*connect.Response[authv1.LoginWithEmailResponse], error) {

	// ユーザートークン生成
	token, err := GenerateToken(request.Msg.Email, request.Msg.Password, false)
	if err != nil {
		// todo : エラーをしっかり作る
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// トークンの期限(秒数)を取得
	expiresIn := token.ExpiresAt.Sub(time.Now()).Seconds()

	// 管理者用トークンを生成
	adminToken, err := GenerateToken(os.Getenv("ADMIN_FB_EMAIL"), os.Getenv("ADMIN_FB_PASSWORD"), false)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}
	// supervisorに報告する内容
	recordReq := connect.NewRequest(&supervisorv1.RecordOperationRequest{Operations: []*typesv1.Operation{
		{
			Type:        typesv1.OperationType_OPERATION_TYPE_LOGIN_WITH_EMAIL,
			Source:      token.UID,
			Destination: []string{token.UID},
		},
	}})
	// リクエストに管理者用トークンをつけてあげる
	recordReq.Header().Set("Authorization", fmt.Sprintf("Bearer %s", adminToken.IdToken))
	go func() {
		_, err = (*as.SupervisorClient).RecordOperation(context.Background(), recordReq)
		as.Logger.Error("failed to record op", zap.Error(err))
	}()

	// ユーザーへの返却内容
	resp := connect.NewResponse(&authv1.LoginWithEmailResponse{
		AuthToken: &typesv1.AuthToken{
			Token:        token.IdToken,
			ExpiresIn:    int64(expiresIn),
			RefreshToken: token.Refresh,
		},
	})

	return resp, nil
}

func (as *AuthService) UpdatePassword(
	_ context.Context, request *connect.Request[authv1.UpdatePasswordRequest]) (
	*connect.Response[authv1.UpdatePasswordResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}

func (as *AuthService) TokenRefresh(
	_ context.Context, request *connect.Request[authv1.TokenRefreshRequest]) (
	*connect.Response[authv1.TokenRefreshResponse], error) {

	tokenData, err := GenTokenWithRefresh(request.Msg.RefreshToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	expiresIn := tokenData.ExpiresAt.Sub(time.Now()).Seconds()

	// 管理者用トークンを生成
	adminToken, err := GenerateToken(os.Getenv("ADMIN_FB_EMAIL"), os.Getenv("ADMIN_FB_PASSWORD"), false)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}
	// supervisorに報告する内容
	recordReq := connect.NewRequest(&supervisorv1.RecordOperationRequest{Operations: []*typesv1.Operation{
		{
			Type:        typesv1.OperationType_OPERATION_TYPE_TOKEN_REFRESH,
			Source:      tokenData.UID,
			Destination: []string{tokenData.UID},
		},
	}})
	// リクエストに管理者用トークンをつけてあげる
	recordReq.Header().Set("Authorization", fmt.Sprintf("Bearer %s", adminToken.IdToken))
	go func() {
		_, err = (*as.SupervisorClient).RecordOperation(context.Background(), recordReq)
		as.Logger.Error("failed to record op", zap.Error(err))
	}()

	return connect.NewResponse(&authv1.TokenRefreshResponse{AuthToken: &typesv1.AuthToken{
		Token:        tokenData.IdToken,
		ExpiresIn:    int64(expiresIn),
		RefreshToken: tokenData.Refresh,
	}}), nil
}
