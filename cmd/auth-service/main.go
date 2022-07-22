package main

import (
	"fmt"
	authService "github.com/submaline/auth-service"
	"github.com/submaline/auth-service/gen/auth/v1/authv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
)

func main() {
	authServiceHandler := &authService.AuthService{}

	mux := http.NewServeMux()
	mux.Handle(authv1connect.NewAuthServiceHandler(
		authServiceHandler))

	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	addr := fmt.Sprintf(":%s", port)

	// h2cはインセキュア通信でhttp2を実現するための実装
	log.Printf("Service listenging on %v", port)
	if err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
