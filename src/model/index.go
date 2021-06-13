package model

import (
	"context"
	"github.com/KSkun/doub-freshman/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ChatModel interface {
	AddMessage(msg, association, newerID string) error
	GetMessage(association, newerID string) ([]string, error)
	ClearMessage(association, newerID string) error
	IncrChattingCounter(association, newerID string) error
	DecrChattingCounter(association, newerID string) error
	GetChatCount(association, newerID string) (int, error)
	Close()
}

type Model interface {
	// 关闭数据库连接
	Close()
	// 终止操作，用于如事务的取消
	Abort()

	cStage() *mongo.Collection
	GetStageWithCondition(condition bson.M) ([]Stage, error)
	GetStageWithFlags(flag []string) ([]Stage, error)
	GetStageWithFlag(flag string) ([]Stage, error)
	GetStage(id primitive.ObjectID) (Stage, error)
	GetStageByHex(hex string) (Stage, error)
	GetStageWithFlagExclude(flag string) ([]Stage, error)

	SetPlayer(player Player) error
	GetPlayer(id string) (Player, error)

	cRecord() *mongo.Collection
	GetRecordByName(name string) (Record, bool, error)
	NewRecord(name string, playerID string) error
}

type model struct {
	dbTrait
	ctx   context.Context
	abort bool
}

func GetModel() Model {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	if config.C.Debug {
		ctx = context.Background()
	}

	ret := &model{
		dbTrait: getDBTx(ctx),
		ctx:     ctx,
		abort:   false,
	}

	return ret
}
