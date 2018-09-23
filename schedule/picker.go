package schedule

import (
	"fmt"
	"strings"

	"github.com/scottjustin5000/nfl-cli/ascii"

	"github.com/manifoldco/promptui"
)

// CallbackFunc function to be called on game selection
type CallbackFunc func(game Game)

// LoadSchedule loads the NFL schedule for the given season week
func LoadSchedule(seasonType string, season int, week int, apiKey string, cb CallbackFunc) {
	scheduledWeek := GetScheduleWeek(seasonType, season, week, apiKey)

	fmt.Println("")
	funcMap := promptui.FuncMap
	funcMap["truncate"] = func(size int, input string) string {
		if len(input) <= size {
			return input
		}
		return input[:size-3]
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F3C8 {{ .Display | green | truncate 50  }} | {{ .GameDate | cyan | truncate 25 }} | {{ .Status | red }} ",
		Inactive: "  {{ .Display | green  | truncate 50  }} | {{ .GameDate | cyan | truncate 25 }} | {{ .Status | red }}",
		Selected: "\U0001F3C8 {{ .Display | red | cyan }}",
		Details: `
--------------- Selected Game ----------------
{{ "Game:" | faint }}	{{ .Display }}
{{ "Date:" | faint }}	{{ .GameDate }}
{{ "Status:" | faint }}	{{ .Status }}`,
	}

	searcher := func(input string, index int) bool {
		game := scheduledWeek.Games[index]
		name := strings.Replace(strings.ToLower(game.Display), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Which Game",
		Items:     scheduledWeek.Games,
		Templates: templates,
		Size:      10,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Println("LOADING.....")

	home := strings.Fields(scheduledWeek.Games[i].Home.Name)
	away := strings.Fields(scheduledWeek.Games[i].Away.Name)

	homeName := home[len(home)-1]
	awayName := away[len(away)-1]
	ascii.Print(strings.ToLower(awayName) + " v. " + strings.ToLower(homeName))
	cb(scheduledWeek.Games[i])
}
