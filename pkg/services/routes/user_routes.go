package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/globaljoules/energy-usage-simulator/pkg/services/controllers"
	// "github.com/globaljoules/energy-usage-simulator/pkg/services/middleware"
)

type UserRouter struct {
	userController controllers.UserController
}

func NewUserRouter() *UserRouter {
	return &UserRouter{}
}

func (ur *UserRouter) Initialize(rg *gin.RouterGroup) {
	router := rg.Group("user")

	router.POST("/addEntity", ur.userController.AddEntity)
	router.GET("/getEntities", ur.userController.GetEntities)
}
