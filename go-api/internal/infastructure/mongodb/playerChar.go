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

type mongoPlayerCharRepository struct {
	collection *mongo.Collection
}

func NewPlayerCharRepository(db *mongo.Database) repository.PlayerCharRepository {
	return &mongoPlayerCharRepository{
		collection: db.Collection("playerChars"),
	}
}

func (r *mongoPlayerCharRepository) Insert(ctx context.Context, item *entity.PlayerChar) (*entity.ID, error) {
	item.Character.ID = entity.ID(uuid.New().String())
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

func (r *mongoPlayerCharRepository) GetByID(ctx context.Context, id entity.ID) (*entity.PlayerChar, error) {
	var result entity.PlayerChar
	err := r.collection.FindOne(ctx, bson.M{"_id": string(id)}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *mongoPlayerCharRepository) GetArray(ctx context.Context, ids []entity.ID) ([]entity.PlayerChar, error) {
	objectIDs := make([]string, len(ids))
	for i := range ids {
		objectIDs[i] = ids[i].String()
	}
	cursor, err := r.collection.Find(ctx, bson.M{"_id": bson.M{"$in": objectIDs}})
	if err != nil {
		return nil, err
	}
	var results []entity.PlayerChar
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoPlayerCharRepository) GetByCreatorID(ctx context.Context, id entity.ID) ([]entity.PlayerChar, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"creator_id": id.String()})
	if err != nil {
		return nil, err
	}
	var results []entity.PlayerChar
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoPlayerCharRepository) GetAll(ctx context.Context) ([]entity.PlayerChar, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var results []entity.PlayerChar
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoPlayerCharRepository) Update(ctx context.Context, item *entity.PlayerChar) error {
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": item.Character.ID}, item)
	return err
}

func (r *mongoPlayerCharRepository) Delete(ctx context.Context, id entity.ID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id.String()})
	return err
}
