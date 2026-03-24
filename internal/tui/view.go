package tui

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"

	"github.com/kenvez/csviq/internal/csv"
)

var highlightStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#3C3489")).
	Foreground(lipgloss.Color("#EEEDFE"))

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
