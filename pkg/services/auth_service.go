package services

import (
	"github.com/gin-gonic/gin"

	"github.com/healthfusiongenai/health-risk-api/pkg/initializer"
	"github.com/healthfusiongenai/health-risk-api/pkg/services/controllers"
	"github.com/healthfusiongenai/health-risk-api/pkg/services/middleware"
	"github.com/healthfusiongenai/health-risk-api/pkg/services/routes"
)

type AuthService struct {
	AuthController controllers.AuthController
	AuthRouter     routes.AuthRouter
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) Initialize(rg *gin.RouterGroup, initializer *initializer.Initializer) {
	as.AuthController = controllers.NewAuthController(initializer.DB)
	userMiddleware := middleware.NewUserMiddleware(initializer)
	as.AuthRouter = routes.NewAuthRouter(as.AuthController, *userMiddleware)
	// connect the routes to the url path auth/*
	as.AuthRouter.Initialize(rg)
}
