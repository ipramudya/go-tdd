package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("should repeat", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		if repeated != expected {
			t.Errorf("expected '%q', but got '%q'", expected, repeated)
		}
	})
}

func Repeat(character string, count int) string {
	var r string

	for i := 1; i <= count; i++ {
		r = r + character
	}

	return r
}

func BenchRepeat(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		Repeat("a", 3)
	}
}
