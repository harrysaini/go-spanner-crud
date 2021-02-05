package repository

import (
	"context"
	"go-spanner-crud/src/models"
)

// StudentRepository - handle all repo interactions for student
type StudentRepository interface {
	AddNewStudent(ctx context.Context, student models.Student) error
	GetStudent(ctx context.Context, uuid string) (models.Student, error)
	GetAllStudents(ctx context.Context, limit int64, offset int64) ([]models.Student, error)
	UpdateStudent(ctx context.Context, student models.Student) error
	DeleteStudent(ctx context.Context, uuid string) error
}
