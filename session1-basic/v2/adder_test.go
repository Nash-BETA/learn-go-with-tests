package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertCorrectMassage := func(t testing.TB, sum, expected int) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}
	t.Run("2+2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertCorrectMassage(t, sum, expected)
	})
	t.Run("1+5", func(t *testing.T) {
		sum := Add(1, 5)
		expected := 6
		assertCorrectMassage(t, sum, expected)
	})
}
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
