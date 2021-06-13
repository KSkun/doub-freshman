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
	Title  string   `json:"title"`
	Text   string   `json:"text"`
	Option []string `json:"option"`
	Delay  int      `json:"delay"`
}

type RspGameSync struct {
	Next      RspNextStage   `json:"next"`
	Selection []RspNextStage `json:"selection"`
	Flag      []string       `json:"flag"`
	Dead      bool           `json:"dead"`
}

type RspNewGame struct {
	RspGameSync
	ID string `json:"id"`
}

type ReqSelectOption struct {
	Player string `json:"player"`
	Option int    `json:"option"`
}

type ReqSelectStage struct {
	Player string `json:"player"`
	Stage  string `json:"stage"`
}

type FlagDiff struct {
	Type  string  `json:"type"`
	Flag  string  `json:"flag"`
	Value float64 `json:"value"`
}

type RspSelectStage struct {
	RspGameSync
	FlagDiff []FlagDiff `json:"flag_diff"`
}
