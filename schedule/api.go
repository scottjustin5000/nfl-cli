package schedule

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// ScheduleWeek structure for NFL schedule week
type ScheduleWeek struct {
	Season     int
	SeasonType string
	WeekNumber int
	Week       int
	Games      []Game
}

func getSchedule(seasonType string, season int, apiKey string) Schedule {
	rs, err := http.Get("http://api.sportradar.us/nfl-radar360/games/" + strconv.Itoa(season) + "/" + seasonType + "/schedule.json?api_key=" + apiKey)
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	var data Schedule
	json.Unmarshal(body, &data)
	return data
}

// GetWeekGames gets games for the given NFL schedule week
func GetWeekGames(seasonType string, season int, week int, apiKey string) ScheduleWeek {
	var games []Game
	schedule := getSchedule(seasonType, season, apiKey)
	for _, schWeek := range schedule.Weeks {
		if schWeek.Sequence == week {
			games = schWeek.Games
			break
		}
	}
	return ScheduleWeek{Season: season, SeasonType: seasonType, WeekNumber: week, Games: games}
}
