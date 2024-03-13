package main

import (
	"fmt"
	"strings"
)

var COLORS = map[int]int{
	PcO: 3,
	PcS: 2,
	PcZ: 1,
	PcT: 5,
	PcL: 7,
	PcJ: 4,
	PcI: 6,
}

func title() {
	var title = colored(COLORS[1], "G") +
		colored(COLORS[2], "O") +
		colored(COLORS[3], "T") +
		colored(COLORS[4], "R") +
		colored(COLORS[5], "I") +
		colored(COLORS[6], "S")

	fmt.Println("\n          " + title)
}

const clear_line = "\033[1A\033[2K"

func clear_lines(lines int) {
	fmt.Print(strings.Repeat(clear_line, lines))
}

func colored(color int, value string) string {
	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", color, value)
}

func block(color int) string {
	return colored(COLORS[color], "██")
}

func empty(odd bool) string {
	if odd {
		return "\033[48;5;0m  \033[0m"
	}
	return "  "
}

func container(lines []string, w int) []string {
	for i := range lines {
		lines[i] = "│" + lines[i] + "│"
	}
	lines = append([]string{"┌" + strings.Repeat("─", w) + "┐"}, lines...)
	lines = append(lines, "└"+strings.Repeat("─", w)+"┘")
	return lines
}
