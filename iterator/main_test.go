package iterator

import "testing"

func TestRepetition(t *testing.T) {
	repetitions := Repetition("a", 5)
	expected := "aaaaa"

	if repetitions != expected {
		t.Errorf("Expected('%s') but return ('%s')", expected, repetitions)
	}
}

func BenchmarkRepetition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetition("a", 5)
	}
}
