package rules_test

import (
	"testing"

	"github.com/TropicalDog17/wordle/x/wordle/rules"
	"github.com/stretchr/testify/assert"
)

func TestHappyPath(t *testing.T) {
	game := rules.New("inside")
	assert.Equal(t, 6, game.GetSecretLength())
	// All letter are wrong
	state, err := game.Guess("malloy")
	assert.Equal(t, rules.GuessState{rules.Wrong, rules.Wrong, rules.Wrong, rules.Wrong, rules.Wrong, rules.Wrong}, state)
	assert.NoError(t, err)
	assert.Equal(t, game.MoveCount, 1)

	// All letter are wrong order
	state, err = game.Guess("niised")
	assert.Equal(t, rules.GuessState{rules.Partial, rules.Partial, rules.Partial, rules.Partial, rules.Partial, rules.Partial}, state)
	assert.NoError(t, err)
	assert.Equal(t, game.MoveCount, 2)

	// Some mixed cases
	state, err = game.Guess("thrive")
	assert.Equal(t, rules.GuessState{rules.Wrong, rules.Wrong, rules.Wrong, rules.True, rules.Wrong, rules.True}, state)
	assert.NoError(t, err)
	assert.Equal(t, game.MoveCount, 3)

	state, err = game.Guess("melodi")
	assert.Equal(t, rules.GuessState{rules.Wrong, rules.Partial, rules.Wrong, rules.Wrong, rules.True, rules.Partial}, state)
	assert.NoError(t, err)
	assert.Equal(t, game.MoveCount, 4)

	state, err = game.Guess("insidy")
	assert.Equal(t, rules.GuessState{rules.True, rules.True, rules.True, rules.True, rules.True, rules.Wrong}, state)
	assert.NoError(t, err)
	assert.Equal(t, game.MoveCount, 5)
	// All letter are correct
	state, err = game.Guess("inside")
	assert.Equal(t, rules.GuessState{rules.True, rules.True, rules.True, rules.True, rules.True, rules.True}, state)
	assert.NoError(t, err)
	assert.Equal(t, game.MoveCount, 6)
	assert.Equal(t, game.IsWin, true)
	// Shouldn't be able to guess
	state, err = game.Guess("dsfdss")
	assert.Equal(t, rules.GuessState{}, state)
	assert.Error(t, err)
	assert.Equal(t, game.IsOutOfMove(), true)
}

func TestWinWithoutReachLastMove(t *testing.T) {
	game := rules.New("inside")
	assert.Equal(t, 6, game.GetSecretLength())
	// All letter are wrong
	state, err := game.Guess("malloy")
	assert.Equal(t, rules.GuessState{rules.Wrong, rules.Wrong, rules.Wrong, rules.Wrong, rules.Wrong, rules.Wrong}, state)
	assert.NoError(t, err)
	assert.Equal(t, game.MoveCount, 1)

	// All letter are wrong order
	state, err = game.Guess("niised")
	assert.Equal(t, rules.GuessState{rules.Partial, rules.Partial, rules.Partial, rules.Partial, rules.Partial, rules.Partial}, state)
	assert.NoError(t, err)
	assert.Equal(t, game.MoveCount, 2)

	// Some mixed cases
	state, err = game.Guess("thrive")
	assert.Equal(t, rules.GuessState{rules.Wrong, rules.Wrong, rules.Wrong, rules.True, rules.Wrong, rules.True}, state)
	assert.NoError(t, err)
	assert.Equal(t, game.MoveCount, 3)

	state, err = game.Guess("melodi")
	assert.Equal(t, rules.GuessState{rules.Wrong, rules.Partial, rules.Wrong, rules.Wrong, rules.True, rules.Partial}, state)
	assert.NoError(t, err)
	assert.Equal(t, game.MoveCount, 4)

	// All letter are correct
	state, err = game.Guess("inside")
	assert.Equal(t, rules.GuessState{rules.True, rules.True, rules.True, rules.True, rules.True, rules.True}, state)
	assert.NoError(t, err)
	assert.Equal(t, game.MoveCount, 5)
	assert.Equal(t, game.IsWin, true)

	// Shouldn't be able to guess
	state, err = game.Guess("dsfdss")
	assert.Equal(t, rules.GuessState{}, state)
	assert.Error(t, err)
	assert.Equal(t, game.MoveCount, 5)
}

// func TestInvalidSecretWord(t *testing.T) {
// 	// Contains numbers
// 	_, err := rules.New("abc2kd")
// 	assert.Error(t, err)
// 	// Contains special character
// 	_, err = rules.New("a.fd#kd")
// 	assert.Error(t, err)
// 	// Correct
// 	_, err = rules.New("fdlskjfldsjf")
// 	assert.NoError(t, err)
// }
