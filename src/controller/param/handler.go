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
	End       bool           `json:"end"`
	Result    string         `json:"result"`
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

type RspSyncWithDiff struct {
	RspGameSync
	FlagDiff []FlagDiff `json:"flag_diff"`
}

type RspSyncWithResult struct {
	RspSyncWithDiff
	Result string `json:"result"`
}

type ReqFlag struct {
	Text  string  `json:"text"`
	Value float64 `json:"value"`
	Hide  bool    `json:"hide"`
}

type ReqCondition struct {
	Flag  string  `json:"flag"`
	Op    string  `json:"op"`
	Value float64 `json:"value"`
}

type ReqOptionBranch struct {
	Next string `json:"next"`
	Text string `json:"text"`
}

type ReqOption struct {
	Text      string          `json:"text"`
	Success   ReqOptionBranch `json:"success"`
	Failed    ReqOptionBranch `json:"failed"`
	Condition []ReqCondition  `json:"condition"`
}

type ReqEvent struct {
	Type  string                 `json:"type"`
	Value map[string]interface{} `json:"value"`
}

type ReqStage struct {
	ID        string         `json:"_id"`
	Title     string         `json:"title"`
	Text      string         `json:"text"`
	Dead      bool           `json:"dead"`
	EnterCond []ReqCondition `json:"enter_cond"`
	Option    []ReqOption    `json:"option"`
	Event     []ReqEvent     `json:"event"`
	Tag       string         `json:"tag"`

	Continue bool `json:"continue"`
	Delay    int  `json:"delay"`
}
