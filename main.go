package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	"github.com/kenvez/csviq/internal/csv"
	"github.com/kenvez/csviq/internal/tui"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("please provide a path to the csv file")

		os.Exit(1)
	}

	path := os.Args[1]
	table, err := csv.LoadFromFile(path)

	if err != nil {
		fmt.Printf("cannot load file: %v", err)

		os.Exit(1)
	}

	model := tui.InitialModel(table)
	program := tea.NewProgram(model)

	if _, err := program.Run(); err != nil {
		fmt.Printf("error running csviq: %v", err)

		os.Exit(1)
	}
}
