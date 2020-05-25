package app

import "github.com/gin-gonic/gin"

var (
	router *gin.Engine
)

func init(){
	router = gin.Default()
}



func StartApp()  {


	HandleMapping()


	if err := router.Run(":4005"); err != nil {
		panic(err)
	}

}
