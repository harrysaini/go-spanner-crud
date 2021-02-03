package main

import (
	"fmt"
	"go-spanner-crud/src/controllers"
	"go-spanner-crud/src/libs"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	var config = libs.Conf

	fmt.Println(config)
	router := httprouter.New()

	router.GET("/", controllers.Index)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), router))
}
