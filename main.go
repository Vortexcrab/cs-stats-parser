package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

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
		fmt.Println(record)
	}
}
