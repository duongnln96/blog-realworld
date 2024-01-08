package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/duongnln96/blog-realworld/pkg/config"
	"github.com/urfave/cli/v2"
)

func main() {
	configPath, _ := os.Getwd()
	configs := config.LoadConfig(fmt.Sprintf("%s/config/auth", configPath))

	_ = configs

	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name:    "grpc_server",
			Aliases: []string{"grpc"},
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "grpc")
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("Server application running error", "err_info", err.Error())
		return
	}
}
