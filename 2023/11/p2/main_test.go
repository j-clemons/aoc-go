package main

import (
    "testing"
)

func TestPromptFunc(t *testing.T) {

    i := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

    var expected int = 8410

    actual := promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
