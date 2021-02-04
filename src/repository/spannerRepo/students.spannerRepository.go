package spannerrepo

import (
	"context"
	"go-spanner-crud/src/models"

	"cloud.google.com/go/spanner"
)

// StudentSpannerRepository - functions for creating data in spanner
type StudentSpannerRepository struct {
	client *spanner.Client
}

// NewStudentSpannerRepository - Create new spanner repo object
func NewStudentSpannerRepository(client *spanner.Client) *StudentSpannerRepository {
	return &StudentSpannerRepository{
		client: client,
	}
}

// AddNewStudent - add new student to db
func (repo *StudentSpannerRepository) AddNewStudent(ctx context.Context, student models.Student) error {
	studentColumns := []string{"UUID", "RollNumber", "FirstName", "LastName", "BirthDate", "Branch"}

	m := []*spanner.Mutation{
		spanner.InsertOrUpdate(
			"Students",
			studentColumns,
			[]interface{}{student.UUID, student.RollNumber, student.FirstName, student.LastName, student.BirthDate, student.Branch},
		),
	}
	_, err := repo.client.Apply(ctx, m)

	return err
}
