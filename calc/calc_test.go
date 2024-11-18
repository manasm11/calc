package calc

import "testing"

func TestAdd(t *testing.T) {
	assertEqual(t, Add(1, 2), 3)
	assertEqual(t, Add(2, 2), 4)
	assertEqual(t, Add(2, 2, 3), 7)
}

func TestMul(t *testing.T) {
	assertEqual(t, Mul(1, 2), 2)
	assertEqual(t, Mul(2, 7), 14)
	assertEqual(t, Mul(2, 2, 3), 12)
}

func TestSub(t *testing.T) {
	assertEqual(t, Sub(1, 2), -1)
	assertEqual(t, Sub(2, 2), 0)
}

func assertEqual(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
