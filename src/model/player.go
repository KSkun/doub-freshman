package model

type Player struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Selection []string `json:"selection"`
	Next      string   `json:"next"`
	Flag      []Flag   `json:"flag"`
	LeftRound int      `json:"left_round"`
}
