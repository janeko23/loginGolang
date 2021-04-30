package app

import (
	"log"
	"login-go/app/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (a *App) Initialize() {
	a.Router = gin.Default()
}

// Run the app on it's router
func (a *App) Run(port string) {
	log.Fatal(a.Router.Run(port))
}
// Set all required routers
func (a *App) SetRouters() {
	a.Router.GET("/", handler.Simple)
	a.Router.POST("/login", handler.Login)

}
