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

func ScoreNOfAKind(scorables []Scorable, n int) int {
	counter := map[int]int{}
	sum := 0

	for _, scorable := range scorables {
		num := scorable.Num()
		if _, exist := counter[num]; !exist {
			counter[num] = 0
		}
		counter[num] += 1
		sum += num
	}

	for _, value := range counter {
		if value >= n {
			return sum
		}
	}

	return 0
}
