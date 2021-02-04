package models

import (
	"go-spanner-crud/src/models/requests"

	"github.com/google/uuid"
)

// Student - Represent student db object
type Student struct {
	UUID       string
	RollNumber int
	FirstName  string
	LastName   string
	BirthDate  string
	Branch     string
}

// NewStudent - creates new student obj
func NewStudent(studentReq requests.StudentCreateRequest) (Student, error) {

	uuid, err := uuid.NewRandom()
	if err != nil {
		return Student{}, err
	}

	return Student{
		UUID:       uuid.String(),
		RollNumber: studentReq.RollNumber,
		FirstName:  studentReq.FirstName,
		LastName:   studentReq.LastName,
		BirthDate:  studentReq.BirthDate,
		Branch:     studentReq.Branch,
	}, nil
}
