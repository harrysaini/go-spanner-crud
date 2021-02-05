package services

import (
	"context"
	"go-spanner-crud/src/cache"
	"go-spanner-crud/src/models"
	"go-spanner-crud/src/models/requests"
	"go-spanner-crud/src/repository"
	"log"
)

// StudentService - performs all student related actions
type StudentService struct {
	repo  repository.StudentRepository
	cache *cache.StudentCache
}

// NewStudentService - created student service instance
func NewStudentService(repo repository.StudentRepository, cache *cache.StudentCache) *StudentService {
	return &StudentService{
		repo:  repo,
		cache: cache,
	}
}

// AddNewStudent - adds new student to our database
func (studentService *StudentService) AddNewStudent(ctx context.Context, studentCreateRequest requests.StudentCreateRequest) (models.Student, error) {
	student := models.NewStudent(studentCreateRequest)

	err := studentService.repo.AddNewStudent(ctx, student)
	if err != nil {
		return models.Student{}, err
	}

	err = studentService.cache.AddNew(ctx, student)
	if err != nil {
		return student, err
	}

	return student, nil
}

// GetStudent get student by uuid
func (studentService *StudentService) GetStudent(ctx context.Context, uuid string) (models.Student, error) {
	log.Println("StudentService", "GetStudent", uuid)

	studentPtr, err := studentService.cache.Get(ctx, uuid)
	if err != nil {
		return models.Student{}, err
	}

	if studentPtr == nil {

		log.Println("StudentService", "GetStudent", "Not Found user in cache")

		student, err := studentService.repo.GetStudent(ctx, uuid)
		if err != nil {
			return models.Student{}, err
		}

		err = studentService.cache.AddNew(ctx, student)
		if err != nil {
			return student, err
		}

		return student, nil
	}

	return *studentPtr, err
}

// GetAllStudents get all students
func (studentService *StudentService) GetAllStudents(ctx context.Context, limit int64, offset int64) ([]models.Student, error) {
	return studentService.repo.GetAllStudents(ctx, limit, offset)
}

// UpdateStudent - updates student
func (studentService *StudentService) UpdateStudent(ctx context.Context, uuid string, studentUpdateRequest requests.StudentUpdateRequest) (models.Student, error) {
	student, err := studentService.GetStudent(ctx, uuid)
	if err != nil {
		return models.Student{}, err
	}

	// Update these fields only
	student.FirstName = studentUpdateRequest.FirstName
	student.LastName = studentUpdateRequest.LastName
	student.BirthDate = studentUpdateRequest.BirthDate

	err = studentService.repo.UpdateStudent(ctx, student)
	if err != nil {
		return models.Student{}, err
	}

	err = studentService.cache.AddNew(ctx, student)
	if err != nil {
		return student, err
	}

	return student, nil
}

// DeleteStudent - delete student
func (studentService *StudentService) DeleteStudent(ctx context.Context, uuid string) error {
	_, err := studentService.GetStudent(ctx, uuid)
	if err != nil {
		return err
	}

	err = studentService.repo.DeleteStudent(ctx, uuid)
	if err != nil {
		return err
	}

	err = studentService.cache.Purge(ctx, uuid)
	if err != nil {
		return err
	}

	return nil

}
