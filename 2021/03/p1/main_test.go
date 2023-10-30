package main

import (
    "testing"
)

func TestPowerConsumption(t *testing.T) {

    i := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

    var expected int64 = 198

    actual := powerConsumption(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
