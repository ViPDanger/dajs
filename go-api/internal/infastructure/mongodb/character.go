package mongodb

import (
	"context"
	"errors"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoCharacterRepository struct {
	collection *mongo.Collection
}

func NewCharacterRepository(db *mongo.Database) repository.CharacterRepository {
	return &mongoCharacterRepository{
		collection: db.Collection("characters"),
	}
}

func (r *mongoCharacterRepository) Insert(ctx context.Context, item *entity.Character) (*string, error) {
	if item.ID == "" {
		item.ID = uuid.New().String()
	}
	res, err := r.collection.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}
	oid, ok := res.InsertedID.(string)
	if !ok {
		return nil, errors.New("failed to cast inserted ID to ObjectID")
	}
	id := string(oid)
	return &id, nil
}

func (r *mongoCharacterRepository) Get(ctx context.Context, creator_id string, ids ...string) ([]*entity.Character, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"creator_id": creator_id, "_id": bson.M{"$in": ids}})
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}
	var results []*entity.Character

	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoCharacterRepository) Update(ctx context.Context, item *entity.Character) error {
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": item.ID}, item)
	return err
}

func (r *mongoCharacterRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
