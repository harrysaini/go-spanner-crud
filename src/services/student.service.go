package services

import "go-spanner-crud/src/repository"

type StudentService struct {
	repo *repository.StudentRepository
}

func New(repo *repository.StudentRepository) *StudentService {
	return &StudentService{
		repo: repo,
	}
}
