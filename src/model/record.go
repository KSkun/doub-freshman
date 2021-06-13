package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StageResult struct {
	Stage   primitive.ObjectID `bson:"stage"`
	Option  int                `bson:"option"`
	Success bool               `bson:"success"`
}

type Record struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	PlayerID string             `bson:"player_id"`
	Flag     []Flag             `bson:"flag"`
	Result   []StageResult      `bson:"result"`
}

func (m *model) cRecord() *mongo.Collection {
	return m.db.Collection("record")
}

func (m *model) GetRecordByName(name string) (Record, bool, error) {
	record := Record{}
	err := m.cRecord().FindOne(m.ctx, bson.M{"name": name}).Decode(&record)
	if err == mongo.ErrNoDocuments {
		return record, false, nil
	}
	if err != nil {
		return record, false, err
	}
	return record, true, nil
}

func (m *model) NewRecord(name string, playerID string) error {
	record := Record{
		ID:       primitive.NewObjectID(),
		Name:     name,
		PlayerID: playerID,
		Flag:     []Flag{},
		Result:   []StageResult{},
	}
	_, err := m.cRecord().InsertOne(m.ctx, record)
	return err
}
