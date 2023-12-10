package main

import (
    "testing"
)

func TestPromptFunc(t *testing.T) {

    i := `.....
.S-7.
.|.|.
.L-J.
.....`

    var expected int = 4

    actual := promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

    i = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`
    expected = 8

    actual = promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}

func TestSwapOutputDirection(t *testing.T) {
    i := West
    expected := East

    actual := swapOutputDirection(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }
}
