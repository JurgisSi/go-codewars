package kata

import (
	"codewars/tools"
	"testing"
)

func MakeNegative(x int) int {
	if (x > 0) {
		return -x;
	}
	return x;
}

func TestMakeNegative(t *testing.T){
	tools.CheckResult(MakeNegative(4), -4, t)
	tools.CheckResult(MakeNegative(0), 0, t)
	tools.CheckResult(MakeNegative(-8), -8, t)
}
