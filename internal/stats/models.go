package stats

type Match struct {
	Nick        string
	Kills       int
	Deaths      int
	RoundResult bool
}

type Player struct {
	Nick    string
	Wins    int
	Defeats int
	Kills   int
	Deaths  int
}
