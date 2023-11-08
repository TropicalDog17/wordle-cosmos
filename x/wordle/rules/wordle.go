package rules

import (
	"fmt"
	"regexp"
	"strings"
)

type State int64

const (
	Wrong   State = 0
	Partial State = 1
	True    State = 2
)
const maxMovesAllowed = 6

type Game struct {
	SecretWord string
	MoveCount  int
	IsWin      bool
}

type GuessState []State

func New(secret string) *Game {
	game := &Game{SecretWord: secret, MoveCount: 0, IsWin: false}
	return game
}

func (game *Game) GetSecretLength() int {
	return len(game.SecretWord)
}

func (game *Game) Guess(word string) (GuessState, error) {
	if game.IsWin || game.IsOutOfMove() {
		return GuessState{}, fmt.Errorf("invalid guess on ended game")
	}
	if _, err := game.ValidGuess(word); err != nil {
		return nil, err
	}
	guessState := []State{}
	for pos, char := range word {
		if !strings.Contains(game.SecretWord, string(char)) {
			guessState = append(guessState, Wrong)
		} else {
			if string(char) == string(game.SecretWord[pos]) {
				guessState = append(guessState, True)
			} else {
				guessState = append(guessState, Partial)
			}
		}
	}
	game.MoveCount++
	if word == game.SecretWord {
		game.IsWin = true
	}
	return guessState, nil
}

func ValidWord(word string) bool {
	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	return IsLetter(word)
}

func (game *Game) ValidGuess(word string) (bool, error) {
	if !ValidWord(word) {
		return false, nil
	}
	if len(word) != game.GetSecretLength() {
		return false, fmt.Errorf("invalid word length, expect %d, got %d", len(word), game.GetSecretLength())
	}

	return true, nil
}

func (game *Game) IsOutOfMove() bool {
	return game.MoveCount >= maxMovesAllowed
}

func ParseGuessState(guessState []State) string {
	var result string
	for _, state := range guessState {
		switch state {
		case Partial:
			result += "P"
		case Wrong:
			result += "W"
		case True:
			result += "T"
		}
	}
	return result
}
