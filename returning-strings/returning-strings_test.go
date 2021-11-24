package kata

import (
	"fmt"
	"testing"
)

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name);
}

func TestMakeNegative(t *testing.T){
	Check(Greet("Billy"), "Hello, Billy how are you doing today?", t)
	Check(Greet(""), "Hello,  how are you doing today?", t)
}

func Check(got string, want string, t *testing.T) {
    if got != want {
        t.Errorf("got %s wanted %s", got, want)
    }
}