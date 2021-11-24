package kata

import (
	"testing"
)

func MakeNegative(x int) int {
	if (x > 0) {
		return -x;
	}
	return x;
}

func TestMakeNegative(t *testing.T){
	Check(MakeNegative(4), -4, t)
	Check(MakeNegative(0), 0, t)
	Check(MakeNegative(-8), -8, t)
}

func Check(got int, want int, t *testing.T) {
    if got != want {
        t.Errorf("got %d, wanted %d", got, want)
    }
}