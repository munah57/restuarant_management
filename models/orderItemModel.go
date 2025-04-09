package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID primitive.ObjectID `bson:"_id"`
	Quantity *string `json:"quantity" validate:"required,eq=S|eq=M|eq=L"`  //small medium or large portions 
	Unit_price *float64 `json:"unit_price" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Food_id *string `json:"food_id" validate:"required"`//its a string because it is a foreign key from the food model
	Order_item_id string `json:"order_item_id"`//not a pointer because it is not required
	Order_id string `json:"order_id"` 

}

//If Order_id is always set in your logic (e.g., you're assigning it yourself during processing), then it's okay to leave it as a plain string.No difference in capability: Both Food_id and Order_id can be pointers.

//Pointers are used when you want to detect "not provided", handle optional fields, or work with validators.
//You’re probably using *string for Food_id because you’re validating input where Food_id is required and may come in as nil/empty.
//If you want similar flexibility with Order_id, feel free to switch it to *string.