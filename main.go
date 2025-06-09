package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "failed:", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("the first argument must be a path to the config file")
	}

	configFilePath := os.Args[1]
	config, err := LoadConfig(configFilePath)
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	activeColorEscapeSequence, err := hexColorToEscapeSequence(config.ActiveColor)
	if err != nil {
		return fmt.Errorf("convert activeColor: %w", err)
	}

	menuEntries, err := MakeMenuEntries(config)
	if err != nil {
		return fmt.Errorf("make menu entries: %w", err)
	}

	app := tea.NewProgram(initialModel(menuEntries, activeColorEscapeSequence), tea.WithAltScreen(), tea.WithOutput(os.Stderr))

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
	fmt.Println(selected.Value)

	return nil
}
