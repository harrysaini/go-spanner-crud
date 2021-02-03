package spannerRepo

import (
	"go-spanner-crud/src/models"

	"cloud.google.com/go/spanner"
)

type StudentSpannerRepository struct {
	client *spanner.Client
}

func New(client *spanner.Client) *StudentSpannerRepository {
	return &StudentSpannerRepository{
		client: client,
	}
}

func (repo *StudentSpannerRepository) AddNewStudent(student models.Student) (models.Student, error) {

}
