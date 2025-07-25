package mongodb

import (
	"context"
	"errors"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoItemRepository struct {
	collection *mongo.Collection
}

func NewMongoItemRepository(db *mongo.Database) *mongoItemRepository {
	return &mongoItemRepository{
		collection: db.Collection("items"),
	}
}

func (r *mongoItemRepository) Insert(ctx context.Context, item entity.Item) (*entity.ID, error) {
	if item.GetSimpleItem().ID.String() == "" {
		item.GetSimpleItem().ID = entity.ID(uuid.New().String())
	}
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

func (r *mongoItemRepository) GetByID(ctx context.Context, id entity.ID) (entity.Item, error) {
	oid, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return nil, err
	}
	var result entity.Item
	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *mongoItemRepository) GetArray(ctx context.Context, ids []entity.ID) ([]entity.Item, error) {
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
	var results []entity.Item
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoItemRepository) GetAll(ctx context.Context) ([]entity.Item, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var results []entity.Item
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoItemRepository) Update(ctx context.Context, item entity.Item) error {
	oid, err := primitive.ObjectIDFromHex(item.GetSimpleItem().ID.String())
	if err != nil {
		return err
	}
	_, err = r.collection.ReplaceOne(ctx, bson.M{"_id": oid}, item)
	return err
}

func (r *mongoItemRepository) Delete(ctx context.Context, id entity.ID) error {
	oid, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}
