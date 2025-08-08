package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoPlayerCharRepository struct {
	collection *mongo.Collection
}

func NewPlayerCharRepository(db *mongo.Database) repository.PlayerCharRepository {
	return &mongoPlayerCharRepository{
		collection: db.Collection("playerChars"),
	}
}

func (r *mongoPlayerCharRepository) Insert(ctx context.Context, item *entity.PlayerChar) (*string, error) {
	item.Character.ID = uuid.New().String()
	res, err := r.collection.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}
	id, ok := res.InsertedID.(string)
	if !ok {
		return nil, errors.New("failed to cast inserted ID to ObjectID")
	}
	return &id, nil
}
func (r *mongoPlayerCharRepository) Get(ctx context.Context, creator_id string, ids ...string) ([]*entity.PlayerChar, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"creator_id": creator_id, "_id": bson.M{"$in": ids}})
	defer cursor.Close(ctx)
	if err != nil {
		return nil, fmt.Errorf("CharcterRepository.GetArray():  %w", err)
	}
	var results []*entity.PlayerChar

	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}
func (r *mongoPlayerCharRepository) Update(ctx context.Context, item *entity.PlayerChar) error {
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": item.Character.ID}, item)
	return err
}

func (r *mongoPlayerCharRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
