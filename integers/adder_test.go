package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("should add", func(t *testing.T) {
		got := Add(2, 2)
		expected := 4

		if got != expected {
			t.Errorf("\nexpected '%d'\nbut got '%d'", expected, got)
		}
	})
}

func Add(first, last int) int {
	return first + last
}
