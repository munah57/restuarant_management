package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context){
	}
}

func GetTable() gin.HandlerFunc{
	return func(c *gin.Context){
	}
}

func CreateTable() gin.HandlerFunc{
	return func(c *gin.Context){
	}
}

func UpdateTable() gin.HandlerFunc{
	return func(c *gin.Context){
	}
}

// create gin handler function the syntax here is that we are taking a function that takes a gin context and returns a function that takes a gin context
// this is a common pattern in gin where we create a handler function that takes a context and returns a function that takes a context
// this allows us to create middleware functions that can be used in the router
