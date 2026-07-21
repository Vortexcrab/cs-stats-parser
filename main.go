package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Vortexcrab/cs-stats-parser/internal/stats"
)

func main() {
	fmt.Println("CS Stats Parser starting...")

	match := stats.Match{
		Nick:        "s1mple",
		Kills:       10,
		Deaths:      5,
		RoundResult: true,
	}
	match2 := stats.Match{
		Nick:        "ZywOo",
		Kills:       8,
		Deaths:      7,
		RoundResult: false,
	}
	match3 := stats.Match{
		Nick:        "device",
		Kills:       12,
		Deaths:      4,
		RoundResult: true,
	}
	for _, m := range []stats.Match{match, match2, match3} {
		fmt.Printf("%s kills: %d, deaths: %d, won: %v\n", m.Nick, m.Kills, m.Deaths, m.RoundResult)
	}

	fmt.Println("---------------------------------")

	file, err := os.Open("testdata/matches.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records[1:] {
		nick := record[0]
		kills, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}
		deaths, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatal(err)
		}
		result, err := strconv.ParseBool(record[3])
		if err != nil {
			log.Fatal(err)
		}
		match := stats.Match{
			Nick:        nick,
			Kills:       kills,
			Deaths:      deaths,
			RoundResult: result}
		fmt.Printf("%s kills: %d, deaths: %d, won: %v\n", match.Nick, match.Kills, match.Deaths, match.RoundResult)
	}
}
