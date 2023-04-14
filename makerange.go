package piscine

func MakeRange(min, max int) []int {
	var aSlice []int
	if max > min {
		aSlice = make([]int, max-min)
		for i := 0; i < max-min; i++ {
			aSlice[i] = min + i
		}
		return aSlice
	}
	return aSlice
}
