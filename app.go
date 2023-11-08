package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type model struct {
	noResult        bool
	entries         []entry
	filteredEntries []entry
	selectedIdx     int
	height          int
	searchInput     textinput.Model
	lastSearchValue string
}

func initialModel(entries []entry) *model {
	return &model{
		entries:     entries,
		searchInput: textinput.New(),
	}
}

func (m *model) Init() tea.Cmd {
	m.searchInput.Focus()
	m.UpdateSearch()
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.noResult = true
			return m, tea.Quit

		case "up":
			m.SelectUp()

		case "down":
			m.SelectDown()

		case "enter":
			if len(m.filteredEntries) != 0 {
				return m, tea.Quit
			}

		case "esc":
			m.searchInput.SetValue("")
			m.selectedIdx = 0
		}

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.UpdateSearch()
	}

	var cmd tea.Cmd
	m.searchInput, cmd = m.searchInput.Update(msg)

	if m.lastSearchValue != m.searchInput.Value() {
		m.UpdateSearch()
		m.lastSearchValue = m.searchInput.Value()
	}

	return m, cmd
}

func (m *model) View() string {
	b := strings.Builder{}

	skip := m.MaxEntriesHeight() - len(m.filteredEntries)
	for i := 0; i < skip; i++ {
		b.WriteRune('\n')
	}

	for i, entry := range m.filteredEntries {
		if i == m.MaxEntriesHeight() {
			break
		}

		isActive := m.selectedIdx == i
		if isActive {
			b.WriteString("\x1b[38;2;140;24;226m") // set foreground color to #8C18E2
			b.WriteString(entry.prefix)
			b.WriteString(" / ")
			b.WriteString(entry.name)
			b.WriteString("\x1b[1;0m") // reset foreground color
		} else {
			b.WriteString(entry.prefix)
			b.WriteString(" / ")
			b.WriteString(entry.name)
		}
		b.WriteRune('\n')
	}

	if len(m.filteredEntries) > m.MaxEntriesHeight() {
		b.WriteString("...\n")
	} else {
		b.WriteRune('\n')
	}

	b.WriteString("\n")

	b.WriteString(m.searchInput.View())

	return b.String()
}

func (m *model) UpdateSearch() {
	search := m.searchInput.Value()

	if search == "" {
		m.filteredEntries = m.entries
	} else {
		results := make([]entry, 0)

		for _, entry := range m.entries {
			rank := fuzzy.RankMatchFold(search, entry.searchVector)
			if rank != -1 {
				entry.rank = rank
				results = append(results, entry)
			}
		}

		slices.SortStableFunc(results, func(left, right entry) int { return left.rank - right.rank })
		m.filteredEntries = results
	}

	m.selectedIdx = 0
	m.searchInput.Prompt = fmt.Sprintf("%d/%d > ", len(m.filteredEntries), len(m.entries))
}

func (m *model) SelectUp() {
	m.selectedIdx--
	if m.selectedIdx < 0 {
		m.selectedIdx = 0
	}
}

func (m *model) SelectDown() {
	m.selectedIdx++
	maxIdx := min(len(m.filteredEntries), m.MaxEntriesHeight()) - 1
	if m.selectedIdx >= maxIdx {
		m.selectedIdx = maxIdx
	}
}

func (m *model) MaxEntriesHeight() int {
	return m.height - 3
}
