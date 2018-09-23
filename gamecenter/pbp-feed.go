package gamecenter

type PbpFeed struct {
	Status  string   `json:"status"`
	Periods []Period `json:"periods"`
}

type Period struct {
	Number int   `json:"number"`
	Pbp    []Pbp `json:"pbp"`
}

type Pbp struct {
	Type      string  `json:"type"`
	Duration  string  `json:"duration"`
	EndReason string  `json:"end_reason"`
	Events    []Event `json:"events"`
	Clock     string  `json:"clock"`
	PlayCount int     `json:"play_count"`
	Gain      int     `json:"gain"`
}

type Event struct {
	Clock          string    `json:"clock"`
	Description    string    `json:"description"`
	StartSituation Situation `json:"start_situation"`
}
