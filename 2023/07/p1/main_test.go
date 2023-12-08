package main

import (
    "testing"
)

func TestPromptFunc(t *testing.T) {

    i := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

    var expected int = 6440

    actual := promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
