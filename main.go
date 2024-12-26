package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	hexColor := flag.String("c", "#8C18E2", "active item color")
	showHidden := flag.Bool("a", false, "show hidden directories (prefixed by a dot)")

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

	entries, err := collectEntries(paths, prefixes, *showHidden)
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
