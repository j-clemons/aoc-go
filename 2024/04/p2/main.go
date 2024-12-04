package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"time"

	"strconv"
	"strings"
)

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func toInt(str string) int {
    i, err := strconv.Atoi(str)
    if err != nil {
        log.Fatal(err)
    }

    return i
}

func checkDirections(row int, col int, ws []string ) int {
    diag1 := [2][2]int{
        {-1,1},  // up left diag
        {1,-1},  // down right diag
    }

    diag2 := [2][2]int{
        {1,1},   // up right diag
        {-1,-1}, // down left diag
    }

    maxRow := len(ws) - 1
    maxCol := len(ws[0]) - 1
    chars := []byte{'M', 'S'}

    if row > 0 && row < maxRow && col > 0 && col < maxCol {
        // check for mas
        if slices.Contains(chars, ws[row+diag1[0][1]][col+diag1[0][0]]) &&
        slices.Contains(chars, ws[row+diag1[1][1]][col+diag1[1][0]]) &&
        ws[row+diag1[0][1]][col+diag1[0][0]] != ws[row+diag1[1][1]][col+diag1[1][0]] &&
        slices.Contains(chars, ws[row+diag2[0][1]][col+diag2[0][0]]) &&
        slices.Contains(chars, ws[row+diag2[1][1]][col+diag2[1][0]]) &&
        ws[row+diag2[0][1]][col+diag2[0][0]] != ws[row+diag2[1][1]][col+diag2[1][0]] {
            return 1
        }
    }

    return 0
}

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    raw_input_lines := strings.Split(input, "\n")

    input_lines := []string{}

    for _, l := range raw_input_lines {
        if len(l) == 0 {
            continue
        }
        input_lines = append(input_lines, l)
    }

    result := 0
    reX := regexp.MustCompile("A")
    for i, l := range input_lines {
        xLocs := reX.FindAllIndex([]byte(l), -1)
        for _, x := range xLocs {
            loc := x[0]
            result = result + checkDirections(i, loc, input_lines)
        }
    }

    return result
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Println(res)
}
