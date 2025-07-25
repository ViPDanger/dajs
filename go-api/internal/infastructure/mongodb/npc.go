package mongodb

import (
	"context"
	"errors"

	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoNPCRepository struct {
	collection *mongo.Collection
}

func NewMongoNPCRepository(db *mongo.Database) *mongoNPCRepository {
	return &mongoNPCRepository{
		collection: db.Collection("npcs"),
	}
}

func (r *mongoNPCRepository) Insert(ctx context.Context, item *entity.NPC) (*entity.ID, error) {
	res, err := r.collection.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to cast inserted ID to ObjectID")
	}
	id := entity.ID(oid.Hex())
	return &id, nil
}

func (r *mongoNPCRepository) GetByID(ctx context.Context, id entity.ID) (*entity.NPC, error) {
	oid, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return nil, err
	}
	var result entity.NPC
	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *mongoNPCRepository) GetByCreatorID(ctx context.Context, id entity.ID) ([]*entity.NPC, error) {
	oid, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return nil, err
	}
	cursor, err := r.collection.Find(ctx, bson.M{"_id": bson.M{"$in": oid}})
	if err != nil {
		return nil, err
	}
	var results []*entity.NPC
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoNPCRepository) GetArray(ctx context.Context, ids []entity.ID) ([]*entity.NPC, error) {
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
	var results []*entity.NPC
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoNPCRepository) GetAll(ctx context.Context) ([]*entity.NPC, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var results []*entity.NPC
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoNPCRepository) Update(ctx context.Context, item *entity.NPC) error {
	oid, err := primitive.ObjectIDFromHex(item.ID.String())
	if err != nil {
		return err
	}
	_, err = r.collection.ReplaceOne(ctx, bson.M{"_id": oid}, item)
	return err
}

func (r *mongoNPCRepository) Delete(ctx context.Context, id entity.ID) error {
	oid, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}
