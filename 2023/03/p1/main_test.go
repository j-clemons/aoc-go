package main

import (
    "testing"
)

func TestPromptFunc(t *testing.T) {

    i := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

    var expected int = 4361

    actual := promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
