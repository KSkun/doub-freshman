package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type Flag struct {
	Text  string  `bson:"text" json:"text"`
	Value float64 `bson:"value" json:"value"`
	Hide  bool    `bson:"hide" json:"hide"`
}

/*
	Condition types:
	flag:               require a flag
	flag, op, value:    require a flag with value satisfying the expression
		op: gt, lt, gte, lte, eq, neq
	flag, exclude:      when a flag exists, exclude this stage in selections (entrance only)
	<empty>, prob, 0.1: set success probability (option only)
*/

type Condition struct {
	Flag  string  `bson:"flag"`
	Op    string  `bson:"op"`
	Value float64 `bson:"value"`
}

type OptionBranch struct {
	Next   primitive.ObjectID `bson:"next"`
	Text   string             `bson:"text"`
	Event  []Event            `bson:"event"`
}

type Option struct {
	Text      string       `bson:"text"`
	Success   OptionBranch `bson:"success"`
	Failed    OptionBranch `bson:"failed"`
	Condition []Condition  `bson:"condition"`
}

/*
	Event types: flag, death, extend
	Event patterns:
	flag, flag=name:                add a flag
	flag, flag=name, delete=true:   delete a flag
	flag, flag=name, inc=10:        increase the value of a flag
	death:                          set the player dead
	extend, extend=1:               extend the round number of current phase
*/

type Event struct {
	Type  string                 `bson:"type"`
	Value map[string]interface{} `bson:"value"`
}

type Stage struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Text      string             `bson:"text"`
	Dead      bool               `bson:"dead"`
	EnterCond []Condition        `bson:"enter_cond"`
	Option    []Option           `bson:"option"`
	Tag       string             `bson:"tag"`

	Continue bool `bson:"continue"`
	Delay    int  `bson:"delay"`
}

func (m *model) cStage() *mongo.Collection {
	return m.db.Collection("stage")
}

func (m *model) GetStageWithCondition(condition bson.M) ([]Stage, error) {
	result, err := m.cStage().Find(m.ctx, condition)
	if err != nil {
		return nil, err
	}
	var stage []Stage
	err = result.All(m.ctx, &stage)
	return stage, err
}

func (m *model) GetStageWithFlags(flag []string) ([]Stage, error) {
	var filter []bson.M
	for _, _flag := range flag {
		filter = append(filter, bson.M{"enter_cond": bson.M{"flag": _flag}})
	}
	return m.GetStageWithCondition(bson.M{"$and": filter})
}

func (m *model) GetStageWithFlag(flag string) ([]Stage, error) {
	return m.GetStageWithFlags([]string{flag})
}

func (m *model) GetStageWithFlagExclude(flag string) ([]Stage, error) {
	return m.GetStageWithCondition(bson.M{"enter_cond": bson.M{"flag": flag, "op": "exclude"}})
}

func (m *model) GetStage(id primitive.ObjectID) (Stage, error) {
	stage := Stage{}
	err := m.cStage().FindOne(m.ctx, bson.M{"_id": id}).Decode(&stage)
	return stage, err
}

func (m *model) GetStageByHex(hex string) (Stage, error) {
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return Stage{}, err
	}
	return m.GetStage(id)
}

func (m *model) AddStages(stages []Stage) error {
	var stagesInterface []interface{}
	for _, stage := range stages {
		stagesInterface = append(stagesInterface, stage)
	}

	_, err := m.cStage().InsertMany(context.Background(), stagesInterface)
	if err != nil {
		return err
	}
	return nil
}
