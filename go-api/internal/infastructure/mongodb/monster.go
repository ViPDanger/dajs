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

type mongoMonsterRepository struct {
	collection *mongo.Collection
}

func NewMonsterRepository(db *mongo.Database) repository.MonsterRepository {
	return &mongoMonsterRepository{
		collection: db.Collection("monsters"),
	}
}

func (r *mongoMonsterRepository) Insert(ctx context.Context, item *entity.Monster) (*entity.ID, error) {
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

func (r *mongoMonsterRepository) GetByID(ctx context.Context, id entity.ID) (*entity.Monster, error) {
	var result entity.Monster
	err := r.collection.FindOne(ctx, bson.M{"_id": string(id)}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *mongoMonsterRepository) GetArray(ctx context.Context, ids []entity.ID) ([]entity.Monster, error) {
	objectIDs := make([]string, len(ids))
	for i := range ids {
		objectIDs[i] = ids[i].String()
	}
	cursor, err := r.collection.Find(ctx, bson.M{"_id": bson.M{"$in": objectIDs}})
	if err != nil {
		return nil, err
	}
	var results []entity.Monster
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoMonsterRepository) GetByCreatorID(ctx context.Context, id entity.ID) ([]entity.Monster, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"creator_id": id.String()})
	if err != nil {
		return nil, err
	}
	var results []entity.Monster
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoMonsterRepository) GetAll(ctx context.Context) ([]entity.Monster, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var results []entity.Monster
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoMonsterRepository) Update(ctx context.Context, item *entity.Monster) error {
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": item.ID}, item)
	return err
}

func (r *mongoMonsterRepository) Delete(ctx context.Context, id entity.ID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id.String()})
	return err
}
