package routes

import (
	"api/rest"
	sentryGin "github.com/getsentry/sentry-go/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(s *rest.ServerRest) {

	s.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.Engine.Use(sentryGin.New(sentryGin.Options{
		WaitForDelivery: true,
	}))

	route1 := s.Engine.Group("/")
	{
		userGroup := route1.Group("user")
		{
			userGroup.POST("/", s.Controllers.UserController.Create)
			userGroup.GET("/:id", s.Controllers.UserController.GetByID)
			userGroup.PUT("/", s.Controllers.UserController.Update)
			userGroup.DELETE("/:id", s.Controllers.UserController.Delete)
		}
	}
}
