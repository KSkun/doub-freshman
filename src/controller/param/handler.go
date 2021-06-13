package param

type ReqNewGame struct {
	Name  string `json:"name"`
	Major string `json:"major"`
}

type RspPlayerInfo struct {
	Name  string `json:"name"`
	Major string `json:"major"`
}

type RspNextStage struct {
	ID     string   `json:"id"`
	Text   string   `json:"text"`
	Option []string `json:"option"`
}

type RspGameSync struct {
	Next      RspNextStage   `json:"next"`
	Selection []RspNextStage `json:"selection"`
	Flag      []string       `json:"flag"`
}

type RspNewGame struct {
	RspGameSync
	ID string `json:"id"`
}
