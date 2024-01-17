package compositions

import "testing"

func TestShapes(t *testing.T) {
	rect := Rectange{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %g want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	/* shape is kind of any-data that Implements Method/Funcs inside of it */
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rect := Rectange{10.0, 10.0}
		want := 100.0

		checkArea(t, rect, want)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})

	// t.Run("triangle", func(t *testing.T) {
	// 	triangle := Triangle{10.0, 10.0}
	// 	want := 314.1592653589793

	// 	checkArea(t, triangle, want)
	// 	/** ^? cannot use triangle (variable of type Triangle) as Shape value in argument to checkArea:
	// 	 *	Triangle does not implement Shape (missing method Area)
	// 	 */
	// })
}

func TestAreaSimplified(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{
			shape: Rectange{12, 6},
			want:  72.0,
		},
		{
			shape: Circle{10},
			want:  314.1592653589793,
		},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
