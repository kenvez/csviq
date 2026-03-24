package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"

	"github.com/kenvez/csviq/internal/csv"
)

type mode int

const (
	modeNormal mode = iota
	modeEdit
)

type model struct {
	table        *csv.Table
	cursorRow    int
	cursorColumn int
	scrollX      int
	scrollY      int
	width        int
	height       int
	mode         mode
	editBuffer   string
	path         string
}

func InitialModel(table *csv.Table, path string) model {
	return model{
		table: table,
		path:  path,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch m.mode {
		case modeNormal:
			switch msg.String() {

			case "ctrl+c", "q":
				return m, tea.Quit

			case "ctrl+s":
				m.table.SaveToFile(m.path)

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

			case "left", "h":
				if m.cursorColumn > 0 {
					m.cursorColumn--
				}

			case "right", "l":
				if m.cursorColumn < len(m.table.Columns)-1 {
					m.cursorColumn++
				}

			case "g":
				m.cursorRow = 0
				m.scrollY = 0
			case "G":
				m.cursorRow = len(m.table.Rows) - 1

				if m.cursorRow >= m.visibleRows() {
					m.scrollY = m.cursorRow - m.visibleRows() + 1
				}
			case "e":
				m.mode = modeEdit
				m.editBuffer = m.table.Rows[m.cursorRow][m.cursorColumn]
			}

		case modeEdit:
			switch msg.String() {
			case "enter":
				m.table.Rows[m.cursorRow][m.cursorColumn] = m.editBuffer
				m.mode = modeNormal
			case "escape":
				m.mode = modeNormal
			case "backspace":
				if len(m.editBuffer) > 0 {
					m.editBuffer = m.editBuffer[:len(m.editBuffer)-1]
				}
			case "space":
				m.editBuffer += " "
			default:
				m.editBuffer += msg.String()
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
		var line strings.Builder

		for j, cell := range m.table.Rows[i] {
			if j > 0 {
				line.WriteString("  ")
			}

			display := cell

			if m.mode == modeEdit && i == m.cursorRow && j == m.cursorColumn {
				display = m.editBuffer + "▎"
			}

			formatted := fmt.Sprintf("%-*s", widths[j], display)

			if i == m.cursorRow && j == m.cursorColumn {
				formatted = highlightStyle.Render(formatted)
			}

			line.WriteString(formatted)
		}

		b.WriteString(line.String())
		b.WriteString("\n")
	}

	return tea.NewView(b.String())
}

func (m model) visibleRows() int {
	return m.height - 4
}
