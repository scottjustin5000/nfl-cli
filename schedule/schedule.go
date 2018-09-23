package schedule

import (
	"fmt"
	"strconv"
	"time"
)

type Team struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	Points int    `json:"points"`
}

type Game struct {
	Id        string `json:"id"`
	Status    string `json:"status"`
	Scheduled string `json:"scheduled"`
	Home      Team   `json:"home"`
	Away      Team   `json:"away"`
	Display   string
	GameDate  string
}

func ordinal(x int) string {
	suffix := "th"
	switch x % 10 {
	case 1:
		suffix = "st"
	case 2:
		suffix = "nd"
	case 3:
		suffix = "rd"
	}
	return strconv.Itoa(x) + suffix
}

func (g *Game) SetDisplay() {
	g.Display = g.Away.Name + " @ " + g.Home.Name + "                            "
}
func (g *Game) SetGameDateDisplay() {
	str := g.Scheduled
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		fmt.Println(err)
	}
	loc, _ := time.LoadLocation("EST")
	t = t.In(loc)
	g.GameDate = t.Format("Monday") + "," + t.Format("Jan") + " " + ordinal(t.Day()) + "        "
}

type Week struct {
	Id       string `json:"id"`
	Sequence int    `json:"sequence"`
	Title    int    `json:"title"`
	Games    []Game `json:"games"`
}
type Schedule struct {
	Id    string `json:"id"`
	Year  int    `json:"year"`
	Weeks []Week `json:"weeks"`
}

func GetScheduleWeek(seasonType string, season int, week int, apiKey string) ScheduleWeek {
	sg := GetWeekGames(seasonType, season, week, apiKey)

	for i := 0; i < len(sg.Games); i++ {
		var g = &sg.Games[i]
		g.SetDisplay()
		g.SetGameDateDisplay()
	}

	return sg

}
