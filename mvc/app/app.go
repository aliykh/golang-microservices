package app

import (
	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func init() {
	engine = gin.Default()
}

func StartApp(){

	HandleMapping()

	err := engine.Run(":8083")
	if err != nil {
		panic(err)
	}
}
