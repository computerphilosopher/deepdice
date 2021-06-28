package dice_test

import (
	"github.com/computerphilosopher/deepdice/pkg/dice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDice(t *testing.T) {
	d := dice.NewDice(1)
	assert.Equal(t, 1, d.Roll())

	d = dice.NewDice(100)
	num := d.Roll()
	assert.GreaterOrEqual(t, num, 1)
	assert.LessOrEqual(t, num, 100)
	assert.Equal(t, num, d.Num())

	dices := dice.Roll(100, 5)
	assert.Equal(t, 5, len(dices))
	for _, d := range dices {
		assert.GreaterOrEqual(t, d.Num(), 1)
		assert.LessOrEqual(t, d.Num(), 100)
	}
}

func TestKeeper(t *testing.T) {
	keeper := dice.NewDiceKeeper(5)
	dices := dice.Roll(100, 5)

	assert.Equal(t, 5, keeper.Size())

	err := keeper.Keep(dices[0], 6)
	assert.NotNil(t, err)

	isEmpty, err := keeper.IsEmpty(0)
	assert.Nil(t, err)
	assert.True(t, isEmpty)

	err = keeper.Keep(dices[0], 0)
	assert.Nil(t, err)

	d, err := keeper.Pop(0)
	assert.Equal(t, d, dices[0])
	isEmpty, _ = keeper.IsEmpty(0)
	assert.True(t, isEmpty)
}
