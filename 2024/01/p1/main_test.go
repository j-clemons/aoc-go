package main

import (
    "testing"
)

func TestPromptFunc(t *testing.T) {

    i := `3   4
4   3
2   5
1   3
3   9
3   3`

    var expected int = 11

    actual := PromptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
