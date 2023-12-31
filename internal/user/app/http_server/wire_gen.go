// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package http_server

import (
	profile2 "github.com/duongnln96/blog-realworld/internal/user/app/http_server/handler/profile"
	user3 "github.com/duongnln96/blog-realworld/internal/user/app/http_server/handler/user"
	"github.com/duongnln96/blog-realworld/internal/user/core/service/profile"
	user2 "github.com/duongnln96/blog-realworld/internal/user/core/service/user"
	"github.com/duongnln96/blog-realworld/internal/user/infras/repo/follow"
	"github.com/duongnln96/blog-realworld/internal/user/infras/repo/user"
	"github.com/duongnln96/blog-realworld/pkg/adapter/postgres"
	"github.com/duongnln96/blog-realworld/pkg/config"
)

// Injectors from wire.go:

func InitNewApp(config2 *config.Configs) *app {
	postgresDBAdapterI := postgresDbAdapter(config2)
	userRepoI := user.NewRepoManager(postgresDBAdapterI)
	userServiceI := user2.NewService(config2, userRepoI)
	handlerI := user3.NewHandler(userServiceI)
	followRepoI := follow.NewRepoManager(postgresDBAdapterI)
	followServiceI := profile.NewService(config2, followRepoI, userRepoI)
	profileHandlerI := profile2.NewHandler(followServiceI, userServiceI)
	http_serverApp := NewApp(config2, handlerI, profileHandlerI)
	return http_serverApp
}

// wire.go:

func postgresDbAdapter(cfg *config.Configs) postgres.PostgresDBAdapterI {
	adapter := postgres.NewPostgresDBAdapter(cfg.PostgresConfigMap.Get("postgres"))

	return adapter
}
