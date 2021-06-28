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

type Dice interface {
	Roll() int
	Num() int
}

type BasicDice struct {
	sides int
	num   int
}

func NewDice(sides int) Dice {
	return &BasicDice{
		sides: sides,
	}
}

func (dice *BasicDice) Roll() int {
	dice.num = rand.Intn(dice.sides) + 1
	return dice.num
}

func (dice *BasicDice) Num() int {
	return dice.num
}

func Roll(sides int, cnt int) []*BasicDice {
	ret := []*BasicDice{}
	for i := 0; i < cnt; i++ {
		ret = append(ret, &BasicDice{sides: sides})
		ret[i].Roll()
	}
	return ret
}

type DiceKeeper struct {
	dices []*BasicDice
}

func NewDiceKeeper(size int) *DiceKeeper {
	return &DiceKeeper{
		dices: make([]*BasicDice, size),
	}
}

func (keeper *DiceKeeper) IsEmpty(idx int) (bool, error) {
	if idx >= len(keeper.dices) {
		return false, makeError(outOfRangeMsg, "DiceKeeper.IsEmpty")
	}
	return keeper.dices[idx] == nil, nil
}

func (keeper *DiceKeeper) Keep(dice *BasicDice, idx int) error {
	if idx >= len(keeper.dices) {
		return makeError(outOfRangeMsg, "DiceKeeper.Keep")
	}
	keeper.dices[idx] = dice
	return nil
}

func (keeper *DiceKeeper) Pop(idx int) (*BasicDice, error) {
	ret := keeper.dices[idx]
	keeper.dices[idx] = nil
	return ret, nil
}

func (keeper *DiceKeeper) Size() int {
	return len(keeper.dices)
}
