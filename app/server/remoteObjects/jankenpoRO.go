package remoteObjects

type IJankenpo interface {
	Play(player1Move, player2Move string) (float64, error)
}

type Jankenpo struct {
}

func (Jankenpo) Play(player1Move, player2Move string) (float64, error) {
	possibilities := []string{"P", "A", "T"}

	if !inArray(player1Move, possibilities) {
		return -1, nil
	}
	if !inArray(player2Move, possibilities) {
		return -1, nil
	}
	if player1Move == player2Move {
		return 0, nil
	}
	switch player1Move {
	case "P":
		if player2Move == "A" {
			return 2, nil
		} else {
			return 1, nil
		}
	case "A":
		if player2Move == "P" {
			return 1, nil
		} else {
			return 2, nil
		}
	case "T":
		if player2Move == "P" {
			return 2, nil
		} else {
			return 1, nil
		}
	default:
		return -1, nil
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
