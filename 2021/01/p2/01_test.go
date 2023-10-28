package main

import (
    "testing"
)

func TestDepthCheck(t *testing.T) {

    i := `199
    200
    208
    210
    200
    207
    240
    269
    260
    263`

    expected := 5

    actual := depthCheck(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
