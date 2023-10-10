package model

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Character struct {
	Id    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Owner primitive.ObjectID `bson:"owner" json:"owner"`

	Name    string    `bson:"name" json:"name"`
	Created time.Time `bson:"created" json:"created"`

	Race  string `bson:"race" json:"race"`
	Class string `bson:"class" json:"class"`
}

func GetCharacter(c *gin.Context, cid primitive.ObjectID) (*Character, error) {
	var cha Character
	res := Collection("characters").FindOne(c.Request.Context(), bson.M{"_id": cid})
	err := res.Err()

	if err == mongo.ErrNoDocuments {
		c.AbortWithError(404, fmt.Errorf("character not found: %s", cid)).SetType(gin.ErrorTypePublic)
		return nil, err
	}
	if err != nil {
		c.AbortWithError(500, err).SetType(gin.ErrorTypePrivate)
		return nil, err
	}

	err = res.Decode(&cha)
	if err != nil {
		c.AbortWithError(500, err).SetType(gin.ErrorTypePrivate)
		return nil, err
	}
	return &cha, nil
}

func InsertCharacter(c *gin.Context, uid primitive.ObjectID, cha *Character) (primitive.ObjectID, error) {
	cha.Owner = uid
	cha.Created = time.Now()
	res, err := Collection("characters").InsertOne(c.Request.Context(), cha)
	if err != nil {
		c.AbortWithError(500, err).SetType(gin.ErrorTypePrivate)
		return primitive.NilObjectID, err
	}

	_, err = Collection("users").UpdateByID(c.Request.Context(), uid, bson.M{"$push": bson.M{"characters": res.InsertedID}})
	if err != nil {
		c.AbortWithError(500, err).SetType(gin.ErrorTypePrivate)
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}
