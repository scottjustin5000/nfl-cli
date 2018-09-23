package gamecenter

import (
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// BuildPbpTable builds the tview.Yable for the game's play-by-play
func BuildPbpTable(pbpFeed PbpFeed, table *tview.Table) int {

	var rowCounter = 0

	table.SetCell(rowCounter, 0,
		tview.NewTableCell("Quarter").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))
	table.SetCell(rowCounter, 1,
		tview.NewTableCell("Pos").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 2,
		tview.NewTableCell("Loc").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 3,
		tview.NewTableCell("D&D").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))

	table.SetCell(rowCounter, 4,
		tview.NewTableCell("Description").
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))

	rowCounter++

	for _, period := range pbpFeed.Periods {

		table.SetCell(rowCounter, 0,
			tview.NewTableCell("Quarter "+strconv.Itoa(period.Number)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		rowCounter++
		for _, pbp := range period.Pbp {

			if pbp.Type == "drive" {
				var clock = ""
				if len(pbp.Events) > 0 {
					clock = pbp.Events[0].Clock
				}
				table.SetCell(rowCounter, 0,
					tview.NewTableCell(clock).
						SetTextColor(tcell.ColorGreen).
						SetAlign(tview.AlignLeft))

				table.SetCell(rowCounter, 1,
					tview.NewTableCell("").
						SetTextColor(tcell.ColorGreen).
						SetAlign(tview.AlignLeft))
				table.SetCell(rowCounter, 2,
					tview.NewTableCell("").
						SetTextColor(tcell.ColorGreen).
						SetAlign(tview.AlignLeft))
				table.SetCell(rowCounter, 3,
					tview.NewTableCell(pbp.EndReason).
						SetTextColor(tcell.ColorGreen).
						SetAlign(tview.AlignLeft))
				table.SetCell(rowCounter, 4,
					tview.NewTableCell(strconv.Itoa(pbp.PlayCount)+" Plays, "+strconv.Itoa(pbp.Gain)+" Yards, "+pbp.Duration+" Minutes").
						SetTextColor(tcell.ColorGreen).
						SetAlign(tview.AlignCenter))
				rowCounter++
				for _, event := range pbp.Events {
					table.SetCell(rowCounter, 0,
						tview.NewTableCell(event.Clock).
							SetTextColor(tcell.ColorWhite).
							SetAlign(tview.AlignLeft))

					table.SetCell(rowCounter, 1,
						tview.NewTableCell(event.StartSituation.Possession.Alias).
							SetTextColor(tcell.ColorWhite).
							SetAlign(tview.AlignLeft))

					table.SetCell(rowCounter, 2,
						tview.NewTableCell(event.StartSituation.Location.Alias+" "+strconv.Itoa(event.StartSituation.Location.Yardline)).
							SetTextColor(tcell.ColorWhite).
							SetAlign(tview.AlignLeft))

					table.SetCell(rowCounter, 3,
						tview.NewTableCell(strconv.Itoa(event.StartSituation.Down)+" & "+strconv.Itoa(event.StartSituation.Yfd)).
							SetTextColor(tcell.ColorWhite).
							SetAlign(tview.AlignLeft))

					table.SetCell(rowCounter, 4,
						tview.NewTableCell(event.Description).
							SetTextColor(tcell.ColorWhite).
							SetAlign(tview.AlignLeft))

					rowCounter++

				}

			}
		}
	}
	return rowCounter
}
