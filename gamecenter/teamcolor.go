package gamecenter

import (
	"github.com/gdamore/tcell"
)

func getTeamColor(alias string) tcell.Color {
	var c tcell.Color
	switch alias {
	case "DEN", "CIN", "CHI", "CLE":
		c = tcell.ColorOrange
	case "ARZ", "ATL", "HOU", "KC", "NE", "SF":
		c = tcell.ColorRed
	case "BAL", "MIN":
		c = tcell.ColorPurple
	case "BUF", "IND", "NYG":
		c = tcell.ColorBlue
	case "CAR":
		c = tcell.ColorAliceBlue
	case "DAL":
		c = tcell.ColorBlue
	case "GB", "PIT":
		c = tcell.ColorYellow
	case "LAC", "TEN":
		c = tcell.ColorPowderBlue
	case "DET":
		c = tcell.ColorLightBlue
	case "MIA":
		c = tcell.ColorTeal
	case "JAC":
		c = tcell.ColorPaleTurquoise
	case "NO":
		c = tcell.ColorGold
	case "NYJ":
	case "PHI":
		c = tcell.ColorGreen
	case "OAK":
		c = tcell.ColorSilver
	case "LAR":
		c = tcell.ColorRoyalBlue
	case "SEA":
		c = tcell.ColorLimeGreen
	case "TB":
		c = tcell.ColorLightSalmon
	case "WAS":
		c = tcell.ColorMaroon
	default:
		c = tcell.ColorWhite
	}
	return c
}
