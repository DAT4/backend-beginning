package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email"`
	Name     string             `json:"name" bson:"name"`
	Password string             `json:"-" bson:"password"`
	Mac      string             `json:"-" bson:"mac"`
	Ip       string             `json:"-" bson:"ip"`
}

const URI string = "mongodb://localhost:27017"

func GetUsers() (users []User) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := client.Database("game").Collection("users").Find(ctx, bson.M{})
	var user User
	for cursor.TryNext(context.Background()) {
		cursor.Decode(&user)
		users = append(users, user)
	}
	defer client.Disconnect(ctx)
	return users
}

func CreateUser(user User) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{}
	cursor, err := client.Database("game").Collection("users").Find(ctx, filter)
	if err != nil {
		fmt.Println(err)
	}

	ok := cursor.TryNext(context.Background())
	if ok {
		fmt.Println("Exist already")
		return
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	_, err = client.Database("game").Collection("users").InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err)
	}
}

func RemoveUser(user User) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": user.Id}
	_, err = client.Database("game").Collection("users").DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateUser(user User) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": user}
	_, err = client.Database("game").Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}
}

func GetUser(id string) (user *User, err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))

	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	defer client.Disconnect(ctx)

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"_id": objectId,
	}

	cursor, err := client.Database("game").Collection("users").Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	cursor.TryNext(context.Background())

	err = cursor.Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (self *User) Authenticate() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	defer client.Disconnect(ctx)

	if err != nil {
		fmt.Println(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{
		"email":    self.Email,
		"password": self.Password,
	}

	cursor, err := client.Database("game").Collection("users").Find(ctx, filter)

	if err != nil {
		return err
	}

	cursor.TryNext(context.Background())

	err = cursor.Decode(&self)

	if err != nil {
		return err
	}

	return nil
}
