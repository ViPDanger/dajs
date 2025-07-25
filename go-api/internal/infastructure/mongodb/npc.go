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

type mongoNPCRepository struct {
	collection *mongo.Collection
}

func NewNPCRepository(db *mongo.Database) repository.NPCRepository {
	return &mongoNPCRepository{
		collection: db.Collection("npcs"),
	}
}

func (r *mongoNPCRepository) Insert(ctx context.Context, item *entity.NPC) (*entity.ID, error) {
	item.ID = entity.ID(uuid.New().String())
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

func (r *mongoNPCRepository) GetByID(ctx context.Context, id entity.ID) (*entity.NPC, error) {
	var result entity.NPC
	err := r.collection.FindOne(ctx, bson.M{"_id": string(id)}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *mongoNPCRepository) GetArray(ctx context.Context, ids []entity.ID) ([]entity.NPC, error) {
	objectIDs := make([]string, len(ids))
	for i := range ids {
		objectIDs[i] = ids[i].String()
	}
	cursor, err := r.collection.Find(ctx, bson.M{"_id": bson.M{"$in": objectIDs}})
	if err != nil {
		return nil, err
	}
	var results []entity.NPC
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoNPCRepository) GetByCreatorID(ctx context.Context, id entity.ID) ([]entity.NPC, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"creator_id": id.String()})
	if err != nil {
		return nil, err
	}
	var results []entity.NPC
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoNPCRepository) GetAll(ctx context.Context) ([]entity.NPC, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var results []entity.NPC
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoNPCRepository) Update(ctx context.Context, item *entity.NPC) error {
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": item.ID}, item)
	return err
}

func (r *mongoNPCRepository) Delete(ctx context.Context, id entity.ID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id.String()})
	return err
}
