package tools

import "testing"

func CheckResult(got interface{}, want interface{}, t *testing.T) {
    if got != want {
        t.Errorf("got %v wanted %v", got, want)
    }
}