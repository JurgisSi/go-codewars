package kata

import (
	"codewars/tools"
	"fmt"
	"testing"
)

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name);
}

func TestGreet(t *testing.T){
	tools.CheckResult(Greet("Billy"), "Hello, Billy how are you doing today?", t)
	tools.CheckResult(Greet(""), "Hello,  how are you doing today?", t)
}
