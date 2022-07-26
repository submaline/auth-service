package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/bufbuild/connect-go"
	authService "github.com/submaline/auth-service"
	"github.com/submaline/auth-service/gen/auth/v1/authv1connect"
	"github.com/submaline/auth-service/gen/supervisor/v1/supervisorv1connect"
	"github.com/submaline/interceptors"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
)

func main() {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("failed to setup firebase app: %v", err)
	}
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("failed to setup firebase auth: %v", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to setup logger: %v", err)
	}
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	supervisorServiceClient := supervisorv1connect.NewSupervisorServiceClient(
		http.DefaultClient,
		os.Getenv("SUPERVISOR_SERVICE_URL"))

	authServiceHandler := &authService.AuthService{
		AuthClient:       authClient,
		SupervisorClient: &supervisorServiceClient,
		Logger:           logger,
	}

	mux := http.NewServeMux()
	inters := connect.WithInterceptors(
		interceptors.NewFirebaseAuthInterceptor(authClient, interceptors.AuthPolicy{"/auth.v1.AuthService/LoginWithEmail": false}),
		interceptors.NewLoggingInterceptor(logger))
	mux.Handle(authv1connect.NewAuthServiceHandler(
		authServiceHandler,
		inters))

	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	addr := fmt.Sprintf(":%s", port)

	// h2cはインセキュア通信でhttp2を実現するための実装
	log.Printf("Service listening on %v", port)
	if err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
