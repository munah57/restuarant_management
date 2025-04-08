package controllers

import (
	"github.com/gin-gonic/gin"
)


// create gin handler function the syntax here is that we are taking a function that takes a gin context and returns a function that takes a gin context
// this is a common pattern in gin where we create a handler function that takes a context and returns a function that takes a context
// this allows us to create middleware functions that can be used in the router


func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context){
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context){
	}
}


func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context){
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context){
	}
}


//

func round(num float64) float64 {

}


func toFixed(num float64, precision int) float64{
	
}