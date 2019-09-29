package tenniskata

type Game interface {
	WonPoint(playerName string)
	GetScore() string
}

type tennisGame struct {
	p2  int
	p1  int
	p1N string
	p2N string
}

func TennisGame(p1N string, p2N string) Game {
	game := &tennisGame{
		p1N: p1N,
		p2N: p2N}

	return game
}

func (game *tennisGame) GetScore() string {
	var s string
	if game.p1 < 4 && game.p2 < 4 && !(game.p1+game.p2 == 6) {
		p := []string{"Love", "Fifteen", "Thirty", "Forty"}
		s = p[game.p1]
		if game.p1 == game.p2 {
			return s + "-All"
		}
		return s + "-" + p[game.p2]
	} else {
		if game.p1 == game.p2 {
			return "Deuce"
		}
		if game.p1 > game.p2 {
			s = game.p1N
		} else {
			s = game.p2N
		}
		if (game.p1-game.p2)*(game.p1-game.p2) == 1 {
			return "Advantage " + s
		}
		return "Win for " + s
	}
}

func (game *tennisGame) WonPoint(playerName string) {
	if playerName == "player1" {
		game.p1 += 1
	} else {
		game.p2 += 1
	}
}
