package app

import (
	"github.com/aliykh/golang-microservices/mvc/controllers"
	"net/http"
)

func StartApp(){
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("it is working fine"))
	})
	http.HandleFunc("/user", controllers.GetUsers)

	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		panic(err)
	}
}
