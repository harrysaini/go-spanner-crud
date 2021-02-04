package repository

import (
	"context"
	"go-spanner-crud/src/models"
)

// StudentRepository - handle all repo interactions for student
type StudentRepository interface {
	AddNewStudent(ctx context.Context, student models.Student) error
}
