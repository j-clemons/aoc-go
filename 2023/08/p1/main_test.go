package main

import (
    "testing"
)

func TestPromptFunc(t *testing.T) {

    i := `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

    var expected int = 2

    actual := promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

    i = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

    expected = 6

    actual = promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }
}
