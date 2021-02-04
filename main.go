package main

import (
	"fmt"
	"go-spanner-crud/src/controllers"
	"go-spanner-crud/src/libs"
	spannerrepo "go-spanner-crud/src/repository/spannerRepo"
	"go-spanner-crud/src/services"
	"log"
	"net/http"

	"cloud.google.com/go/spanner"
	"github.com/julienschmidt/httprouter"
)

func setUpStudentsRoutes(dbClient *spanner.Client) *httprouter.Router {
	studentsRepo := spannerrepo.NewStudentSpannerRepository(dbClient)
	studentService := services.NewStudentService(studentsRepo)
	studentController := controllers.NewStudentsController(studentService)
	studentRouter := httprouter.New()

	studentRouter.POST("/", studentController.HandleAddNewStudent)

	return studentRouter

}

func main() {

	var config = libs.Conf
	fmt.Println(config)
	dbClient := libs.GetSpannerClientInstance()

	studentRouter := setUpStudentsRoutes(dbClient)

	router := httprouter.New()
	router.GET("/", controllers.Index)
	http.Handle("/api/student", studentRouter)

	address := fmt.Sprintf(":%d", config.Server.Port)
	log.Println("Will listen on -> ", address)
	log.Fatalln(http.ListenAndServe(address, router))
}
