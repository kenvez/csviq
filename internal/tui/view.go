package tui

import (
	"fmt"
	"path/filepath"
	"strings"

	"charm.land/lipgloss/v2"

	"github.com/kenvez/csviq/internal/csv"
)

var highlightStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#3C3489")).
	Foreground(lipgloss.Color("#EEEDFE"))

var statusBarStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#87fc77"))

func columnWidths(table *csv.Table) []int {
	widths := make([]int, len(table.Columns))

	for i := range widths {
		widths[i] = len(table.Columns[i])
	}

	for _, row := range table.Rows {
		for i, cell := range row {
			if len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}

	return widths
}

func renderRow(cells []string, widths []int) string {
	var b strings.Builder

	for i, cell := range cells {
		if i > 0 {
			b.WriteString("  ")
		}

		fmt.Fprintf(&b, "%-*s", widths[i], cell)
	}

	return b.String()
}

func renderStatusBar(m model) string {
	var mode string

	if m.mode == modeNormal {
		mode = "NORMAL"
	} else {
		mode = "EDIT"
	}

	filename := filepath.Base(m.path)

	return fmt.Sprintf(" %s | Row %d, Col %d | %s",
		mode,
		m.cursorRow+1,
		m.cursorColumn+1,
		filename,
	)
}
