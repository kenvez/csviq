package csv

import (
	"encoding/csv"
	"os"
)

func (t *Table) SaveToFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	writer.Write(t.Columns)

	for _, row := range t.Rows {
		writer.Write(row)
	}

	writer.Flush()

	return writer.Error()
}
