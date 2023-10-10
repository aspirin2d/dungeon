package controller

import (
	"errors"
	"fmt"

	"github.com/aspirin2ds/dungeon/model"
	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	_ "github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
)

var usernameReg = regexp2.MustCompile("^(?=[a-zA-Z0-9._]{6,20}$)(?!.*[_.]{2})[^_.].*[^_.]$", 0)

func GetObjectId(c *gin.Context, name string) (primitive.ObjectID, error) {
	id := c.Param(name)
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("invalid uid: %s", id)
	}

	return uid, nil
}

// @Summary		get user by id
// @Tags			user
// @Param uid  path	string	true	"user id"
// @Produce		json
// @Success		200 {object} model.User "user information"
// @Failure		400 {object} controller.ErrorMessage "bad request"
// @Failure		500 {object} controller.ErrorMessage "internal error"
// @Router			/{uid} [get]
func GetUser(c *gin.Context) {
	uid, err := GetObjectId(c, "uid")
	if err != nil {
		c.AbortWithError(400, err).SetType(gin.ErrorTypePublic)
	}

	usr, err := model.GetUser(c.Request.Context(), uid)

	if err != nil && err == mongo.ErrNoDocuments {
		c.AbortWithError(400, fmt.Errorf("user not found")).SetType(gin.ErrorTypePublic)
		return
	}
	if err != nil {
		c.AbortWithError(500, err).SetType(gin.ErrorTypePrivate)
		return
	}

	c.JSON(200, usr)
}

type InsertedResponse struct {
	InsertedId primitive.ObjectID `json:"inserted" example:"652270c0305160d377fda9d6"`
}

// @Summary		create a new user
// @Tags			user
// @Accept		json
// @Param		  user_form	body controller.NewUser.UserForm	true "user form body"
// @Produce		json
// @Success		200 {object} controller.InsertedResponse "inserted user id"
// @Failure		400 {object} controller.ErrorMessage "bad request"
// @Failure		500 {object} controller.ErrorMessage "internal error"
// @Router			/new [post]
func NewUser(c *gin.Context) {
	type UserForm struct {
		Username string `json:"username" binding:"required" example:"your_unique_name"`
	}
	var form UserForm
	if err := c.Bind(&form); err != nil {
		c.AbortWithError(400, err)
		return
	}

	if matched, err := usernameReg.MatchString(form.Username); matched == false || err != nil {
		c.AbortWithError(400, errors.New("username invalid")).SetType(gin.ErrorTypePublic)
		return
	}

	inserted, err := model.NewUser(c, &model.User{Username: form.Username})
	if err != nil {
		return
	}

	c.JSON(200, InsertedResponse{InsertedId: inserted})
}
