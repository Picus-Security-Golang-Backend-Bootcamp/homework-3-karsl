package book

import (
	"strconv"
)

// lineToBook parses single line to Book
func lineToBook(line []string) (Book, error) {
	name := line[0]
	stockCode := line[1]
	authorId, err := strconv.Atoi(line[2])
	if err != nil {
		return Book{}, err
	}

	data, err := Construct(name, stockCode, authorId)
	if err != nil {
		return Book{}, err
	}

	return data, nil
}

// linesToBook parses lines to Books
func linesToBook(lines [][]string) ([]Book, error) {
	var result []Book
	for _, line := range lines[1:] {
		data, err := lineToBook(line)
		if err != nil {
			return nil, err
		}

		result = append(result, data)
	}

	return result, nil
}
