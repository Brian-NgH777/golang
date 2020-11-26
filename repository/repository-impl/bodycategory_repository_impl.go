package repositoryimpl

import (
	"context"
	"fmt"
	models "learn/model"
	repo "learn/repository"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BodycategoryRepoImpl struct {
	Db *mongo.Database
}

func NewBodycategoryRepo(db *mongo.Database) repo.BodycategoriesRepository {
	return &BodycategoryRepoImpl{
		Db: db,
	}
}

func (mongo *BodycategoryRepoImpl) FindBodycategories() error {
	// Here's an array in which you can store the decoded documents
	// bodycategories := []interface{}

	// Pass these options to the Find method
	findOptions := options.Find()
	// findOptions.SetLimit(2)

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := mongo.Db.Collection("bodycategories").
		Find(context.Background(),
			bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem bson.M
		if err := cur.Decode(&elem); err != nil {
			log.Fatal(err)
		}
		fmt.Println("elemelem", elem)
		// bodycategories = append(bodycategories, &elem)
	}

	// Close the cursor once finished
	// cur.Close(context.TODO())

	return nil

}

// InsertBodycategory is install Bodycategory
func (mongo *BodycategoryRepoImpl) InsertBodycategory(bodycategory models.Bodycategories) error {
	bbytes, _ := bson.Marshal(bodycategory)

	_, err := mongo.Db.Collection("bodycategories").InsertOne(context.Background(), bbytes)
	if err != nil {
		return err
	}

	return nil
}
