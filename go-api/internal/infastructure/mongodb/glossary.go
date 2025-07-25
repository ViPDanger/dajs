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

type mongoGlossaryRepository struct {
	collection *mongo.Collection
}

func NewMongoGlossaryRepository(db *mongo.Database) *mongoGlossaryRepository {
	return &mongoGlossaryRepository{
		collection: db.Collection("glossarys"),
	}
}

func (r *mongoGlossaryRepository) Insert(ctx context.Context, item *entity.Glossary) (*entity.ID, error) {
	if item.ID.String() == "" {
		item.ID = entity.ID(uuid.New().String())
	}
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

func (r *mongoGlossaryRepository) GetByID(ctx context.Context, id entity.ID) (*entity.Glossary, error) {
	oid, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return nil, err
	}
	var result entity.Glossary
	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *mongoGlossaryRepository) GetArray(ctx context.Context, ids []entity.ID) ([]*entity.Glossary, error) {
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
	var results []*entity.Glossary
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoGlossaryRepository) GetAll(ctx context.Context) ([]*entity.Glossary, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var results []*entity.Glossary
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *mongoGlossaryRepository) Update(ctx context.Context, item *entity.Glossary) error {
	oid, err := primitive.ObjectIDFromHex(item.ID.String())
	if err != nil {
		return err
	}
	_, err = r.collection.ReplaceOne(ctx, bson.M{"_id": oid}, item)
	return err
}

func (r *mongoGlossaryRepository) Delete(ctx context.Context, id entity.ID) error {
	oid, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}
