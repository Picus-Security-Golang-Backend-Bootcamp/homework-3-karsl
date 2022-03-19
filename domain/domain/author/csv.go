package author

import (
	"encoding/csv"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
	"time"
)

func readFromCsv(filename string) ([]Author, error) {
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

	var result []Author
	for _, line := range lines[1:] {
		var birthDate = time.Now()

		if line[2] != "" {
			splittedDate := strings.Split(line[2], "/")
			year, err := strconv.Atoi(splittedDate[0])
			if err != nil {
				return nil, err
			}
			month, err := strconv.Atoi(splittedDate[1])
			if err != nil {
				return nil, err
			}
			day, err := strconv.Atoi(splittedDate[2])
			if err != nil {
				return nil, err
			}

			birthDate = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		}

		id, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, err
		}

		data := Author{
			Model: gorm.Model{
				ID: uint(id),
			},
			Name:      line[1],
			BirthDate: birthDate,
		}

		result = append(result, data)
	}

	return result, nil
}
