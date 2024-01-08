// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package grpc_server

import (
	"github.com/duongnln96/blog-realworld/internal/auth/adapter/repo/syclladb/auth_token"
	auth_token3 "github.com/duongnln96/blog-realworld/internal/auth/app/grpc_server/handler/auth_token"
	auth_token2 "github.com/duongnln96/blog-realworld/internal/auth/usecases/auth_token"
	"github.com/duongnln96/blog-realworld/internal/pkg/token"
	"github.com/duongnln96/blog-realworld/pkg/adapter/scylladb"
	"github.com/duongnln96/blog-realworld/pkg/config"
	"log"
	"log/slog"
)

// Injectors from wire.go:

func InitNewApp(config2 *config.Configs, logger *slog.Logger) (*app, func()) {
	scyllaDBAdaterI, cleanup := scylladbAdapter(config2)
	authTokenRepoI := auth_token.NewRepoManager(scyllaDBAdaterI)
	tokenMakerI := jwtTokenAdapter(config2)
	authTokenUseCasesI := auth_token2.NewUsecases(authTokenRepoI, tokenMakerI)
	authTokenServiceServer := auth_token3.NewHandler(authTokenUseCasesI)
	grpc_serverApp := NewApp(config2, logger, authTokenServiceServer)
	return grpc_serverApp, func() {
		cleanup()
	}
}

// wire.go:

func scylladbAdapter(cfg *config.Configs) (scylladb.ScyllaDBAdaterI, func()) {
	adapter := scylladb.NewScyllaDBAdapter(cfg.ScyllaDBConfigMap.Get("scylladb"))

	return adapter, func() { adapter.Close() }
}

func jwtTokenAdapter(cfg *config.Configs) token.TokenMakerI {
	secret, ok := cfg.Other.Get("jwt_secret_key").(string)
	if !ok {
		log.Panic("Cannot get jwt_secret_key")
	}

	tokenMaker, err := token.NewJWTTokenMaker(secret)
	if err != nil {
		log.Panic("Cannot create new instance for token maker")
	}

	return tokenMaker
}