package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// FruitModel represents the fruit object
//
// swagger:model
type Fruit struct {
	//  id
	// required: true
	// example: apple
	Id primitive.ObjectID `json:"id" bson:"_id"`

	// fruit name
	// required: true
	// example: apple
	Fruit string `json:"fruit" bson:"fruit"`

	// fruit color
	// required: true
	// example: red
	Color string `json:"color" bson:"color"`
}
