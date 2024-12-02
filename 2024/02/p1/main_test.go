package main

import (
    "testing"
)

func TestPromptFunc(t *testing.T) {

    i := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

    var expected int = 2

    actual := promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
