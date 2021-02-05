package main

import (
	"fmt"
	"go-spanner-crud/src/cache"
	"go-spanner-crud/src/controllers"
	"go-spanner-crud/src/libs"
	spannerrepo "go-spanner-crud/src/repository/spannerRepo"
	"go-spanner-crud/src/services"
	"log"
	"net/http"

	"cloud.google.com/go/spanner"
	"github.com/julienschmidt/httprouter"
)

func setUpStudentsRoutes(dbClient *spanner.Client, router *httprouter.Router, cacheClient *libs.RedisCache) {
	studentsRepo := spannerrepo.NewStudentSpannerRepository(dbClient)
	studentCache := cache.NewStudentCache(cacheClient)
	studentService := services.NewStudentService(studentsRepo, studentCache)
	studentController := controllers.NewStudentsController(studentService)

	router.GET("/api/student/", studentController.HandleGetAllStudents)
	router.POST("/api/student/", studentController.HandleAddNewStudent)
	router.GET("/api/student/:uuid/", studentController.HandleGetStudent)
	router.PUT("/api/student/:uuid/", studentController.HandleUpdateStudent)
	router.DELETE("/api/student/:uuid/", studentController.HandleDeleteStudent)

}

func main() {

	var config = libs.Conf
	dbClient := libs.GetSpannerClientInstance()
	cache := libs.NewRedisCacheClient()

	router := httprouter.New()
	router.GET("/", controllers.Index)

	setUpStudentsRoutes(dbClient, router, cache)

	address := fmt.Sprintf(":%d", config.Server.Port)
	log.Println("Will listen on -> ", address)
	log.Fatalln(http.ListenAndServe(address, router))
}
