package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Args)%2 == 0 {
		fmt.Println("provide an even number of arguments")
		os.Exit(1)
	}

	paths := make([]string, 0)
	prefixes := make([]string, 0)
	for i, arg := range os.Args {
		if i == 0 {
			continue //skip executable
		}

		if i%2 == 1 {
			paths = append(paths, arg)
		} else {
			prefixes = append(prefixes, arg)
		}
	}

	entries, err := collectEntries(paths, prefixes)
	if err != nil {
		fmt.Println("failed to collect entries:", err)
		os.Exit(1)
	}

	app := tea.NewProgram(initialModel(entries), tea.WithAltScreen(), tea.WithOutput(os.Stderr))

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
