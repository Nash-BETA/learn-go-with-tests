package iteration

import "testing"

func TestName(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but repeated %q", expected, repeated)
	}
}
