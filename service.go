package auth_service

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	authv1 "github.com/submaline/auth-service/gen/auth/v1"
)

type AuthService struct{}

func (as *AuthService) LoginWithEmail(
	_ context.Context, request *connect.Request[authv1.LoginWithEmailRequest]) (
	*connect.Response[authv1.LoginWithEmailResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
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
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}
