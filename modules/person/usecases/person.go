package usecases

import (
	"context"
	"fmt"

	"github.com/GDSC-UIT/egreenbin-api/models"
	"github.com/GDSC-UIT/egreenbin-api/modules/person/repositories"
)

type PersonUsecase interface {
	GetAll(ctx context.Context) ([]models.Person, error)
	GetByID(ctx context.Context, personID string) (*models.Person, error)
	Create(ctx context.Context, person *models.Person) error
	Update(ctx context.Context, personID string, updates map[string]interface{}) error
	Delete(ctx context.Context, personID string) error
}

type personUsecase struct {
	personRepo repositories.PersonRepository
}

func NewPersonUsecase(personRepo repositories.PersonRepository) *personUsecase {
	return &personUsecase{personRepo: personRepo}
}

func (u *personUsecase) GetAll(ctx context.Context) ([]models.Person, error) {
	return u.personRepo.GetAll(ctx)
}

func (u *personUsecase) GetByID(ctx context.Context, personID string) (*models.Person, error) {
	// Validate the Person ID
	// if !bson.Ob(personID) {
	// 	return nil, fmt.Errorf("invalid Person ID")
	// }
	// Get the Person from the database
	return u.personRepo.GetByID(ctx, personID)
}

// ... implement the other methods of the PersonUsecase interface ...
func (u *personUsecase) Create(ctx context.Context, person *models.Person) error {
	// Validate the Person input
	if person.Name == "" || person.Genre < 0 {
		return fmt.Errorf("invalid Person input")
	}

	// Create the Person in the database
	return u.personRepo.Create(ctx, person)
}

// ... implement the other methods of the PersonUsecase interface ...
func (u *personUsecase) Update(ctx context.Context, personID string, updates map[string]interface{}) error {
	return u.personRepo.Update(ctx, personID, updates)
}

// ... implement the other methods of the PersonUsecase interface ...
func (u *personUsecase) Delete(ctx context.Context, personID string) error {
	return u.personRepo.Delete(ctx, personID)
}
