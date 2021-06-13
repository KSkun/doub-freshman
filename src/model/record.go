package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type StageResult struct {
	Stage   primitive.ObjectID `bson:"stage"`
	Option  int                `bson:"option"`
	Success bool               `bson:"success"`
}

type Record struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Flag   []Flag             `bson:"flag"`
	Result []StageResult      `bson:"result"`
}
