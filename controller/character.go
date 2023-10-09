package controller

import (
	"fmt"
	"time"

	"github.com/aspirin2ds/dungeon/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary		get character by id
// @Tags			character
// @Param		uid	path	string	true	"user id"
// @Param		cid	path	string	true	"character id"
// @Produce		json
// @Success		200 {object} model.Character "character information"
// @Failure		400 {object} controller.ErrorMessage "bad request"
// @Failure		500 {object} controller.ErrorMessage "internal error"
// @Router			/{uid}/{cid} [get]
func GetCharacter(c *gin.Context) {
	var err error

	uid, err := GetObjectId(c, "uid")
	if err != nil {
		c.AbortWithError(400, err).SetType(gin.ErrorTypePublic)
		return
	}

	cid, err := GetObjectId(c, "cid")
	if err != nil {
		c.AbortWithError(400, err).SetType(gin.ErrorTypePublic)
		return
	}

	cha, err := model.GetCharacter(c, cid)
	if err != nil {
		// already handled in model.GetCharacter
		return
	}

	if cha.Owner.Hex() != uid.Hex() {
		c.AbortWithError(400, fmt.Errorf("character's owner is not %s", uid)).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(200, cha)
}

// @Summary		create a new character
// @Tags			character
// @Accept		multipart/form-data
// @Param		uid	path	string	true	"user id"
// @Param		name	formData	string	true	"character name"
// @Param		race	formData	string	true	"race id"
// @Param		class	formData	string	true	"class id"
// @Produce		json
// @Success		200 {object} controller.InsertedResponse "inserted character id"
// @Failure		400 {object} controller.ErrorMessage "bad request"
// @Failure		500 {object} controller.ErrorMessage "internal error"
// @Router			/{uid}/new [post]
func NewCharacter(c *gin.Context) {
	var ctx = c.Request.Context()
	var err error

	uid, err := GetObjectId(c, "uid")
	if err != nil {
		c.AbortWithError(400, err).SetType(gin.ErrorTypePublic)
		return
	}

	type CharacterForm struct {
		Name  string `form:"name" binding:"required"`
		Race  string `form:"race" binding:"required"`
		Class string `form:"class" binding:"required"`
	}

	var form CharacterForm
	if err = c.Bind(&form); err != nil {
		c.AbortWithError(400, err)
		return
	}

	// TODO: validate username

	race := model.GetRace(form.Race)
	if race == nil {
		c.AbortWithError(400, fmt.Errorf("race not found: %s", form.Race)).SetType(gin.ErrorTypePublic)
		return
	}

	class := model.GetClass(form.Class)
	if class == nil {
		c.AbortWithError(400, fmt.Errorf("class not found: %s", form.Class)).SetType(gin.ErrorTypePublic)
		return
	}

	char := &model.Character{
		Name:    form.Name,
		Created: time.Now(),
		Owner:   uid,
		Race:    form.Race,
		Class:   form.Class,
	}

	res, err := model.Collection("characters").InsertOne(ctx, char)
	if err != nil {
		c.AbortWithError(500, err).SetType(gin.ErrorTypePrivate)
		return
	}

	c.JSON(200, InsertedResponse{InsertedId: res.InsertedID.(primitive.ObjectID)})
}
