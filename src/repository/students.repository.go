package repository

import "go-spanner-crud/src/models"

type StudentRepository interface {
	AddStudent(student models.Student) (models.Student, error)
}
