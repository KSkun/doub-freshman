package model

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Major     string   `json:"major"`
	Selection []string `json:"selection"`
	Next      string   `json:"next"`
	Flag      []Flag   `json:"flag"`
	LeftRound int      `json:"left_round"`
}

const keyPlayer = "player:%s"

func (m *model) SetPlayer(player Player) error {
	doc, err := json.Marshal(player)
	if err != nil {
		return err
	}
	result := redisClient.Set(fmt.Sprintf(keyPlayer, player.ID), string(doc), 0)
	return result.Err()
}

func (m *model) GetPlayer(id string) (Player, error) {
	player := Player{}
	doc, err := redisClient.Get(fmt.Sprintf(keyPlayer, id)).Result()
	if err != nil {
		return player, err
	}
	err = json.Unmarshal([]byte(doc), &player)
	return player, err
}
