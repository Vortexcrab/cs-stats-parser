package main

import (
	/*"encoding/csv"*/
	"fmt"
	"log"

	/*"os"
	"strconv"*/

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
	matches, err := stats.ParseCSV("testdata/matches.csv")
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range matches {
		fmt.Printf("%s kills: %d, deaths: %d, won: %v\n", m.Nick, m.Kills, m.Deaths, m.RoundResult)
	}

	grouped := make(map[string][]stats.Match)
	for _, m := range matches {
		grouped[m.Nick] = append(grouped[m.Nick], m)
	}

	fmt.Println("---------------------------------")
	for _, playerMatches := range grouped {
		player := stats.CalcStats(playerMatches)
		fmt.Printf("nick: %s kills: %d, deaths %d, wins: %d, defeats: %d\n", player.Nick, player.Kills, player.Deaths, player.Wins, player.Defeats)
	}
}
