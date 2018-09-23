package gamecenter

type Boxscore struct {
	Id            string         `json:"id"`
	Status        string         `json:"status"`
	Scheduled     string         `json:"scheduled"`
	Weather       string         `json:"weather"`
	Clock         string         `json:"clock"`
	Quarter       int            `json:"quarter"`
	Summary       Summary        `json:"summary"`
	Situation     Situation      `json:"situation"`
	LastEvent     LastEvent      `json:"last_event"`
	Scoring       []Scoring      `json:"scoring"`
	ScoringDrives []ScoringDrive `json:"scoring_drives"`
}

type Team struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	Points int    `json:"points"`
}

type Summary struct {
	Season Season  `json:"season"`
	Week   BoxWeek `json:"week"`
	Venue  Venue   `json:"venue"`
	Home   Team    `json:"home"`
	Away   Team    `json:"away"`
}

type Season struct {
	Year       int    `json:"year"`
	SeasonType string `json:"type"`
}

type BoxWeek struct {
	Sequence int    `json:"sequence"`
	Title    string `json:"title"`
}

type Venue struct {
	Name  string `json:"name"`
	City  string `json:"city"`
	State string `json:"state"`
}

type Situation struct {
	Clock      string     `json:"clock"`
	Down       int        `json:"down"`
	Yfd        int        `json:"yfd"`
	Possession Possession `json:"possession"`
	Location   Location   `json:"location"`
}

type Possession struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

type Location struct {
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	Yardline int    `json:"yardline"`
}

type LastEvent struct {
	EventType   string `json:"type"`
	Clock       string `json:"clock"`
	Description string `json:"description"`
}

type Scoring struct {
	PeriodType string `json:"period_type"`
	Number     int    `json:"number"`
	Sequence   int    `json:"sequence"`
	HomePoints int    `json:"home_points"`
	AwayPoints int    `json:"away_points"`
}
type Play struct {
	Clock          string    `json:"clock"`
	Description    string    `json:"description"`
	StartSituation Situation `json:"start_situation"`
}
type Quarter struct {
	Number   int `json:"number"`
	Sequence int `json:"sequence"`
}
type ScoringDrive struct {
	Quarter     Quarter `json:"quarter"`
	Clock       string  `json:"clock"`
	Duration    string  `json:"duration"`
	StartReason string  `json:"start_reason"`
	EndReason   string  `json:"end_reason"`
	PlayCount   int     `json:"play_count"`
	Gain        int     `json:"gain"`
	Plays       []Play  `json:"plays"`
}
