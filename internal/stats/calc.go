package stats

func CalcStats(matches []Match) Player {
	var totalKills, totalDeaths, totalWins, totalDefeats int
	for _, m := range matches {
		totalKills += m.Kills
		totalDeaths += m.Deaths
		if m.RoundResult {
			totalWins += 1
		} else {
			totalDefeats += 1
		}
	}
	return Player{
		Nick:    matches[0].Nick,
		Kills:   totalKills,
		Deaths:  totalDeaths,
		Wins:    totalWins,
		Defeats: totalDefeats,
	}
}
