package repositories

import (
	"context"

	"github.com/GDSC-UIT/egreenbin-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PersonRepository interface {
	GetAll(ctx context.Context) ([]models.Person, error)
	GetByID(ctx context.Context, personID string) (*models.Person, error)
	Create(ctx context.Context, person *models.Person) error
	Update(ctx context.Context, personID string, updates map[string]interface{}) error
	Delete(ctx context.Context, personID string) error
}

type personRepository struct {
	db *mongo.Database
}

func NewPersonRepository(db *mongo.Database) *personRepository {
	return &personRepository{db: db}
}

func (r *personRepository) GetAll(ctx context.Context) ([]models.Person, error) {
	// get all persons in collection persons
	var persons []models.Person
	cursor, err := r.db.Collection("persons").Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	// end find

	if err = cursor.All(context.TODO(), &persons); err != nil {
		panic(err)
	}
	return persons, err

}
func (r *personRepository) GetByID(ctx context.Context, perID string) (*models.Person, error) {
	// Convert the user ID string to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(perID)
	if err != nil {
		return nil, err
	}
	// Find the user with the matching ID in the "persons" collection
	var person models.Person
	err = r.db.Collection("persons").FindOne(ctx, bson.M{"_id": objectID}).Decode(&person)
	if err != nil {
		return nil, err
	}
	return &person, nil
}
func (r *personRepository) Create(ctx context.Context, person *models.Person) error {
	// Insert the person into the "persons" collection
	person.ID = primitive.NewObjectID()
	_, err := r.db.Collection("persons").InsertOne(ctx, person)
	return err
}

func (r *personRepository) Update(ctx context.Context, perID string, updates map[string]interface{}) error {
	// Convert the user ID string to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(perID)
	if err != nil {
		return err
	}
	// Update the user with the matching ID in the "persons" collection
	_, err = r.db.Collection("persons").UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": updates})
	return err
}

func (r *personRepository) Delete(ctx context.Context, perID string) error {
	// Convert the user ID string to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(perID)
	if err != nil {
		return err
	}
	// Delete the user with the matching ID in the "persons" collection
	_, err = r.db.Collection("persons").DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
