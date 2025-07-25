package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) repository.UserRepository {
	return &mongoUserRepository{
		collection: db.Collection("users"),
	}
}

func (r *mongoUserRepository) Insert(ctx context.Context, item *entity.User) error {
	if item.ID.String() == "" {
		item.ID = entity.ID(uuid.New().String())
	}
	if r.collection.FindOne(ctx, bson.M{"username": item.Username}).Decode(&entity.User{}) == nil {
		return errors.New("UserRepository.Insert(): Finded user with username " + item.Username)
	}
	res, err := r.collection.InsertOne(ctx, item)
	if err != nil {
		return fmt.Errorf("UserRepository.Insert(): %w", err)
	}
	_, ok := res.InsertedID.(string)
	if !ok {
		return errors.New("UserRepository.Insert(): failed to cast inserted ID to ObjectID")
	}
	return nil
}

func (r *mongoUserRepository) Get(ctx context.Context, username string) (*entity.User, error) {
	var result entity.User
	if err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&result); err != nil {
		return nil, fmt.Errorf("UserRepository.Get()/")
	}
	return &result, nil
}

func (r *mongoUserRepository) Update(ctx context.Context, item *entity.User) error {
	oid, err := primitive.ObjectIDFromHex(item.ID.String())
	if err != nil {
		return err
	}
	_, err = r.collection.ReplaceOne(ctx, bson.M{"_id": oid}, item)
	return err
}

func (r *mongoUserRepository) Delete(ctx context.Context, id entity.ID) error {
	oid, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}
