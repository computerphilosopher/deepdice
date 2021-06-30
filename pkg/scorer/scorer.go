package scorer

type Scorable interface {
	Num() int
}

func ScoreSingle(scorables []Scorable, num int) int {
	ret := 0
	for _, scorable := range scorables {
		if scorable.Num() == num {
			ret += num
		}
	}
	return ret
}
