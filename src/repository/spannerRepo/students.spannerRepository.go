package spannerrepo

import (
	"context"
	"errors"
	"go-spanner-crud/src/models"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
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
		spanner.Insert(
			"Students",
			studentColumns,
			[]interface{}{student.UUID, student.RollNumber, student.FirstName, student.LastName, student.BirthDate, student.Branch},
		),
	}
	_, err := repo.client.Apply(ctx, m)

	if spanner.ErrCode(err) == codes.AlreadyExists {
		return errors.New("Duplicate user")
	}

	return err
}

// GetStudent - get student by
func (repo *StudentSpannerRepository) GetStudent(ctx context.Context, uuid string) (models.Student, error) {

	stmt := spanner.Statement{
		SQL: `SELECT UUID, RollNumber, FirstName, LastName, BirthDate,  Branch FROM Students
			WHERE UUID = @uuid`,
		Params: map[string]interface{}{
			"uuid": uuid,
		},
	}
	iter := repo.client.Single().Query(ctx, stmt)
	defer iter.Stop()

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return models.Student{}, errors.New("User Not Found")
		}
		if err != nil {
			return models.Student{}, err
		}

		var student models.Student
		var birthDate spanner.NullDate

		if err := row.Columns(&student.UUID, &student.RollNumber, &student.FirstName, &student.LastName, &birthDate, &student.Branch); err != nil {
			return models.Student{}, err
		}

		student.BirthDate = birthDate.Date.String()

		return student, nil
	}
}

// GetAllStudents - get all students
func (repo *StudentSpannerRepository) GetAllStudents(ctx context.Context, limit int64, offset int64) ([]models.Student, error) {

	stmt := spanner.Statement{
		SQL: `SELECT UUID, RollNumber, FirstName, LastName, BirthDate,  Branch FROM Students
			LIMIT @limit OFFSET @offset`,
		Params: map[string]interface{}{
			"limit":  limit,
			"offset": offset,
		},
	}
	iter := repo.client.Single().Query(ctx, stmt)
	defer iter.Stop()

	students := []models.Student{}

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return students, nil
		}
		if err != nil {
			return students, err
		}

		var student models.Student
		var birthDate spanner.NullDate

		if err := row.Columns(&student.UUID, &student.RollNumber, &student.FirstName, &student.LastName, &birthDate, &student.Branch); err != nil {
			return students, err
		}

		student.BirthDate = birthDate.Date.String()

		students = append(students, student)

	}
}

// UpdateStudent - update student in database
func (repo *StudentSpannerRepository) UpdateStudent(ctx context.Context, student models.Student) error {
	studentColumns := []string{"UUID", "FirstName", "LastName", "BirthDate"}

	m := []*spanner.Mutation{
		spanner.Update(
			"Students",
			studentColumns,
			[]interface{}{student.UUID, student.FirstName, student.LastName, student.BirthDate},
		),
	}
	_, err := repo.client.Apply(ctx, m)

	return err
}

// DeleteStudent - delete student in database
func (repo *StudentSpannerRepository) DeleteStudent(ctx context.Context, uuid string) error {
	m := []*spanner.Mutation{
		spanner.Delete("Students", spanner.Key{uuid}),
	}
	_, err := repo.client.Apply(ctx, m)
	return err
}
