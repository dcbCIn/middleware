package remoteObjects

type Jankenpo struct {
}

func (Jankenpo) Process(player1Move, player2Move string) int {
	possibilities := []string{"P", "A", "T"}

	if !inArray(player1Move, possibilities) {
		return -1
	}
	if !inArray(player2Move, possibilities) {
		return -1
	}
	if player1Move == player2Move {
		return 0
	}
	switch player1Move {
	case "P":
		if player2Move == "A" {
			return 2
		} else {
			return 1
		}
	case "A":
		if player2Move == "P" {
			return 1
		} else {
			return 2
		}
	case "T":
		if player2Move == "P" {
			return 2
		} else {
			return 1
		}
	default:
		return -1
	}
}

func inArray(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
