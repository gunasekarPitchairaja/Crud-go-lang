package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"makecrud/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type BookController struct {
	session *mongo.Client
}

func NewBookController(s *mongo.Client) *BookController {
	return &BookController{s}
}


func (db BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var dbCollection = db.session.Database("booksDB").Collection("books")
	book := models.Books{}
	err := json.NewDecoder(r.Body).Decode(&book) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(book.Id)
	dbCollection.InsertOne(context.TODO(), book)
	json.NewEncoder(w).Encode(book)

}

func (db BookController) Books(w http.ResponseWriter, r *http.Request) {
	var dbCollection = db.session.Database("booksDB").Collection("books")
	var results []primitive.M

	cur, err := dbCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	for cur.Next(context.TODO()) {
		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	json.NewEncoder(w).Encode(results)
}

func (db BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	var dbCollection = db.session.Database("booksDB").Collection("books")
	results := models.Books{}
	val := mux.Vars(r)
	docID, err := primitive.ObjectIDFromHex(val["id"])
	if err != nil {
		fmt.Println(err)
	}
	errr := dbCollection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&results)
	if errr != nil {
		fmt.Println(errr)
	}

	json.NewEncoder(w).Encode(results)
}

func (db BookController) SearchBook(w http.ResponseWriter, r *http.Request) {
	var dbCollection = db.session.Database("booksDB").Collection("books")
	title := mux.Vars(r)
	var result primitive.M
	err := dbCollection.FindOne(context.TODO(), bson.M{"title": title["title"]}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(result)

}

func (db BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var dbCollection = db.session.Database("booksDB").Collection("books")
	body := models.UpdateBody{}

	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	docID, err := primitive.ObjectIDFromHex(body.Id.Hex())
	if err != nil {
		fmt.Println(err)
	}
	filter := bson.M{"title": body.Title, "_id": docID}

	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.M{"$set": bson.M{"price": body.Price}}
	updateResult := dbCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

func (db BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	var dbCollection = db.session.Database("booksDB").Collection("books")
	params := mux.Vars(r)
	res := models.Books{}
	docID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		fmt.Println(err)
	}

	filter := bson.M{"_id": docID}
	e := dbCollection.FindOneAndDelete(context.TODO(), filter).Decode(&res)
	if e != nil {

		fmt.Print(e)
	}
	json.NewEncoder(w).Encode(res)

}
