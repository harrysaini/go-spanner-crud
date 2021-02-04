package services

import (
	"context"
	"go-spanner-crud/src/models"
	"go-spanner-crud/src/models/requests"
	"go-spanner-crud/src/repository"
)

// StudentService - performs all student related actions
type StudentService struct {
	repo repository.StudentRepository
}

// NewStudentService - created student service instance
func NewStudentService(repo repository.StudentRepository) *StudentService {
	return &StudentService{
		repo: repo,
	}
}

// AddNewStudent - adds new student to our database
func (studentService *StudentService) AddNewStudent(ctx context.Context, studentCreateRequest requests.StudentCreateRequest) (models.Student, error) {
	student, err := models.NewStudent(studentCreateRequest)
	if err != nil {
		return models.Student{}, err
	}

	err = studentService.repo.AddNewStudent(ctx, student)
	if err != nil {
		return models.Student{}, err
	}

	return student, nil
}
