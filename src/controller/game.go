package controller

import (
	"github.com/KSkun/doub-freshman/constant"
	"github.com/KSkun/doub-freshman/model"
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
			return checkCompCond(_flag, cond)
		}
	}
	return false
}

func checkStageCond(player model.Player, stage model.Stage) bool {
	for _, cond := range stage.EnterCond {
		if !checkSingleCond(player, cond) {
			return false
		}
	}
	return true
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
