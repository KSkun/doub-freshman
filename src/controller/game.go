package controller

import (
	"errors"
	"github.com/KSkun/doub-freshman/constant"
	"github.com/KSkun/doub-freshman/controller/param"
	"github.com/KSkun/doub-freshman/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
)

func strArrayToSet(arr []string) map[string]bool {
	result := map[string]bool{}
	for _, ele := range arr {
		result[ele] = true
	}
	return result
}

func strSetToArray(_map map[string]bool) []string {
	var result []string
	for ele, _ := range _map {
		result = append(result, ele)
	}
	return result
}

func stageSetToArray(_map map[string]model.Stage) []string {
	var result []string
	for ele, _ := range _map {
		result = append(result, ele)
	}
	return result
}

func deleteFlagFromSlice(slice []model.Flag, flag model.Flag) []model.Flag {
	idx := -1
	for i, _flag := range slice {
		if _flag.Text == flag.Text {
			idx = i
			break
		}
	}
	if idx == -1 {
		return slice
	}
	return append(slice[:idx], slice[idx+1:]...)
}

func checkCompCond(flag model.Flag, cond model.Condition) bool {
	switch cond.Op {
	case "lt":
		return flag.Value < cond.Value
	case "lte":
		return flag.Value <= cond.Value
	case "gt":
		return flag.Value > cond.Value
	case "gte":
		return flag.Value >= cond.Value
	case "eq":
		return flag.Value == cond.Value
	case "neq":
		return flag.Value != cond.Value
	default:
	}
	return true
}

func checkSingleCond(player model.Player, cond model.Condition) bool {
	for _, _flag := range player.Flag {
		if _flag.Text == cond.Flag {
			if cond.Op == "" {
				return true
			}
			if cond.Op == "exclude" {
				return false
			}
			if cond.Op == "prob" {
				r := rand.Float64()
				return r < cond.Value
			}
			return checkCompCond(_flag, cond)
		}
	}
	return false
}

func checkCondList(player model.Player, list []model.Condition) bool {
	for _, cond := range list {
		if !checkSingleCond(player, cond) {
			return false
		}
	}
	return true
}

func checkStageCond(player model.Player, stage model.Stage) bool {
	return checkCondList(player, stage.EnterCond)
}

func checkStageExclude(stage model.Stage, flag model.Flag) bool {
	for _, cond := range stage.EnterCond {
		if cond.Op == "exclude" && cond.Flag == flag.Text {
			return true
		}
	}
	return false
}

func addFlag(_player model.Player, flag model.Flag) (model.Player, error) {
	for _, _flag := range _player.Flag {
		if _flag.Text == flag.Text {
			return _player, nil
		}
	}
	player := _player
	player.Flag = append(player.Flag, flag)
	m := model.GetModel()
	defer m.Close()
	stageSet := map[string]model.Stage{}
	for _, stageID := range player.Selection {
		stage, err := m.GetStageByHex(stageID)
		if err != nil {
			return _player, err
		}
		stageSet[stageID] = stage
	}

	// check new available stages
	avaStage, err := m.GetStageWithFlag(flag.Text)
	if err != nil {
		return player, err
	}
	for _, stage := range avaStage {
		if checkStageCond(player, stage) {
			stageSet[stage.ID.Hex()] = stage
		}
	}
	// check old stages excluded
	for _, stage := range stageSet {
		if checkStageExclude(stage, flag) {
			delete(stageSet, stage.ID.Hex())
		}
	}

	player.Selection = stageSetToArray(stageSet)
	return player, nil
}

func deleteFlag(_player model.Player, flag model.Flag) (model.Player, error) {
	found := false
	for _, _flag := range _player.Flag {
		if _flag.Text == flag.Text {
			found = true
			break
		}
	}
	if !found {
		return _player, nil
	}
	player := _player
	player.Flag = deleteFlagFromSlice(player.Flag, flag)
	m := model.GetModel()
	defer m.Close()
	stageSet := map[string]model.Stage{}
	for _, stageID := range player.Selection {
		stage, err := m.GetStageByHex(stageID)
		if err != nil {
			return _player, err
		}
		stageSet[stageID] = stage
	}

	// check unsatisfied stages
	for _, stage := range stageSet {
		if !checkStageCond(player, stage) {
			delete(stageSet, stage.ID.Hex())
		}
	}
	// check if excluded stages can be added
	excStage, err := m.GetStageWithFlagExclude(flag.Text)
	if err != nil {
		return _player, err
	}
	for _, stage := range excStage {
		if checkStageCond(player, stage) {
			stageSet[stage.ID.Hex()] = stage
		}
	}

	player.Selection = stageSetToArray(stageSet)
	return player, nil
}

func increaseFlag(_player model.Player, flag model.Flag, inc float64) (model.Player, error) {
	found := false
	idx := -1
	for i, _flag := range _player.Flag {
		if _flag.Text == flag.Text {
			found = true
			idx = i
			break
		}
	}
	if !found {
		return _player, nil
	}
	player := _player
	// update flag
	player.Flag[idx].Value += inc
	m := model.GetModel()
	defer m.Close()
	stageSet := map[string]model.Stage{}
	for _, stageID := range player.Selection {
		stage, err := m.GetStageByHex(stageID)
		if err != nil {
			return _player, err
		}
		stageSet[stageID] = stage
	}

	// check new available stages
	avaStage, err := m.GetStageWithFlag(flag.Text)
	if err != nil {
		return player, err
	}
	for _, stage := range avaStage {
		if checkStageCond(player, stage) {
			stageSet[stage.ID.Hex()] = stage
		}
	}
	// check old stages unsatisfied
	for _, stage := range stageSet {
		if !checkStageCond(player, stage) {
			delete(stageSet, stage.ID.Hex())
		}
	}

	player.Selection = stageSetToArray(stageSet)
	return player, nil
}

func newPlayer(name string, major string) (model.Player, error) {
	gpa := rand.Int()%(95-80+1) + 80 // random initial gpa in [80, 95]
	player := model.Player{
		Name:      name,
		Major:     major,
		Selection: []string{},
		Next:      "",
		Flag:      []model.Flag{{Text: "加权", Value: float64(gpa)}},
		LeftRound: constant.RoundHoliday,
	}
	player, err := addFlag(player, model.Flag{Text: "假期"})
	if err != nil {
		return player, err
	}
	player, err = addFlag(player, model.Flag{Text: "萌新"})
	if err != nil {
		return player, err
	}
	player, err = addFlag(player, model.Flag{Text: major})
	if err != nil {
		return player, err
	}
	return player, nil
}

func selectOption(_player model.Player, _option int) (model.Player, error) {
	m := model.GetModel()
	defer m.Close()
	stage, err := m.GetStageByHex(_player.NowStage)
	if err != nil {
		return _player, err
	}
	if _option < 0 || _option >= len(stage.Option) {
		return _player, errors.New("option index out of bound")
	}
	option := stage.Option[_option]
	nextID := primitive.ObjectID{}
	if checkCondList(_player, option.Condition) {
		nextID = option.Success.Next
	} else {
		nextID = option.Failed.Next
	}
	next, err := m.GetStage(nextID)
	if err != nil {
		return _player, err
	}
	player := _player
	if next.Continue {
		player.Next = nextID.Hex()
	} else {
		player.Selection = append(player.Selection, nextID.Hex())
	}
	return player, err
}

func applySingleEvent(player model.Player, event model.Event) (model.Player, []param.FlagDiff, error) {
	switch event.Type {
	case "flag":
		flag := event.Value["flag"].(string)
		del, delOK := event.Value["delete"].(bool)
		inc, incOK := event.Value["inc"].(float64)
		if delOK && del {
			_player, err := deleteFlag(player, model.Flag{Text: flag})
			return _player, []param.FlagDiff{{Type: "del", Flag: flag}}, err
		} else if incOK && inc != 0 {
			_player, err := increaseFlag(player, model.Flag{Text: flag}, inc)
			return _player, []param.FlagDiff{{Type: "inc", Flag: flag, Value: inc}}, err
		} else {
			_player, err := addFlag(player, model.Flag{Text: flag})
			return _player, []param.FlagDiff{{Type: "add", Flag: flag}}, err
		}
	case "death":
		player.Dead = true
		return player, nil, nil
	case "extend":
		ext, extOK := event.Value["extend"].(int)
		if !extOK {
			return player, nil, errors.New("extend value not found")
		}
		player.LeftRound += ext
		return player, nil, nil
	default:
	}
	return player, nil, nil
}

func applyEventList(player model.Player, event []model.Event) (model.Player, []param.FlagDiff, error) {
	var err error
	var _flagDiff []param.FlagDiff
	for _, _event := range event {
		var __flagDiff []param.FlagDiff
		player, __flagDiff, err = applySingleEvent(player, _event)
		if err != nil {
			return player, nil, err
		}
		_flagDiff = append(_flagDiff, __flagDiff...)
	}
	return player, _flagDiff, nil
}

func switchPhase(player model.Player) (model.Player, []param.FlagDiff, error) {
	var _flagDiff []param.FlagDiff
	idx := -1
	for i, flag := range player.Flag {
		if flag.Text == "假期" || flag.Text == "在校" {
			idx = i
			break
		}
	}
	next := ""
	if player.Flag[idx].Text == "假期" {
		next = "在校"
		player.LeftRound = constant.RoundDaily
	} else {
		next = "假期"
		player.LeftRound = constant.RoundHoliday
	}

	_flagDiff = append(_flagDiff, param.FlagDiff{
		Type: "del",
		Flag: player.Flag[idx].Text,
	})
	_flagDiff = append(_flagDiff, param.FlagDiff{
		Type: "add",
		Flag: next,
	})

	player, err := deleteFlag(player, player.Flag[idx])
	if err != nil {
		return player, nil, err
	}
	player, err = addFlag(player, model.Flag{Text: next})
	return player, _flagDiff, err
}

func selectStage(_player model.Player, stageID string) (model.Player, []param.FlagDiff, error) {
	var _flagDiff []param.FlagDiff
	found := false
	if _player.Next != "" {
		if _player.Next == stageID {
			found = true
		} else {
			return _player, nil, errors.New("continuing stage exists")
		}
	}else {
		for _, stage := range _player.Selection {
			if stage == stageID {
				found = true
				break
			}
		}
	}
	if !found {
		return _player, nil, errors.New("invalid stage")
	}

	m := model.GetModel()
	defer m.Close()
	stage, err := m.GetStageByHex(_player.NowStage)
	if err != nil {
		return _player, nil, err
	}
	player := _player
	player.NowStage = stageID
	player, __flagDiff, err := applyEventList(player, stage.Event)
	if err != nil {
		return _player, nil, err
	}
	_flagDiff = append(_flagDiff, __flagDiff...)
	player.LeftRound--
	// if a continuing stage is waiting to be trigger, do not switch phase
	if player.LeftRound <= 0 && player.Next == "" {
		var __flagDiff []param.FlagDiff
		player, __flagDiff, err = switchPhase(player)
		if err != nil {
			return _player, nil, err
		}
		_flagDiff = append(_flagDiff, __flagDiff...)
	}
	return player, _flagDiff, nil
}
