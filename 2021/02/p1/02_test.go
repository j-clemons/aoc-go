package main

import (
    "testing"
)

func TestSubPos(t *testing.T) {

    i := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

    expected := 150

    actual := subPos(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
