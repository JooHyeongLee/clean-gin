package routes

import (
	"github.com/JooHyeongLee/clean-gin/api/controllers"
	"github.com/JooHyeongLee/clean-gin/lib"
)

// MemberRoutes struct
type MemberRoutes struct {
	logger           lib.Logger
	handler          lib.RequestHandler
	memberController controllers.MemberController
}

// Setup member routes
func (s MemberRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/member/:id", s.memberController.GetOneMember)
	}
}
