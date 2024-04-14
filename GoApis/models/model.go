package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID       primitive.ObjectID `bson:"_id"`
	Quantity int                `bson:"quantity"`
}

// Read implements io.Reader.
func (o Order) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}
