package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"go-spanner-crud/src/models/requests"
	"go-spanner-crud/src/services"
	"go-spanner-crud/src/utils"

	"github.com/julienschmidt/httprouter"
)

// StudentsController - handle all student requests
type StudentsController struct {
	service *services.StudentService
}

// NewStudentsController - creates new controller
func NewStudentsController(service *services.StudentService) *StudentsController {
	return &StudentsController{
		service: service,
	}
}

// HandleAddNewStudent - handle add request
func (controler *StudentsController) HandleAddNewStudent(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

	var studentCreateRequest requests.StudentCreateRequest

	err := json.NewDecoder(req.Body).Decode(&studentCreateRequest)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	err = studentCreateRequest.Validate()
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	ctx := req.Context()

	student, err := controler.service.AddNewStudent(ctx, studentCreateRequest)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, student)
}

// HandleGetStudent - get student by uuid
func (controler *StudentsController) HandleGetStudent(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

	uuid := params.ByName("uuid")

	if uuid == "" {
		utils.SendErrorResponse(w, errors.New("uuid missing"), http.StatusBadRequest)
		return
	}
	log.Println("HandleGetStudent", uuid)

	student, err := controler.service.GetStudent(req.Context(), uuid)

	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, student)

}

// HandleGetAllStudents - get all students
func (controler *StudentsController) HandleGetAllStudents(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

	limitParam := req.FormValue("limit")
	offsetParam := req.FormValue("offset")

	var err error

	var limit int64 = 10
	if limitParam != "" {
		limit, err = strconv.ParseInt(limitParam, 10, 64)
	}
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	var offset int64 = 0
	if offsetParam != "" {
		offset, err = strconv.ParseInt(offsetParam, 10, 64)
	}
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	students, err := controler.service.GetAllStudents(req.Context(), limit, offset)

	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, students)

}

// HandleUpdateStudent - handle update student req
func (controler *StudentsController) HandleUpdateStudent(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	uuid := params.ByName("uuid")

	if uuid == "" {
		utils.SendErrorResponse(w, errors.New("uuid missing"), http.StatusBadRequest)
		return
	}

	var studentUpdateRequest requests.StudentUpdateRequest
	err := json.NewDecoder(req.Body).Decode(&studentUpdateRequest)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	err = studentUpdateRequest.Validate()
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	student, err := controler.service.UpdateStudent(req.Context(), uuid, studentUpdateRequest)

	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, student)

}

// HandleDeleteStudent - delete student by uuid
func (controler *StudentsController) HandleDeleteStudent(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	uuid := params.ByName("uuid")

	if uuid == "" {
		utils.SendErrorResponse(w, errors.New("uuid missing"), http.StatusBadRequest)
		return
	}

	err := controler.service.DeleteStudent(req.Context(), uuid)

	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, nil)

}
