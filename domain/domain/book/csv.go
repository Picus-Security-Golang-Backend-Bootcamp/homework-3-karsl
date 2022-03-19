package book

import (
	"encoding/csv"
	"os"
	"strconv"
)

func readFromCsv(filename string) ([]Book, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var result []Book
	for _, line := range lines[1:] {
		name := line[0]
		stockCode := line[1]
		authorId, err := strconv.Atoi(line[2])
		if err != nil {
			return nil, err
		}

		data, err := Construct(name, stockCode, authorId)
		if err != nil {
			return nil, err
		}

		result = append(result, data)
	}

	return result, nil
}
