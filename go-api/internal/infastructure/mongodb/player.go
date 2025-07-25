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

func (r *mongoPCRepository) GetByCreatorID(ctx context.Context, id entity.ID) ([]*entity.PlayerCharacter, error) {
	oid, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return nil, err
	}
	cursor, err := r.collection.Find(ctx, bson.M{"_id": bson.M{"$in": oid}})
	if err != nil {
		return nil, err
	}
	var results []*entity.PlayerCharacter
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoPCRepository) Insert(ctx context.Context, item *entity.PlayerCharacter) (*entity.ID, error) {
	res, err := r.collection.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}
	oid, ok := res.InsertedID.(string)
	if !ok {
		return nil, errors.New("failed to cast inserted ID to ObjectID")
	}
	id := entity.ID(oid)
	return &id, nil
}

func (r *mongoPCRepository) GetByID(ctx context.Context, id entity.ID) (*entity.PlayerCharacter, error) {
	oid, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return nil, err
	}
	var result entity.PlayerCharacter
	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *mongoPCRepository) GetArray(ctx context.Context, ids []entity.ID) ([]*entity.PlayerCharacter, error) {
	objectIDs := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		oid, err := primitive.ObjectIDFromHex(string(id))
		if err != nil {
			return nil, err
		}
		objectIDs[i] = oid
	}
	cursor, err := r.collection.Find(ctx, bson.M{"_id": bson.M{"$in": objectIDs}})
	if err != nil {
		return nil, err
	}
	var results []*entity.PlayerCharacter
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoPCRepository) GetAll(ctx context.Context) ([]*entity.PlayerCharacter, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
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
	oid, err := primitive.ObjectIDFromHex(item.Character.ID.String())
	if err != nil {
		return err
	}
	_, err = r.collection.ReplaceOne(ctx, bson.M{"_id": oid}, item)
	return err
}

func (r *mongoPCRepository) Delete(ctx context.Context, id entity.ID) error {
	oid, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}
