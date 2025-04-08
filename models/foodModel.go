package models

//we need to import time for updated at and created at feilds
//we can also use GORM but we are not using it in this project
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Food struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          *string            `json:"name" validate:"required,min=2,max=100"` //this means that it should between 2-100 chracters
	Price         *float64           `json:"price" validate:"required"`              //this means that it should be a float number
	Food_image_at time.Time          `json:"food_image" validate:"required"`
	Created_at    time.Time          `json:"created_at" validate:"required"`
	Updated_at    time.Time          `json:"updated_at" validate:"required"`
	Food_id       string             `json:"food_id" validate:"required"`
	Menu_id       *string            `json:"menu_id" validate:"required"`
}
//we put menu_id as a pointer as it is coming from the menu model and it is a pointer in the menu model