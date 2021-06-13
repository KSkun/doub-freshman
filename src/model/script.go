package model

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	prob, 0.1:          set success probability
*/

type Condition struct {
	Flag  string  `bson:"flag"`
	Op    string  `bson:"op"`
	Value float64 `bson:"value"`
}

type OptionBranch struct {
	Next primitive.ObjectID `bson:"next"`
	Text string             `bson:"text"`
}

type Option struct {
	Success   OptionBranch `bson:"success"`
	Failed    OptionBranch `bson:"failed"`
	Condition []Condition  `bson:"condition"`
}

/*
	Event types: flag, death, extend
	Event patterns:
	flag:               add a flag
	flag, delete=true:  delete a flag
	flag, inc=10:       increase the value of a flag
	death:              set the player dead
	extend:             extend the round number of current phase
*/

type Event struct {
	Type  string                 `bson:"type"`
	Value map[string]interface{} `bson:"value"`
}

type Stage struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Text      string             `bson:"text"`
	EnterCond []Condition        `bson:"enter_cond"`
	Option    []Option           `bson:"option"`
	Event     []Event            `bson:"event"`
	Tag       string             `bson:"tag"`

	Continue bool `bson:"continue"`
	Delay    int  `bson:"delay"`
}
