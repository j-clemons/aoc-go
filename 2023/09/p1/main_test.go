package main

import (
    "testing"
)

func TestPromptFunc(t *testing.T) {

    i := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

    var expected int = 114

    actual := promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
