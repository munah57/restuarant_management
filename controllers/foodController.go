package controllers

import (
	"context"
	"go-resturant-management/database"
	"go-resturant-management/models"
	"net/http"
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// create gin handler function the syntax here is that we are taking a function that takes a gin context and returns a function that takes a gin context
// this is a common pattern in gin where we create a handler function that takes a context and returns a function that takes a context
// this allows us to create middleware functions that can be used in the router

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		foodId := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel() //cancel defer to free up resources after the function is done executing
		//if there is an error in finding the food item that means it does not exist and we cannot add
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching the food"})
		} //catch error here
		c.JSON(http.StatusOK, food) //the context is the response we are sending back to the client. Here we are sending the food object back to the client
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Food

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//Validate what we are recieving. - we validate outside the func
		//if error occers we handle error as such
		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		//we will find the menu that exists and then add it to using the food id
		//we know that food items belong to a menu

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		//if there is a an error in finding the menu that means it does not exist and we cannot add
		//the food item to the menu
		if err != nil {
			msg := fmt.Sprintf("menu with id %v not found", food.Menu_id)
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		food.Created_at, _ = time.Parse(time.RFC3339, time.Now()).Format(time.RFC3339)
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now()).Format(time.RFC3339)
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()      // this will give us the hex value of the object id
		var num = toFixed(*food.Price, 2) //precision unit is the number of decimal places we want to keep
		food.Price = &num
		//everything that is required in the food model we have it here so we can add it to the database

		result, insertErr := foodCollection.InsertOne(ctx, food) 
		if insertErr !=nil {
			msg := fmt.Sprintf("food item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
		
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

//

func round(num float64) float64 {

}

func toFixed(num float64, precision int) float64 {

}
