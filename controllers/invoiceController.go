package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetInvoices() gin.HandlerFunc {

	return func(c *gin.Context){
	} //we are returning a function that takes a gin context
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
	}
}