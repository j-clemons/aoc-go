package main

import (
    "testing"
)

func TestPromptFunc(t *testing.T) {

    i := `Time:      71530
Distance:  940200`

    var expected int = 71503

    actual := promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
