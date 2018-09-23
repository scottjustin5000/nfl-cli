package gamecenter

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/scottjustin5000/nfl-cli/ascii"
	"github.com/scottjustin5000/nfl-cli/schedule"
)

type ViewType int

const (
	PlayByPlay    ViewType = 0
	ScoringDrives ViewType = 1
	PlayerStats   ViewType = 2
	TeamStats     ViewType = 3
)

var tref *tview.Table
var app *tview.Application
var boxScore Boxscore
var pbpFeed PbpFeed
var grid *tview.Grid
var game schedule.Game
var selectedView ViewType

var ctx context.Context
var cancel context.CancelFunc

var rowCounter = 0
var maxRow = 0

var rightAlignText = func(text string) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignRight).
		SetText(text)
}

var centerAlignText = func(text string, color tcell.Color) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetTextColor(color).
		SetText(text)
}

var leftAlignText = func(text string) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetText(text)
}

var centerAlignColor = func(text string, color tcell.Color) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetTextColor(color).
		SetText(text)
}

var leftAlignColor = func(text string, color tcell.Color) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetTextColor(color).
		SetText(text)
}

var rightAlignColor = func(text string, color tcell.Color) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignRight).
		SetTextColor(color).
		SetText(text)
}

func loadGameCenter(game schedule.Game, apiKey string) {
	bxChan := make(chan Boxscore, 1)
	pbpChan := make(chan PbpFeed, 1)
	//go get data from sportradar api
	go func() {
		bxChan <- GetBoxScore(game.Id, apiKey)
	}()
	go func() {
		pbpChan <- loadPlayByPlay(game.Id, apiKey)
	}()
	boxScore = <-bxChan
	pbpFeed = <-pbpChan

	//create table object
	tref = tview.NewTable().SetBorders(true).SetBordersColor(tcell.ColorLightGrey)
	//make table selectable
	tref.SetSelectable(true, false)

	//build play-by-play view
	maxRow = BuildPbpTable(pbpFeed, tref)

	//get team names
	home := strings.Fields(game.Home.Name)
	away := strings.Fields(game.Away.Name)
	homeName := home[len(home)-1]
	awayName := away[len(away)-1]
	//create default title (only displayed when console is too small for ASCII ART)
	rawTitle := awayName + " v. " + homeName
	// build alternative ASCII ART title
	teamTitle := ascii.PrintS(strings.ToLower(rawTitle))
	//build boxscore header
	boxHeader := getBoxHeader(game, awayName, homeName)
	//build box, quarter-by-quarter scoring
	box := getBox(boxScore)
	//add menu selections
	flex := tview.NewFlex().AddItem(leftAlignText("(p) play-by-play"), 0, 1, false).
		AddItem(leftAlignText("(s) player stats"), 0, 1, false).
		AddItem(leftAlignText("(t) team stats"), 0, 1, false).
		AddItem(leftAlignText("(x) scoring"), 0, 1, false)
		//team title display and rules
	grid.
		SetRows(7, 1, 7, 1, 1, 0).
		SetColumns(30, 30, 30, 30).
		AddItem(centerAlignText(teamTitle, tcell.ColorWhite), 0, 0, 0, 0, 0, 0, false).
		AddItem(centerAlignText(teamTitle, tcell.ColorWhite), 0, 0, 1, 4, 0, 125, false)
		//add box summary and play-by-play title
	grid.
		AddItem(boxHeader, 1, 0, 1, 4, 0, 0, false).
		AddItem(box, 2, 0, 1, 4, 0, 0, false).
		AddItem(leftAlignText("Last Play: "+boxScore.LastEvent.Description), 3, 1, 1, 4, 0, 0, false).
		AddItem(flex, 4, 1, 1, 3, 0, 0, false).
		AddItem(tref, 5, 0, 1, 5, 0, 0, false)

	grid.SetInputCapture(onInput)
}

func Build(gamer schedule.Game, apiKey string) {
	game = gamer
	app = tview.NewApplication()
	grid = tview.NewGrid()
	loadGameCenter(gamer, apiKey)
	if pbpFeed.Status == "inprogress" {
		pollGame(apiKey)
	}

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}

}

func pollGame(apiKey string) {

	ctx, cancel = context.WithCancel(context.Background())
	forever := make(chan struct{})

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // if cancel() execute
				forever <- struct{}{}
				return
			default:
				time.Sleep(5000 * time.Millisecond)
				grid.Clear()
				loadGameCenter(game, apiKey)
				app.Draw()
				if pbpFeed.Status == "closed" {
					cancel()
				}
			}
		}
	}(ctx)

	<-forever
}

func onInput(event *tcell.EventKey) *tcell.EventKey {
	if strconv.QuoteRune(event.Rune()) == "'x'" {
		tref.Clear()
		mr := BuildScoringTable(boxScore.ScoringDrives, tref)
		rowCounter = 0
		tref.Select(rowCounter, 4)
		maxRow = mr
	} else if strconv.QuoteRune(event.Rune()) == "'p'" {
		tref.Clear()
		mr := BuildPbpTable(pbpFeed, tref)
		rowCounter = 0
		tref.Select(rowCounter, 4)
		maxRow = mr
	} else if strconv.QuoteRune(event.Rune()) == "'b'" {

	} else if strconv.QuoteRune(event.Rune()) == "'s'" {

	} else if strconv.QuoteRune(event.Rune()) == "'t'" {

	} else if event.Name() == "Down" && rowCounter < maxRow {
		rowCounter = rowCounter + 1
		tref.Select(rowCounter, 4)
	} else if event.Name() == "Up" && rowCounter > 0 {

		rowCounter = rowCounter - 1
		tref.Select(rowCounter, 4)
	}
	return event
}

func loadPlayByPlay(gameId string, apiKey string) PbpFeed {
	return GetPlayByPlay(gameId, apiKey)
}

func getBox(boxScore Boxscore) *tview.Flex {

	awayScore := ascii.PrintS(strconv.Itoa(boxScore.Summary.Away.Points))
	homeScore := ascii.PrintS(strconv.Itoa(boxScore.Summary.Home.Points))
	var periods = []string{""}
	if len(boxScore.Scoring) > 4 {
		for _, period := range boxScore.Scoring {
			if period.PeriodType == "overtime" {
				periods = append(periods, "OT"+strconv.Itoa(period.Number))
			} else {
				periods = append(periods, "Q"+strconv.Itoa(period.Number))
			}
		}
		periods = append(periods, "T")
	} else {
		periods = []string{"", "Q1", "Q2", "Q3", "Q4", "T"}
	}
	table3 := ascii.New(periods)

	awayBox := make(map[string]interface{})
	homeBox := make(map[string]interface{})

	for index, period := range periods {
		if index == 0 {
			awayBox[""] = boxScore.Summary.Away.Alias
			homeBox[""] = boxScore.Summary.Home.Alias
		} else if index == len(periods)-1 {
			awayBox["T"] = strconv.Itoa(boxScore.Summary.Away.Points)
			homeBox["T"] = strconv.Itoa(boxScore.Summary.Home.Points)
		} else if index <= len(boxScore.Scoring) {
			awayBox[period] = strconv.Itoa(boxScore.Scoring[index-1].AwayPoints)
			homeBox[period] = strconv.Itoa(boxScore.Scoring[index-1].HomePoints)
		} else {
			awayBox[period] = "-"
			homeBox[period] = "-"
		}
	}
	table3.AddRow(awayBox)
	table3.AddRow(homeBox)

	t := table3.Print()
	grid := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(rightAlignText(awayScore), 0, 1, false).
		AddItem(centerAlignColor(t, tcell.ColorYellow), 0, 2, false).
		AddItem(leftAlignText(homeScore), 0, 1, false)

	return grid
}

func getBoxHeader(game schedule.Game, awayName string, homeName string) *tview.Flex {
	ti, err := time.Parse(time.RFC3339, game.Scheduled)
	if err != nil {
		fmt.Println(err)
	}
	loc, _ := time.LoadLocation("EST")
	ti = ti.In(loc)
	gd := ti.Format("Monday") + "," + ti.Format("Jan") + " " + strconv.Itoa(ti.Day())
	grid := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(rightAlignColor(awayName, getTeamColor(game.Away.Alias)), 0, 1, false).
		AddItem(centerAlignText(gd, tcell.ColorWhite), 0, 2, false).
		AddItem(leftAlignColor(homeName, getTeamColor(game.Home.Alias)), 0, 1, false)
	return grid
}
