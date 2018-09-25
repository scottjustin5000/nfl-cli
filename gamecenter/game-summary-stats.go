package gamecenter

type GameSummary struct {
	Id         string      `json:"id"`
	Status     string      `json:"status"`
	Statistics TeamSummary `json:"statistics"`
}
type TeamStatsSummary struct {
	Id             string      `json:"id"`
	Name           string      `json:"name"`
	Alias          string      `json:"alias"`
	Summary        StatSummary `json:"summary"`
	Passing        StatSummary `json:"passing"`
	Rushing        StatSummary `json:"rushing"`
	Receiving      StatSummary `json:"receiving"`
	Fumbles        StatSummary `json:"fumbles"`
	Kickoffs       StatSummary `json:"kickoffs"`
	KickoffReturns StatSummary `json:"kick_returns"`
	Punts          StatSummary `json:"punts"`
	PuntReturns    StatSummary `json:"punt_returns"`
	Penalties      StatSummary `json:"penalties"`
	FieldGoals     StatSummary `json:"field_goals"`
	Defense        StatSummary `json:"defense"`
}

type TeamSummary struct {
	Home TeamStatsSummary `json:"home"`
	Away TeamStatsSummary `json:"away"`
}

type StatSummary struct {
	Totals  map[string]interface{}   `json:"totals"`
	Players []map[string]interface{} `json:"players"`
}
