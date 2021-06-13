package controller

import (
	"errors"
	"fmt"
	"github.com/KSkun/doub-freshman/controller/param"
	"github.com/KSkun/doub-freshman/model"
	"github.com/KSkun/doub-freshman/util/context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func makeRspNextStage(stage model.Stage) param.RspNextStage {
	var option []string
	for _, _option := range stage.Option {
		option = append(option, _option.Text)
	}
	return param.RspNextStage{
		ID:     stage.ID.Hex(),
		Text:   stage.Text,
		Title:  stage.Title,
		Option: option,
		Delay:  stage.Delay,
	}
}

func makeRspGameSync(player model.Player) (param.RspGameSync, error) {
	rsp := param.RspGameSync{}
	m := model.GetModel()
	defer m.Close()
	if player.Next != "" {
		t, err := m.GetStageByHex(player.Next)
		if err != nil {
			return rsp, err
		}
		rsp.Next = makeRspNextStage(t)
	}
	rsp.Selection = []param.RspNextStage{}
	for _, hex := range player.Selection {
		t, err := m.GetStageByHex(hex)
		if err != nil {
			return rsp, err
		}
		rsp.Selection = append(rsp.Selection, makeRspNextStage(t))
	}
	rsp.Flag = []string{}
	for _, flag := range player.Flag {
		flagStr := flag.Text
		if flag.Value != 0 {
			flagStr += fmt.Sprintf("%.1f", flag.Value)
		}
		rsp.Flag = append(rsp.Flag, flagStr)
	}
	return rsp, nil
}

func HandlerNewGame(ctx echo.Context) error {
	req := param.ReqNewGame{}
	if err := ctx.Bind(&req); err != nil {
		return context.Error(ctx, http.StatusBadRequest, "invalid request", err)
	}
	if err := ctx.Validate(req); err != nil {
		return context.Error(ctx, http.StatusBadRequest, "invalid request", err)
	}

	m := model.GetModel()
	defer m.Close()
	_, found, err := m.GetRecordByName(req.Name)
	if err != nil {
		return context.Error(ctx, http.StatusInternalServerError, "internal error", err)
	}
	if found {
		return context.Error(ctx, http.StatusBadRequest, "name used", errors.New("name used"))
	}

	player, err := newPlayer(req.Name, req.Major)
	if err != nil {
		return context.Error(ctx, http.StatusInternalServerError, "internal error", err)
	}
	player.ID = primitive.NewObjectID().Hex()
	err = m.SetPlayer(player)
	if err != nil {
		return context.Error(ctx, http.StatusInternalServerError, "internal error", err)
	}
	err = m.NewRecord(req.Name, player.ID)
	if err != nil {
		return context.Error(ctx, http.StatusInternalServerError, "internal error", err)
	}
	_rsp, err := makeRspGameSync(player)
	if err != nil {
		return context.Error(ctx, http.StatusInternalServerError, "internal error", err)
	}
	rsp := param.RspNewGame{RspGameSync: _rsp, ID: player.ID}
	return context.Success(ctx, rsp)
}