package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	hexColor := flag.String("c", "#8C18E2", "active item color")

	flag.Parse()

	if flag.NArg()%2 != 0 {
		fmt.Println("provide an even number of arguments")
		os.Exit(1)
	}

	paths := make([]string, 0)
	prefixes := make([]string, 0)
	for i, arg := range flag.Args() {
		if i%2 == 0 {
			paths = append(paths, arg)
		} else {
			prefixes = append(prefixes, arg)
		}
	}

	colorEscapeSeqneuce, err := hexColorToEscapeSequence(*hexColor)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	entries, err := collectEntries(paths, prefixes)
	if err != nil {
		fmt.Println("failed to collect entries:", err)
		os.Exit(1)
	}

	app := tea.NewProgram(initialModel(entries, colorEscapeSeqneuce), tea.WithAltScreen(), tea.WithOutput(os.Stderr))

	finalModel, err := app.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m := finalModel.(*model)
	if m.noResult {
		os.Exit(1)
	}

	selected := m.filteredEntries[m.selectedIdx]
	fmt.Println(selected.fullPath)
}

func hexColorToEscapeSequence(hex string) (string, error) {
	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}

	if len(hex) != 6 {
		return "", fmt.Errorf("'%s' is not a valid hex color", hex)
	}

	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return "", err
	}

	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return "", err
	}

	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b), nil
}
