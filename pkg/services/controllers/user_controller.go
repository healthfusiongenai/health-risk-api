package controllers

import (
	//	"fmt"
	//	"net/http"
	//	"strings"
	//	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	// "github.com/healthfusiongenai/health-risk-api/pkg/utils"
	// "github.com/healthfusiongenai/health-risk-api/pkg/models"
	// "github.com/healthfusiongenai/health-risk-api/pkg/config"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) AddEntity(ctx *gin.Context) {
	// how do we get the user id?
	// where is the token? can it hold the user id?
	//
}

func (uc *UserController) GetEntities(ctx *gin.Context) {

}
