package main

import (
	"fmt"
	"strings"
)

const clear_line = "\033[1A\033[2K"

func clear_lines(lines int) {
	fmt.Print(strings.Repeat(clear_line, lines))
}

func colored(color int, value string) string {
	return fmt.Sprintf("\033[38;5;%dm%s\033[0m", color, value)
}

func block(color int) string {
	return colored(color, "██")
}
func empty() string {
	return "  "
}
