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

		case "up", "k":
			if m.cursorRow > 0 {
				m.cursorRow--
				if m.cursorRow < m.scrollY {
					m.scrollY--
				}
			}

		case "down", "j":
			if m.cursorRow < len(m.table.Rows)-1 {
				m.cursorRow++
				if m.cursorRow >= m.scrollY+m.visibleRows() {
					m.scrollY++
				}
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
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

	start := m.scrollY
	end := min(start+m.visibleRows(), len(m.table.Rows))

	for i := start; i < end; i++ {
		line := renderRow(m.table.Rows[i], widths)

		if i == m.cursorRow {
			line = highlightStyle.Render(line)
		}

		b.WriteString(line)
		b.WriteString("\n")
	}

	return tea.NewView(b.String())
}

func (m model) visibleRows() int {
	return m.height - 4
}
