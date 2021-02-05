package models

import (
	"fmt"
	"go-spanner-crud/src/models/requests"

	"github.com/google/uuid"
)

// Student - Represent student db object
type Student struct {
	UUID       string `json:"uuid"`
	RollNumber int64  `json:"rollNumber"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	BirthDate  string `json:"birthDate"`
	Branch     string `json:"branch"`
}

// NewStudent - creates new student obj
func NewStudent(studentReq requests.StudentCreateRequest) Student {

	id := fmt.Sprintf("%d-%s", studentReq.RollNumber, studentReq.Branch)
	uuid := uuid.NewSHA1(uuid.NameSpaceURL, []byte(id))

	return Student{
		UUID:       uuid.String(),
		RollNumber: studentReq.RollNumber,
		FirstName:  studentReq.FirstName,
		LastName:   studentReq.LastName,
		BirthDate:  studentReq.BirthDate,
		Branch:     studentReq.Branch,
	}
}
