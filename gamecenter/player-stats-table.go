package gamecenter

import (
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func floatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
func floatToIntToString(input float64) string {
	return strconv.Itoa(int(input))
}

func buildPassing(table *tview.Table, summary GameSummary, rowCounter int) int {
	table.SetCell(rowCounter, 0,
		tview.NewTableCell("Passing").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 1,
		tview.NewTableCell("CP/ATT").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 2,
		tview.NewTableCell("YDS").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 3,
		tview.NewTableCell("TD").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 4,
		tview.NewTableCell("INT").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 5,
		tview.NewTableCell("Passing").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 6,
		tview.NewTableCell("CP/ATT").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 7,
		tview.NewTableCell("YDS").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 8,
		tview.NewTableCell("TD").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 9,
		tview.NewTableCell("INT").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	rowCounter++
	statCounter := 0
	for _, player := range summary.Statistics.Away.Passing.Players {
		//for _, v := range player {
		table.SetCell(rowCounter, 0,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(rowCounter, 1,
			tview.NewTableCell(floatToIntToString(player["completions"].(float64))+"/"+floatToIntToString(player["attempts"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 2,
			tview.NewTableCell(floatToIntToString(player["yards"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 3,
			tview.NewTableCell(floatToIntToString(player["touchdowns"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 4,
			tview.NewTableCell(floatToIntToString(player["interceptions"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		rowCounter++
		statCounter++
		//	}
	}
	homeCounter := rowCounter - statCounter
	for _, player := range summary.Statistics.Home.Passing.Players {
		//for _, v := range player {
		table.SetCell(homeCounter, 5,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(homeCounter, 6,
			tview.NewTableCell(floatToIntToString(player["completions"].(float64))+"/"+floatToIntToString(player["attempts"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 7,
			tview.NewTableCell(floatToIntToString(player["yards"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 8,
			tview.NewTableCell(floatToIntToString(player["touchdowns"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 9,
			tview.NewTableCell(floatToIntToString(player["interceptions"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		homeCounter++
		if homeCounter > rowCounter {
			rowCounter++
		}
	}
	return rowCounter
}

func buildRushing(table *tview.Table, summary GameSummary, rowCounter int) int {
	table.SetCell(rowCounter, 0,
		tview.NewTableCell("Rushing").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 1,
		tview.NewTableCell("ATT").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 2,
		tview.NewTableCell("YDS").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 3,
		tview.NewTableCell("TD").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 4,
		tview.NewTableCell("LG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 5,
		tview.NewTableCell("Rushing").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 6,
		tview.NewTableCell("ATT").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 7,
		tview.NewTableCell("YDS").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 8,
		tview.NewTableCell("TD").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 9,
		tview.NewTableCell("LG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	rowCounter++
	statCounter := 0
	for _, player := range summary.Statistics.Away.Rushing.Players {
		//for _, v := range player {
		table.SetCell(rowCounter, 0,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(rowCounter, 1,
			tview.NewTableCell(floatToIntToString(player["attempts"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 2,
			tview.NewTableCell(floatToIntToString(player["yards"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 3,
			tview.NewTableCell(floatToIntToString(player["touchdowns"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 4,
			tview.NewTableCell(floatToIntToString(player["longest"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		rowCounter++
		statCounter++
		//	}
	}
	homeCounter := rowCounter - statCounter
	for _, player := range summary.Statistics.Home.Rushing.Players {
		//for _, v := range player {
		table.SetCell(homeCounter, 5,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(homeCounter, 6,
			tview.NewTableCell(floatToIntToString(player["attempts"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 7,
			tview.NewTableCell(floatToIntToString(player["yards"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 8,
			tview.NewTableCell(floatToIntToString(player["touchdowns"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 9,
			tview.NewTableCell(floatToIntToString(player["longest"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		homeCounter++
		if homeCounter > rowCounter {
			rowCounter++
		}
	}
	return rowCounter
}

func buildReceiving(table *tview.Table, summary GameSummary, rowCounter int) int {
	table.SetCell(rowCounter, 0,
		tview.NewTableCell("Receiving").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 1,
		tview.NewTableCell("T/REC").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 2,
		tview.NewTableCell("YDS").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 3,
		tview.NewTableCell("TD").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 4,
		tview.NewTableCell("LG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 5,
		tview.NewTableCell("Receiving").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 6,
		tview.NewTableCell("T/REC").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 7,
		tview.NewTableCell("YDS").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 8,
		tview.NewTableCell("TD").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 9,
		tview.NewTableCell("LG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	rowCounter++
	statCounter := 0
	for _, player := range summary.Statistics.Away.Receiving.Players {
		table.SetCell(rowCounter, 0,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(rowCounter, 1,
			tview.NewTableCell(floatToIntToString(player["targets"].(float64))+"/"+floatToIntToString(player["receptions"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 2,
			tview.NewTableCell(floatToIntToString(player["yards"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 3,
			tview.NewTableCell(floatToIntToString(player["touchdowns"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 4,
			tview.NewTableCell(floatToIntToString(player["longest"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		rowCounter++
		statCounter++
		//	}
	}
	homeCounter := rowCounter - statCounter
	for _, player := range summary.Statistics.Home.Receiving.Players {
		table.SetCell(homeCounter, 5,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(homeCounter, 6,
			tview.NewTableCell(floatToIntToString(player["targets"].(float64))+"/"+floatToIntToString(player["receptions"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 7,
			tview.NewTableCell(floatToIntToString(player["yards"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 8,
			tview.NewTableCell(floatToIntToString(player["touchdowns"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 9,
			tview.NewTableCell(floatToIntToString(player["longest"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		homeCounter++
		if homeCounter > rowCounter {
			rowCounter++
		}
	}
	return rowCounter
}
func buildFumbles(table *tview.Table, summary GameSummary, rowCounter int) int {
	table.SetCell(rowCounter, 0,
		tview.NewTableCell("Fumbles").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 1,
		tview.NewTableCell("FUM").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 2,
		tview.NewTableCell("LOST").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 3,
		tview.NewTableCell("REC").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 4,
		tview.NewTableCell("YDS").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 5,
		tview.NewTableCell("Fumbles").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 6,
		tview.NewTableCell("FUM").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 7,
		tview.NewTableCell("LOST").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 8,
		tview.NewTableCell("REC").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 9,
		tview.NewTableCell("YDS").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	rowCounter++
	statCounter := 0
	for _, player := range summary.Statistics.Away.Fumbles.Players {
		table.SetCell(rowCounter, 0,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(rowCounter, 1,
			tview.NewTableCell(floatToIntToString(player["fumbles"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 2,
			tview.NewTableCell(floatToIntToString(player["lost_fumbles"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		var recovery string
		if int(player["own_rec"].(float64)) > 0 {
			recovery = floatToIntToString(player["own_rec"].(float64))
		} else {
			recovery = floatToIntToString(player["opp_rec"].(float64))
		}
		table.SetCell(rowCounter, 3,
			tview.NewTableCell(recovery).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		var recoveryYards string
		if int(player["own_rec_yards"].(float64)) > 0 {
			recoveryYards = floatToIntToString(player["own_rec_yards"].(float64))
		} else {
			recoveryYards = floatToIntToString(player["opp_rec_yards"].(float64))
		}

		table.SetCell(rowCounter, 4,
			tview.NewTableCell(recoveryYards).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		rowCounter++
		statCounter++
		//	}
	}
	homeCounter := rowCounter - statCounter
	for _, player := range summary.Statistics.Home.Fumbles.Players {
		table.SetCell(homeCounter, 5,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(homeCounter, 6,
			tview.NewTableCell(floatToIntToString(player["fumbles"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 7,
			tview.NewTableCell(floatToIntToString(player["lost_fumbles"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		var recovery string
		if int(player["own_rec"].(float64)) > 0 {
			recovery = floatToIntToString(player["own_rec"].(float64))
		} else {
			recovery = floatToIntToString(player["opp_rec"].(float64))
		}
		table.SetCell(homeCounter, 8,
			tview.NewTableCell(recovery).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		var recoveryYards string
		if int(player["own_rec_yards"].(float64)) > 0 {
			recoveryYards = floatToIntToString(player["own_rec_yards"].(float64))
		} else {
			recoveryYards = floatToIntToString(player["opp_rec_yards"].(float64))
		}

		table.SetCell(homeCounter, 9,
			tview.NewTableCell(recoveryYards).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		homeCounter++
		if homeCounter > rowCounter {
			rowCounter++
		}
	}
	return rowCounter
}
func buildKicking(table *tview.Table, summary GameSummary, rowCounter int) int {
	table.SetCell(rowCounter, 0,
		tview.NewTableCell("Kicking").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 1,
		tview.NewTableCell("FG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 2,
		tview.NewTableCell("LG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 3,
		tview.NewTableCell("AVG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 4,
		tview.NewTableCell("BLK").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 5,
		tview.NewTableCell("Kicking").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 6,
		tview.NewTableCell("FG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 7,
		tview.NewTableCell("LG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 8,
		tview.NewTableCell("AVG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 9,
		tview.NewTableCell("BLK").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	rowCounter++
	statCounter := 0
	for _, player := range summary.Statistics.Away.FieldGoals.Players {
		table.SetCell(rowCounter, 0,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(rowCounter, 1,
			tview.NewTableCell(floatToIntToString(player["made"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 2,
			tview.NewTableCell(floatToIntToString(player["longest"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(rowCounter, 3,
			tview.NewTableCell(floatToIntToString(player["avg_yards"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 4,
			tview.NewTableCell(floatToIntToString(player["blocked"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		rowCounter++
		statCounter++
	}
	homeCounter := rowCounter - statCounter
	for _, player := range summary.Statistics.Home.FieldGoals.Players {
		table.SetCell(homeCounter, 5,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(homeCounter, 6,
			tview.NewTableCell(floatToIntToString(player["made"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 7,
			tview.NewTableCell(floatToIntToString(player["longest"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 8,
			tview.NewTableCell(floatToIntToString(player["avg_yards"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 9,
			tview.NewTableCell(floatToIntToString(player["blocked"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		homeCounter++
		if homeCounter > rowCounter {
			rowCounter++
		}
	}
	return rowCounter
}
func buildKickReturns(table *tview.Table, summary GameSummary, rowCounter int) int {
	table.SetCell(rowCounter, 0,
		tview.NewTableCell("Kick Returns").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 1,
		tview.NewTableCell("NO").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 2,
		tview.NewTableCell("AVG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 3,
		tview.NewTableCell("TD").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 4,
		tview.NewTableCell("LG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 5,
		tview.NewTableCell("Kick Returns").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 6,
		tview.NewTableCell("NO").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 7,
		tview.NewTableCell("AVG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 8,
		tview.NewTableCell("TD").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 9,
		tview.NewTableCell("LG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	rowCounter++
	statCounter := 0
	for _, player := range summary.Statistics.Away.KickoffReturns.Players {
		table.SetCell(rowCounter, 0,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(rowCounter, 1,
			tview.NewTableCell(floatToIntToString(player["number"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 2,
			tview.NewTableCell(floatToIntToString(player["avg_yards"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(rowCounter, 3,
			tview.NewTableCell(floatToIntToString(player["touchdowns"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 4,
			tview.NewTableCell(floatToIntToString(player["longest"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		rowCounter++
		statCounter++
	}
	homeCounter := rowCounter - statCounter
	for _, player := range summary.Statistics.Home.KickoffReturns.Players {
		table.SetCell(homeCounter, 5,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(homeCounter, 6,
			tview.NewTableCell(floatToIntToString(player["number"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 7,
			tview.NewTableCell(floatToIntToString(player["avg_yards"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 8,
			tview.NewTableCell(floatToIntToString(player["touchdowns"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 9,
			tview.NewTableCell(floatToIntToString(player["longest"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		homeCounter++
		if homeCounter > rowCounter {
			rowCounter++
		}
	}
	return rowCounter
}
func buildPunts(table *tview.Table, summary GameSummary, rowCounter int) int {
	table.SetCell(rowCounter, 0,
		tview.NewTableCell("Punts").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 1,
		tview.NewTableCell("NO").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 2,
		tview.NewTableCell("AVG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 3,
		tview.NewTableCell("I20").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 4,
		tview.NewTableCell("LG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 5,
		tview.NewTableCell("Punts").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 6,
		tview.NewTableCell("NO").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 7,
		tview.NewTableCell("AVG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 8,
		tview.NewTableCell("I20").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 9,
		tview.NewTableCell("LG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	rowCounter++
	statCounter := 0
	for _, player := range summary.Statistics.Away.Punts.Players {
		table.SetCell(rowCounter, 0,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(rowCounter, 1,
			tview.NewTableCell(floatToIntToString(player["attempts"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 2,
			tview.NewTableCell(floatToIntToString(player["avg_yards"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(rowCounter, 3,
			tview.NewTableCell(floatToIntToString(player["inside_20"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 4,
			tview.NewTableCell(floatToIntToString(player["longest"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		rowCounter++
		statCounter++
	}
	homeCounter := rowCounter - statCounter
	for _, player := range summary.Statistics.Home.Punts.Players {
		table.SetCell(homeCounter, 5,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(homeCounter, 6,
			tview.NewTableCell(floatToIntToString(player["attempts"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 7,
			tview.NewTableCell(floatToIntToString(player["avg_yards"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 8,
			tview.NewTableCell(floatToIntToString(player["inside_20"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 9,
			tview.NewTableCell(floatToIntToString(player["longest"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		homeCounter++
		if homeCounter > rowCounter {
			rowCounter++
		}
	}
	return rowCounter
}
func buildPuntReturns(table *tview.Table, summary GameSummary, rowCounter int) int {
	table.SetCell(rowCounter, 0,
		tview.NewTableCell("Punt Returns").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 1,
		tview.NewTableCell("NO").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 2,
		tview.NewTableCell("AVG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 3,
		tview.NewTableCell("TD").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 4,
		tview.NewTableCell("LG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 5,
		tview.NewTableCell("Punt Returns").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 6,
		tview.NewTableCell("NO").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 7,
		tview.NewTableCell("AVG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 8,
		tview.NewTableCell("TD").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 9,
		tview.NewTableCell("LG").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	rowCounter++
	statCounter := 0
	for _, player := range summary.Statistics.Away.PuntReturns.Players {
		table.SetCell(rowCounter, 0,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(rowCounter, 1,
			tview.NewTableCell(floatToIntToString(player["number"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 2,
			tview.NewTableCell(floatToIntToString(player["avg_yards"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(rowCounter, 3,
			tview.NewTableCell(floatToIntToString(player["touchdowns"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 4,
			tview.NewTableCell(floatToIntToString(player["longest"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		rowCounter++
		statCounter++
	}
	homeCounter := rowCounter - statCounter
	for _, player := range summary.Statistics.Home.PuntReturns.Players {
		table.SetCell(homeCounter, 5,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(homeCounter, 6,
			tview.NewTableCell(floatToIntToString(player["number"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 7,
			tview.NewTableCell(floatToIntToString(player["avg_yards"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 8,
			tview.NewTableCell(floatToIntToString(player["touchdowns"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 9,
			tview.NewTableCell(floatToIntToString(player["longest"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		homeCounter++
		if homeCounter > rowCounter {
			rowCounter++
		}
	}
	return rowCounter
}
func buildDefense(table *tview.Table, summary GameSummary, rowCounter int) int {
	table.SetCell(rowCounter, 0,
		tview.NewTableCell("Defense").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 1,
		tview.NewTableCell("T/A").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 2,
		tview.NewTableCell("SCK").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 3,
		tview.NewTableCell("INT").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 4,
		tview.NewTableCell("FF").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 5,
		tview.NewTableCell("Defense").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 6,
		tview.NewTableCell("T/A").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 7,
		tview.NewTableCell("SCK").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 8,
		tview.NewTableCell("INT").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 9,
		tview.NewTableCell("FF").
			SetBackgroundColor(tcell.ColorGray).
			SetTextColor(tcell.ColorGreen).
			SetAlign(tview.AlignLeft))
	rowCounter++
	statCounter := 0
	for _, player := range summary.Statistics.Away.Defense.Players {
		table.SetCell(rowCounter, 0,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(rowCounter, 1,
			tview.NewTableCell(floatToIntToString(player["tackles"].(float64))+"/"+floatToIntToString(player["assists"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 2,
			tview.NewTableCell(floatToIntToString(player["sacks"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(rowCounter, 3,
			tview.NewTableCell(floatToIntToString(player["interceptions"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(rowCounter, 4,
			tview.NewTableCell(floatToIntToString(player["forced_fumbles"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		rowCounter++
		statCounter++
	}
	homeCounter := rowCounter - statCounter
	for _, player := range summary.Statistics.Home.Defense.Players {
		table.SetCell(homeCounter, 5,
			tview.NewTableCell(player["name"].(string)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		table.SetCell(homeCounter, 6,
			tview.NewTableCell(floatToIntToString(player["tackles"].(float64))+"/"+floatToIntToString(player["assists"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 7,
			tview.NewTableCell(floatToIntToString(player["sacks"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 8,
			tview.NewTableCell(floatToIntToString(player["interceptions"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))

		table.SetCell(homeCounter, 9,
			tview.NewTableCell(floatToIntToString(player["forced_fumbles"].(float64))).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		homeCounter++
		if homeCounter > rowCounter {
			rowCounter++
		}
	}
	return rowCounter
}

// BuildPlayerStatsTable builds the tview.Yable for player stats
func BuildPlayerStatsTable(summary GameSummary, table *tview.Table) int {
	var rowCounter = 0

	table.SetCell(rowCounter, 0,
		tview.NewTableCell(summary.Statistics.Away.Name).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 1,
		tview.NewTableCell("Stats").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 2,
		tview.NewTableCell("").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 3,
		tview.NewTableCell("").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 4,
		tview.NewTableCell("").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 5,
		tview.NewTableCell(summary.Statistics.Home.Name).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 6,
		tview.NewTableCell("Stats").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 7,
		tview.NewTableCell("").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 8,
		tview.NewTableCell("").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 9,
		tview.NewTableCell("").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))

	rowCounter++

	rowCounter = buildPassing(table, summary, rowCounter)
	rowCounter = buildRushing(table, summary, rowCounter)
	rowCounter = buildReceiving(table, summary, rowCounter)
	rowCounter = buildFumbles(table, summary, rowCounter)
	rowCounter = buildKicking(table, summary, rowCounter)
	rowCounter = buildKickReturns(table, summary, rowCounter)
	rowCounter = buildPunts(table, summary, rowCounter)
	rowCounter = buildPuntReturns(table, summary, rowCounter)
	rowCounter = buildDefense(table, summary, rowCounter)
	return rowCounter
}
