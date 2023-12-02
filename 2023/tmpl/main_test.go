package main

import (
    "testing"
)

func TestPromptFunc(t *testing.T) {

    i := `TEST 
INPUT
LINES`

    var expected int64 = 0

    actual := promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
