package kata

import (
	"codewars/tools"
	"testing"
)

func Move(position int, roll int) int {
    return roll * 2 + position;
}

func TestMove(t *testing.T){
	tools.CheckResult(Move(0, 4), 8, t)
	tools.CheckResult(Move(3, 6), 15, t)
	tools.CheckResult(Move(2, 5), 12, t)
}
