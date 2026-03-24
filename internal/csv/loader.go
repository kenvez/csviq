package csv

import (
	"encoding/csv"
	"errors"
	"os"
)

func LoadFromFile(path string) (*Table, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, errors.New("empty csv file")
	}

	table := &Table{
		Columns: records[0],
		Rows:    records[1:],
	}

	return table, nil
}
