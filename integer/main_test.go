package integer

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected '%d', Result '%d'", expected, sum)
	}
}
