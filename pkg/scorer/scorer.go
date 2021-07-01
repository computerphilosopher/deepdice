package scorer

type Scorable interface {
	Num() int
}

func numCounter(scorables []Scorable) map[int]int {
	counter := map[int]int{}
	for _, scorable := range scorables {
		num := scorable.Num()
		if _, exist := counter[num]; !exist {
			counter[num] = 0
		}
		counter[num] += 1
	}
	return counter
}

func sum(scorables []Scorable) int {
	ret := 0
	for _, scorable := range scorables {
		ret += scorable.Num()
	}
	return ret
}

func ScoreSingle(scorables []Scorable, num int) int {
	return num * numCounter(scorables)[num]
}

func ScoreNOfAKind(scorables []Scorable, n int) int {
	counter := numCounter(scorables)

	for _, value := range counter {
		if value >= n {
			return sum(scorables)
		}
	}

	return 0
}

func ScoreYacht(scorables []Scorable) int {
	if ScoreNOfAKind(scorables, len(scorables)) > 0 {
		return 50
	}
	return 0
}
