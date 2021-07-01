package scorer_test

import (
	"github.com/computerphilosopher/deepdice/pkg/scorer"
	"github.com/stretchr/testify/assert"
	"testing"
)

type FakeDice struct {
	num int
}

func (dice FakeDice) Num() int {
	return dice.num
}

func TestSingle(t *testing.T) {
	dices := []scorer.Scorable{
		FakeDice{6},
		FakeDice{1},
		FakeDice{1},
		FakeDice{2},
		FakeDice{3},
	}

	assert.Equal(t, 2, scorer.ScoreSingle(dices, 1))
	assert.Equal(t, 2, scorer.ScoreSingle(dices, 2))
	assert.Equal(t, 3, scorer.ScoreSingle(dices, 3))
	assert.Equal(t, 0, scorer.ScoreSingle(dices, 4))
	assert.Equal(t, 0, scorer.ScoreSingle(dices, 5))
	assert.Equal(t, 6, scorer.ScoreSingle(dices, 6))
}

func TestNOfAKind(t *testing.T) {
	dices := [][]scorer.Scorable{
		[]scorer.Scorable{
			FakeDice{1},
			FakeDice{1},
			FakeDice{2},
			FakeDice{1},
			FakeDice{3},
		},
		[]scorer.Scorable{
			FakeDice{3},
			FakeDice{3},
			FakeDice{3},
			FakeDice{1},
			FakeDice{3},
		},
	}

	assert.Equal(t, 8, scorer.ScoreNOfAKind(dices[0], 3))
	assert.Equal(t, 0, scorer.ScoreNOfAKind(dices[0], 4))
	assert.Equal(t, 13, scorer.ScoreNOfAKind(dices[1], 4))
	assert.Equal(t, 0, scorer.ScoreNOfAKind(dices[0], 5))
}

func TestYacht(t *testing.T) {
	dices := [][]scorer.Scorable{
		[]scorer.Scorable{
			FakeDice{5},
			FakeDice{5},
			FakeDice{5},
			FakeDice{5},
			FakeDice{5},
		},
		[]scorer.Scorable{
			FakeDice{3},
			FakeDice{3},
			FakeDice{3},
			FakeDice{1},
			FakeDice{3},
		},
	}

	assert.Equal(t, 50, scorer.ScoreYacht(dices[0]))
	assert.Equal(t, 0, scorer.ScoreYacht(dices[1]))
}
