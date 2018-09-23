package gamecenter

import (
	"sort"
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func groupByQuarter(scoringDrives []ScoringDrive) map[int][]ScoringDrive {
	m := make(map[int][]ScoringDrive)
	for _, drive := range scoringDrives {
		if _, ok := m[drive.Quarter.Sequence]; ok {
			m[drive.Quarter.Sequence] = append(m[drive.Quarter.Sequence], drive)
		} else {
			m[drive.Quarter.Sequence] = []ScoringDrive{drive}
		}
	}
	return m
}

func BuildScoringTable(scoringDrives []ScoringDrive, table *tview.Table) int {
	scoring := groupByQuarter(scoringDrives)
	var keys []int
	for k := range scoring {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	var rowCounter = 0

	for _, key := range keys {
		var scoringDrives = scoring[key]
		table.SetCell(rowCounter, 0,
			tview.NewTableCell("Quarter "+strconv.Itoa(key)).
				SetTextColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft))
		rowCounter++
		for _, drive := range scoringDrives {

			var clock = drive.Clock
			table.SetCell(rowCounter, 0,
				tview.NewTableCell(clock).
					SetTextColor(tcell.ColorGray).
					SetAlign(tview.AlignLeft))

			table.SetCell(rowCounter, 1,
				tview.NewTableCell("").
					SetTextColor(tcell.ColorGray).
					SetAlign(tview.AlignLeft))
			table.SetCell(rowCounter, 2,
				tview.NewTableCell("").
					SetTextColor(tcell.ColorGray).
					SetAlign(tview.AlignLeft))
			table.SetCell(rowCounter, 3,
				tview.NewTableCell(drive.EndReason).
					SetTextColor(tcell.ColorGray).
					SetAlign(tview.AlignLeft))
			table.SetCell(rowCounter, 4,
				tview.NewTableCell(strconv.Itoa(drive.PlayCount)+" Plays, "+strconv.Itoa(drive.Gain)+" Yards, "+drive.Duration+" Minutes").
					SetTextColor(tcell.ColorGray).
					SetAlign(tview.AlignCenter))
			rowCounter++
			for _, event := range drive.Plays {
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
	return rowCounter
}
