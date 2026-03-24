package tui

import (
	"strings"

	tea "charm.land/bubbletea/v2"

	"github.com/kenvez/csviq/internal/csv"
)

type model struct {
	table        *csv.Table
	cursorRow    int
	cursorColumn int
	scrollX      int
	scrollY      int
	width        int
	height       int
}

func InitialModel(table *csv.Table) model {
	return model{
		table: table,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	widths := columnWidths(m.table)
	var b strings.Builder

	b.WriteString(renderRow(m.table.Columns, widths))
	b.WriteString("\n")

	for i, w := range widths {
		if i > 0 {
			b.WriteString("  ")
		}

		b.WriteString(strings.Repeat("─", w))
	}

	b.WriteString("\n")

	for _, row := range m.table.Rows {
		b.WriteString(renderRow(row, widths))
		b.WriteString("\n")
	}

	return tea.NewView(b.String())
}
