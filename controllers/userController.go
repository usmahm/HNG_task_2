package controllers

import (
	"fmt"
	"log"
	"net/http"
	"person-app/database"
	"person-app/helpers"
	"person-app/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var validate = validator.New()

// get mongodb filter depending if id or name is passed
func getFilter(param string) primitive.D {
	var filter primitive.D

	id, err := primitive.ObjectIDFromHex(param)
	if err != nil { // if error is nil, means user passed name not person id
		filter = bson.D{{Key: "name", Value: param}}
	} else {
		filter = bson.D{{Key: "_id", Value: id}}
	}

	return filter
}

func GetPerson(c *gin.Context) {
	var person_result models.Person
	var ctx, cancel = helpers.CreateContext()

	param := c.Param("param")

	filter := getFilter(param)

	err := database.OpenCollection("person").FindOne(ctx, filter).Decode(&person_result)
	defer cancel()

	if err != nil {
		fmt.Println(err)
		msg := fmt.Sprintf("an error occured")

		if err == mongo.ErrNoDocuments {
			msg = fmt.Sprintf("person not found")
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   msg,
		})
		return
	}

	fmt.Println(param)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    person_result,
	})
}

func CreatePerson(c *gin.Context) {
	var ctx, cancel = helpers.CreateContext()
	defer cancel()

	var new_person models.Person

	if err := c.BindJSON(&new_person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}

	if validation_error := validate.Struct(new_person); validation_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   validation_error.Error(),
		})
		return
	}

	new_person.ID = primitive.NewObjectID()
	new_person.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	new_person.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	// Add to db
	result, err := database.OpenCollection("person").InsertOne(ctx, new_person)

	if err != nil {
		log.Println(err)

		msg := fmt.Sprintf("Error creating food item")
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   msg,
		})
		return
	}

	fmt.Println(new_person, result)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    new_person,
		"message": "Person created successfully",
	})
}

func UpdatePerson(c *gin.Context) {
	var ctx, cancel = helpers.CreateContext()
	defer cancel()

	param := c.Param("param")
	var person models.Person

	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: person.Name}}}}
	filter := getFilter(param)

	updateOptions := options.FindOneAndUpdate().SetReturnDocument((options.After))
	err := database.OpenCollection("person").FindOneAndUpdate(ctx, filter, update, updateOptions).Decode(&person)
	if err != nil {
		msg := fmt.Sprintf("an error occured")

		if err == mongo.ErrNoDocuments {
			msg = fmt.Sprintf("person not found")
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   msg,
		})
		return
	}

	fmt.Println(param)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Person updated successfully",
		"data":    person,
	})
}

func DeletePerson(c *gin.Context) {
	var ctx, cancel = helpers.CreateContext()
	defer cancel()
	param := c.Param("param")

	filter := getFilter(param)
	result, err := database.OpenCollection("person").DeleteOne(ctx, filter)
	if err != nil {
		msg := fmt.Sprintf("an error occured")

		if err == mongo.ErrNoDocuments {
			msg = fmt.Sprintf("person not found")
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   msg,
		})
		return
	}

	fmt.Println("REsult", result)
	message := fmt.Sprintf("User %s deleted", param)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
	})
}
