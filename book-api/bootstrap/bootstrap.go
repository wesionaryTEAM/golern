package bootstrap

import (
	"bookapi/api/controllers"
	"bookapi/api/repositories"
	"bookapi/api/routes"
	"bookapi/infrastructure"
	"bookapi/services"
	"context"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	routes.Module,
	infrastructure.Module,
	services.Module,
	repositories.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler infrastructure.Router,
	routes routes.Routes,
	env infrastructure.Env,
	logger infrastructure.Logger,
	database infrastructure.Database,
) {
	conn, _ := database.DB.DB()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("---------------------------")
			logger.Zap.Info("------- BOOKS API 📚  -------")
			logger.Zap.Info("---------------------------")
			conn.SetMaxOpenConns(10)

			go func() {
				routes.Setup()
				if env.ServerPort == "" {
					handler.Gin.Run()
				} else {
					handler.Gin.Run(env.ServerPort)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Zap.Info("Stopping Application")
			conn.Close()
			return nil
		},
	})
}
