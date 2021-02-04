package controllers

import (
	"encoding/json"
	"go-spanner-crud/src/models/requests"
	"go-spanner-crud/src/services"
	"go-spanner-crud/src/utils"
	"net/http"

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
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	ctx := req.Context()

	student, err := controler.service.AddNewStudent(ctx, studentCreateRequest)
	if err != nil {
		utils.SendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	utils.SendResponse(w, student)
}
