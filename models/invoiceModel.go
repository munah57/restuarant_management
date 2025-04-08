package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"` //bson:"_id" is used to create a unique id for each invoice
	Invoice_id       string             `json:"invoice_id"`
	Order_id         string             `json:"order_id"`
	Payment_method   *string            `json:"payment_method" validate:"eq=CARD|eq=CASH|eq="` //either card cash or empty
	Payment_status   *string            `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	Payment_due_date time.Time          `json:"payment_due_date"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
}


/*
primitive.ObjectID is a type from the MongoDB Go driver that represents a MongoDB _id field, which is usually an auto-generated 12-byte unique identifier.
MongoDB uses it by default as the primary key (_id) for documents unless you specify your own.


bson stands for Binary JSON — it’s the format MongoDB uses internally to store documents. It's similar to JSON but optimized for performance and type safety.
In Go, the bson:"..." struct tag tells the MongoDB driver how to map Go struct fields to MongoDB document fields.
*/