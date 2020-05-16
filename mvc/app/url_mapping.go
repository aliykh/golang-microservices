package app

import (
	"github.com/aliykh/golang-microservices/mvc/controllers"
)

func HandleMapping() {
	engine.GET("/user/:user_id", controllers.GetUsers)
}
