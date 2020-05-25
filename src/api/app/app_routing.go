package app

import "github.com/aliykh/golang-microservices/src/api/controllers"

func HandleMapping() {
	router.POST("/github", controllers.CreateRepo)
}
