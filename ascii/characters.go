package ascii

import (
	"fmt"
	"strings"
)

var joker = `
 ___ 
/ _ \
\// /
  \/ 
  ()
`

var jokerInf letterInfos
var jokerLines []string

func init() {
	a, b := processOne(joker)
	jokerInf = b
	jokerLines = a
}

type letterInfos struct {
	lineNum  int
	maxWidth int
}

type Characters struct {
	font map[string][]string
	info map[string]letterInfos
}

func processOne(s string) ([]string, letterInfos) {
	lines := strings.Split(s, "\n")
	maxw := 0
	for i, line := range lines {
		if len(line) > maxw {
			maxw = len(line)
		}
		lines[i] = line
	}
	return lines, letterInfos{
		len(lines),
		maxw,
	}
}

func process(m map[string]string) (map[string][]string, map[string]letterInfos) {
	tr := map[string][]string{}
	inf := map[string]letterInfos{}
	for k, v := range m {
		a, b := processOne(v)
		tr[k] = a
		inf[k] = b
	}
	return tr, inf
}

func NewCharacters(m map[string]string) Characters {
	trimmed, infos := process(m)
	return Characters{trimmed, infos}
}

func padRight(s string, width int) string {
	if len(s) < width {
		s += strings.Repeat(" ", width-len(s))
	}
	return s
}

func (c Characters) getOne(s string) ([]string, letterInfos) {
	linf, ok := c.info[s]
	if !ok {
		joker, okj := c.info["?"]
		if okj {
			return c.font["?"], joker
		}
		return jokerLines, jokerInf
	}
	return c.font[s], linf
}

func (c Characters) print(text string, printOut bool) string {
	maxHeight := 0
	// Calculating max height of banner
	for _, v := range text {
		_, linf := c.getOne(string(v))
		if linf.lineNum > maxHeight {
			maxHeight = linf.lineNum
		}
	}
	ret := ""
	// Render
	for i := 0; i < maxHeight-1; i++ {
		thisLin := ""
		for _, v := range text {
			lines, linf := c.getOne(string(v))
			if linf.lineNum <= i {
				thisLin += padRight("", linf.maxWidth)
			} else {
				thisLin += padRight(lines[i], linf.maxWidth)
			}
		}
		if printOut {
			fmt.Println(thisLin)
		} else {
			ret += thisLin + "\n"
		}
	}
	return ret
}

func (c Characters) GetString(text string) string {
	return c.print(text, false)
}

func (c Characters) Print(text string) {
	c.print(text, true)
}

// Prints the ascii
func Print(s string) {
	NewCharacters(AsciiMap).Print(s)
}

// GetString returns string
func GetString(s string) string {
	return NewCharacters(AsciiMap).GetString(s)
}
