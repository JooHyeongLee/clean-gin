package bootstrap

import (
	"github.com/JooHyeongLee/clean-gin/api/controllers"
	"github.com/JooHyeongLee/clean-gin/api/middlewares"
	"github.com/JooHyeongLee/clean-gin/api/routes"
	"github.com/JooHyeongLee/clean-gin/lib"
	"github.com/JooHyeongLee/clean-gin/repository"
	"github.com/JooHyeongLee/clean-gin/services"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	repository.Module,
)
