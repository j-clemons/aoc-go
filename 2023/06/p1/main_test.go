package main

import (
    "testing"
)

func TestPromptFunc(t *testing.T) {

    i := `Time:      7  15   30
Distance:  9  40  200`

    var expected int = 288

    actual := promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
