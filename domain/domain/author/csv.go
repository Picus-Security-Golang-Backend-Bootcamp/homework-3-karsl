package author

import (
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

func cellToAuthor(line []string) (Author, error) {
	var birthDate = time.Now()

	if line[2] != "" {
		splittedDate := strings.Split(line[2], "/")
		year, err := strconv.Atoi(splittedDate[0])
		if err != nil {
			return Author{}, err
		}
		month, err := strconv.Atoi(splittedDate[1])
		if err != nil {
			return Author{}, err
		}
		day, err := strconv.Atoi(splittedDate[2])
		if err != nil {
			return Author{}, err
		}

		birthDate = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}

	id, err := strconv.Atoi(line[0])
	if err != nil {
		return Author{}, err
	}

	data := Author{
		Model: gorm.Model{
			ID: uint(id),
		},
		Name:      line[1],
		BirthDate: birthDate,
	}

	return data, nil
}

func cellsToAuthors(lines [][]string) ([]Author, error) {
	var result []Author
	for _, line := range lines[1:] {
		data, err := cellToAuthor(line)
		if err != nil {
			return nil, err
		}

		result = append(result, data)
	}

	return result, nil
}
