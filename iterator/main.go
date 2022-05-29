package iterator

func Repetition(s string, limit int) string {
	result := ""
	for i := 0; i < limit; i++ {
		result += s
	}
	return result
}
