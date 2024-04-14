package database

import (
	"context"
	"fmt"
	"log"
	"testApi/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const connecionString2 string = "mongodb+srv://omarBlazi:omarblaziMongoDB54@cluster0.aoo2r6v.mongodb.net/"
const connecionString string = "mongodb+srv://omarBlazi:xhFokVVoQ3qXh704@cluster1.aoo2r6v.mongodb.net/testDb?retryWrites=true&w=majority&appName=Cluster0" 
const dbName = "blazi"

const collectionName = "orders"


var collection *mongo.Collection
func init(){
	clientOptions := options.Client().ApplyURI(connecionString);

	client ,err := mongo.Connect(context.TODO(),clientOptions);

	if err != nil{
		log.Fatal(err);
	}

	fmt.Println("Connected Successfully to mongoDB");

	collection = client.Database(dbName).Collection(collectionName);

	fmt.Println("Collection ref is READY")
}

//MongoDB Helpers

func InsertOrder(order models.Order){
	order.ID = primitive.NewObjectID()

	inserted , err :=collection.InsertOne(context.TODO(),order);
	if err != nil {
		log.Fatal(err);
	}
	fmt.Println("Inserted Succesfully",inserted.InsertedID)
}

func UpdateOrder(orderId string,order models.Order)(*mongo.UpdateResult,error){
	id, err := primitive.ObjectIDFromHex(orderId);
	if err != nil{
	 	log.Fatal(err);
	 }
	filter := bson.M{
		"_id":id,
	}
	updated := bson.M{"$set":bson.M{
		"quantity": order.Quantity,
	}}

	result ,err:= collection.UpdateOne(context.TODO(),filter,updated)
	if err != nil{
		log.Fatal(err);
		return nil,err;
	}
	fmt.Printf("Result is %v ",result)
	return result,nil;
}


func Delete(orderId string)bool{
	id,err:= primitive.ObjectIDFromHex(orderId)
	fmt.Printf("the hex obj id is %v",id);
	if err != nil {
		log.Fatal(err)
	}	
	filter:= bson.M{"_id":id};
	result,err := collection.DeleteOne(context.TODO(),filter) 
	if err != nil{
		log.Fatal(err);
		return false;
	}
	fmt.Printf("The count of delete is %v",result)
	return true;
}

func DeleteAll() int64{
	deleted,err := collection.DeleteOne(context.TODO(),bson.D{{}}) 
	if err != nil{
		log.Fatal(err);
	}

	fmt.Println("Deleted count", deleted.DeletedCount)
	return deleted.DeletedCount
}


func GetAll() []primitive.M{
	result,err := collection.Find(context.TODO(),bson.D{{}});
	if err != nil {
		log.Fatal(err);
	}
	var orders []primitive.M

	for result.Next(context.TODO()){
		var order bson.M
		err := result.Decode(&order)
		if err != nil{
			log.Fatal(err);
		}
		orders = append(orders, order);
	}
	return orders;
}
func GetById(orderId string) (*mongo.SingleResult, error) {
	id,err:= primitive.ObjectIDFromHex(orderId)
	fmt.Printf("the hex obj id is %v",id);
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{
		"_id": id,
	}
	res := collection.FindOne(context.TODO(), filter)
	fmt.Printf("the result of search by id is %v",res)
	if err := res.Err(); err != nil {
		return nil, err
	}
	
	return res, nil
}