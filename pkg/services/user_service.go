package services

import (
	//	"log"
	//	"fmt"
	//	"strings"
	//	"net/http"

	//	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	//	"github.com/globaljoules/energy-usage-simulator/pkg/config"
	"github.com/globaljoules/energy-usage-simulator/pkg/initializer"
	// "github.com/globaljoules/energy-usage-simulator/pkg/utils"
)

type UserService struct {
	Server *gin.Engine
}

func NewUserService() *UserService {
	userService := &UserService{
		Server: gin.Default(),
	}

	return userService
}

// Initialize the service based on DB choice: Postgres, Mongo, Cassandra
func (us *UserService) Initialize(rg *gin.RouterGroup, initializer *initializer.Initializer) error {

	// put in a switch?

	//	config, err :=

	return nil
}
