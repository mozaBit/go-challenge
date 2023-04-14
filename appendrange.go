package piscine

func AppendRange(min, max int) []int {
	var aSlice []int
	if min < max {
		for i := min; i < max; i++ {
			aSlice = append(aSlice, i)
		}
	}
	return aSlice
}
