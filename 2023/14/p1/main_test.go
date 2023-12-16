package main

import (
    "testing"
)

func TestPromptFunc(t *testing.T) {

    i := `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

    var expected int = 136

    actual := promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
