package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello world"))
	})


	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		panic(err)
	}


}