package main

import (
    "testing"
)

func TestPromptFunc(t *testing.T) {

    i := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

    var expected int = 142

    actual := promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
