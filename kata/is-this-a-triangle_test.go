package kata

import (
	"codewars/tools"
	"testing"
)

func IsTriangle(a, b, c int) bool {
	if (isAllPositive(a, b,c)) {
		return a + b > c && a + c > b && b + c > a
	}
	return false
}

func isAllPositive(a int, b int, c int) bool {
	return a > 0 && b > 0 && c > 0
}

func TestIsTriangle(t *testing.T){
	tools.CheckResult(IsTriangle(5, 1, 2), false, t)
	tools.CheckResult(IsTriangle(1, 2, 5), false, t)
	tools.CheckResult(IsTriangle(2, 5, 1), false, t)
	tools.CheckResult(IsTriangle(4, 2, 3), true, t)
	tools.CheckResult(IsTriangle(5, 1, 5), true, t)
	tools.CheckResult(IsTriangle(2, 2, 2), true, t)
	tools.CheckResult(IsTriangle(-1, 2, 3), false, t)
}
