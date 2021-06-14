package controller

import (
	"errors"
	"fmt"
	"github.com/KSkun/doub-freshman/constant"
	"github.com/KSkun/doub-freshman/controller/param"
	"github.com/KSkun/doub-freshman/model"
	"github.com/KSkun/doub-freshman/util/context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

func makeRspNextStage(player model.Player, stage model.Stage) param.RspNextStage {
	var option []string
	for _, _option := range stage.Option {
		text := strings.ReplaceAll(_option.Text, constant.NamePlaceholder, player.Name)
		option = append(option, text)
	}
	text := strings.ReplaceAll(stage.Text, constant.NamePlaceholder, player.Name)
	title := strings.ReplaceAll(stage.Title, constant.NamePlaceholder, player.Name)
	return param.RspNextStage{
		ID:     stage.ID.Hex(),
		Text:   text,
		Title:  title,
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
		rsp.Next = makeRspNextStage(player, t)
	}
	rsp.Selection = []param.RspNextStage{}
	for _, hex := range player.Selection {
		t, err := m.GetStageByHex(hex)
		if err != nil {
			return rsp, err
		}
		rsp.Selection = append(rsp.Selection, makeRspNextStage(player, t))
	}
	rsp.Flag = []string{}
	for _, flag := range player.Flag {
		// hide flag
		if flag.Hide {
			continue
		}
		flagStr := flag.Text
		if flag.Value != 0 {
			flagStr += fmt.Sprintf("%.1f", flag.Value)
		}
		rsp.Flag = append(rsp.Flag, flagStr)
	}
	rsp.Dead = player.Dead
	rsp.End = player.End
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

func HandlerSelectOption(ctx echo.Context) error {
	req := param.ReqSelectOption{}
	if err := ctx.Bind(&req); err != nil {
		return context.Error(ctx, http.StatusBadRequest, "invalid request", err)
	}
	if err := ctx.Validate(req); err != nil {
		return context.Error(ctx, http.StatusBadRequest, "invalid request", err)
	}

	m := model.GetModel()
	defer m.Close()
	player, err := m.GetPlayer(req.Player)
	if err != nil {
		return context.Error(ctx, http.StatusInternalServerError, "internal error", err)
	}
	player, _flagDiff, result, err := selectOption(player, req.Option)
	if err != nil {
		return context.Error(ctx, http.StatusInternalServerError, "internal error", err)
	}
	err = m.SetPlayer(player)
	if err != nil {
		return context.Error(ctx, http.StatusInternalServerError, "internal error", err)
	}
	_rsp, err := makeRspGameSync(player)
	if err != nil {
		return context.Error(ctx, http.StatusInternalServerError, "internal error", err)
	}
	return context.Success(ctx, param.RspSyncWithResult{
		RspSyncWithDiff: param.RspSyncWithDiff{RspGameSync: _rsp,
			FlagDiff: _flagDiff,
		},
		Result: result,
	})
}

func HandlerSelectStage(ctx echo.Context) error {
	req := param.ReqSelectStage{}
	if err := ctx.Bind(&req); err != nil {
		return context.Error(ctx, http.StatusBadRequest, "invalid request", err)
	}
	if err := ctx.Validate(req); err != nil {
		return context.Error(ctx, http.StatusBadRequest, "invalid request", err)
	}

	m := model.GetModel()
	defer m.Close()
	player, err := m.GetPlayer(req.Player)
	if err != nil {
		return context.Error(ctx, http.StatusInternalServerError, "internal error", err)
	}
	player, _flagDiff, err := selectStage(player, req.Stage)
	if err != nil {
		return context.Error(ctx, http.StatusInternalServerError, "internal error", err)
	}
	err = m.SetPlayer(player)
	if err != nil {
		return context.Error(ctx, http.StatusInternalServerError, "internal error", err)
	}
	_rsp, err := makeRspGameSync(player)
	if err != nil {
		return context.Error(ctx, http.StatusInternalServerError, "internal error", err)
	}
	return context.Success(ctx, param.RspSyncWithDiff{
		RspGameSync: _rsp,
		FlagDiff:    _flagDiff,
	})
}
