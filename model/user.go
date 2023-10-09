package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username" json:"username"`
	Created  time.Time          `bson:"created" json:"created"`
}

func GetUser(ctx context.Context, uid primitive.ObjectID) (*User, error) {
	var usr User
	res := Collection("users").FindOne(ctx, bson.M{"_id": uid})
	err := res.Err()
	if err != nil {
		return nil, err
	}
	err = res.Decode(&usr)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}
