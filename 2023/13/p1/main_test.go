package main

import (
    "testing"
)

func TestPromptFunc(t *testing.T) {

//     i := `#.##..##.
// ..#.##.#.
// ##......#
// ##......#
// ..#.##.#.
// ..##..##.
// #.#.##.#.
//
// #...##..#
// #....#..#
// ..##..###
// #####.##.
// #####.##.
// ..##..###
// #....#..#

    i := 
`
.#.##.##
#.#.#.##
#..##.##
`
    // var expected int = 411
    var expected int = 6

    actual := promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }

}
