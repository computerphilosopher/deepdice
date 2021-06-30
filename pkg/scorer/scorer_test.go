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
