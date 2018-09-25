package gamecenter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetBoxScore get boxscore feed for a given game
func GetBoxScore(gameId string, apiKey string) Boxscore {
	rs, err := http.Get("http://api.sportradar.us/nfl-radar360/games/" + gameId + "/boxscore.json?api_key=" + apiKey)
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}
	var data Boxscore
	json.Unmarshal(body, &data)
	return data
}

// GetPlayByPlay gets play-by-play feed for a given game
func GetPlayByPlay(gameId string, apiKey string) PbpFeed {
	rs, err := http.Get("http://api.sportradar.us/nfl-radar360/games/" + gameId + "/pbp.json?api_key=" + apiKey)
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	var data PbpFeed
	json.Unmarshal(body, &data)
	return data
}

// GetGameSummary make call to SR game summary feed
func GetGameSummary(gameId string, apiKey string) GameSummary {
	rs, err := http.Get("http://api.sportradar.us/nfl-radar360/games/" + gameId + "/statistics.json?api_key=" + apiKey)
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	var data GameSummary
	json.Unmarshal(body, &data)
	return data
}
