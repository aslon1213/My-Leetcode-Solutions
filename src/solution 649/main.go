package main

func main() {
	predictPartyVictory("RDD")
}

func predictPartyVictory(senate string) string {
	radiantBan, radiantPick, direBan, direPick, round := 0, -1, 0, -1, []byte(senate)
	for radiantPick*direPick != 0 {
		radiantPick, direPick = 0, 0
		nextRound := []byte{}
		for _, v := range round {
			if v == 'R' {
				if direBan > 0 {
					direBan--
				} else {
					nextRound = append(nextRound, 'R')
					radiantBan++
					radiantPick++
				}
			} else {
				if radiantBan > 0 {
					radiantBan--
				} else {
					nextRound = append(nextRound, 'D')
					direBan++
					direPick++
				}
			}
		}
		round = nextRound
	}

	if radiantPick > 0 {
		return "Radiant"
	}
	if direPick > 0 {
		return "Dire"
	}
	return ""
}
