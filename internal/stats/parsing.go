package stats

import (
	"encoding/csv"
	"os"
	"strconv"
)

func ParseCSV(filename string) ([]Match, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var matches []Match
	for _, record := range records[1:] {
		nick := record[0]
		kills, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}
		deaths, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}
		result, err := strconv.ParseBool(record[3])
		if err != nil {
			return nil, err
		}
		match := Match{
			Nick:        nick,
			Kills:       kills,
			Deaths:      deaths,
			RoundResult: result}
		matches = append(matches, match)
	}
	return matches, nil
}
