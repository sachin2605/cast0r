package svc

import (
	"context"
	"log"
	"time"

	"github.com/sachin2605/cast0r/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FruitService struct{}

func NewFruitService() FruitService {
	return FruitService{}
}

// List list all Fruits
func (br *FruitService) List(page, limit int) []models.Fruit {
	log.Println("[FruitService] List() Getting Fruits from Mongo.")
	offset := (page - 1) * limit
	collection := DB.Collection("fruits")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetLimit(int64(limit)).SetSkip(int64(offset)))

	if err != nil {
		log.Fatal("Error while retriving Fruits..")
	}
	defer cursor.Close(ctx)

	var Fruits []models.Fruit
	for cursor.Next(ctx) {
		var Fruit models.Fruit
		cursor.Decode(&Fruit)
		Fruits = append(Fruits, Fruit)
	}
	if err := cursor.Err(); err != nil {
		log.Println("Error while traversing Curor")
		//return error
	}
	return Fruits
}

// Get Fruit by ID
func (br *FruitService) Get(id string) (f models.Fruit, err error) {
	col := DB.Collection("fruits")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)

	var b models.Fruit
	err = col.FindOne(ctx, bson.M{"_id": objID}).Decode(&b)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return f, err
		}
		panic(err)
	}
	return b, nil
}

// Create creates new Fruit Document in Mongo
func (br *FruitService) Create(Fruit models.Fruit) interface{} {
	collection := DB.Collection("fruits")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	Fruit.Id = primitive.NewObjectID()

	result, err := collection.InsertOne(ctx, Fruit)

	if err != nil {
		log.Fatal("Error while Inserting new Fruit in Mongo.", err.Error())
		log.Fatal(err.Error())
	}
	return result.InsertedID
}
