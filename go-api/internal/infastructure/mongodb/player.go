package mongodb

import (
	"context"
	"errors"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoPCRepository struct {
	collection *mongo.Collection
}

func NewPlayerCharacterRepository(db *mongo.Database) repository.PlayerCharacterRepository {
	return &mongoPCRepository{
		collection: db.Collection("players"),
	}
}

func (r *mongoPCRepository) Insert(ctx context.Context, item *entity.PlayerCharacter) (*string, error) {
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

func (r *mongoPCRepository) Get(ctx context.Context, creator_id string, ids ...string) ([]*entity.PlayerCharacter, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"creator_id": creator_id, "_id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}
	var results []*entity.PlayerCharacter
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoPCRepository) Update(ctx context.Context, item *entity.PlayerCharacter) error {
	oid, err := primitive.ObjectIDFromHex(item.Character.ID)
	if err != nil {
		return err
	}
	_, err = r.collection.ReplaceOne(ctx, bson.M{"_id": oid}, item)
	return err
}

func (r *mongoPCRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}
