package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ImgRepo struct {
	db *mongo.Database
}

func NewImgRepo(db *mongo.Database) *ImgRepo {
	return &ImgRepo{db: db}
}

func (r *ImgRepo) Create(img string) (string, error) {
	coll := r.db.Collection(IMG_COLLECTION)
	resOfSearch := coll.FindOne(context.Background(), bson.M{"img": img})
	if resOfSearch.Err() == mongo.ErrNoDocuments {
		res, err := coll.InsertOne(context.Background(), bson.M{"img": img})
		if err != nil {
			return "", err
		}
		id, ok := res.InsertedID.(primitive.ObjectID)
		if !ok {
			return "", errors.New("error represent insert id as primitive.ObjectID")
		}
		return id.Hex(), nil
	}
	return "", errors.New("img already exist")
}
