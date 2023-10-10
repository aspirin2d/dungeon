package model

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username" json:"username"`
	Created  time.Time          `bson:"created" json:"created"`

	Characters []primitive.ObjectID `bson:"characters,omitempty" json:"characters,omitempty"`
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

func NewUser(c *gin.Context, usr *User) (primitive.ObjectID, error) {
	usr.Created = time.Now()
	res, err := Collection("users").InsertOne(c.Request.Context(), usr)
	if err != nil {
		// if username already exists
		if mongo.IsDuplicateKeyError(err) {
			c.AbortWithError(400, fmt.Errorf("username already exists: %s", usr.Username)).SetType(gin.ErrorTypePublic)
		} else {
			c.AbortWithError(500, err).SetType(gin.ErrorTypePrivate)
		}

		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}
