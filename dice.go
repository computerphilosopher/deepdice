package dice

import (
	"fmt"
	"math/rand"
)

const (
	outOfRangeMsg = "Out of range of slice or array"
)

func makeError(msg string, context string) error {
	return fmt.Errorf("%s: %s", msg, context)
}

type Dice struct {
	sides int
	num   int
}

func NewDice(sides int) *Dice {
	return &Dice{
		sides: sides,
	}
}

func (dice *Dice) Roll() int {
	dice.num = rand.Intn(dice.sides) + 1
	return dice.num
}

func (dice *Dice) Num() int {
	return dice.num
}

func Roll(sides int, cnt int) []*Dice {
	ret := []*Dice{}
	for i := 0; i < cnt; i++ {
		ret = append(ret, NewDice(sides))
		ret[i].Roll()
	}
	return ret
}

type DiceKeeper struct {
	dices []*Dice
}

func NewDiceKeeper(size int) *DiceKeeper {
	return &DiceKeeper{
		dices: make([]*Dice, size),
	}
}

func (keeper *DiceKeeper) IsEmpty(idx int) (bool, error) {
	if idx >= len(keeper.dices) {
		return false, makeError(outOfRangeMsg, "DiceKeeper.IsEmpty")
	}
	return keeper.dices[idx] == nil, nil
}

func (keeper *DiceKeeper) Keep(dice *Dice, idx int) error {
	if idx >= len(keeper.dices) {
		return makeError("Dicekeeper.Keep", outOfRangeMsg)
	}
	keeper.dices[idx] = dice
	return nil
}

func (keeper *DiceKeeper) Pop(idx int) (*Dice, error) {
	ret := keeper.dices[idx]
	keeper.dices[idx] = nil
	return ret, nil
}

func (keeper *DiceKeeper) Size() int {
	return len(keeper.dices)
}
