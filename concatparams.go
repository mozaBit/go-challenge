package piscine

func ConcatParams(args []string) string {
	var s string
	for j := 0; j < len(args)-1; j++ {
		s = s + args[j] + "\n"
	}
	s = s + args[len(args)-1]
	return s
}
