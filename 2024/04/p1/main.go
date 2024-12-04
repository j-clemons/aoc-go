package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"strings"
	"strconv"
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
    dir := [8][2]int{
        {-1,0},  // left
        {-1,1},  // up left diag
        {0,1},   // up
        {1,1},   // up right diag
        {1,0},   // right
        {1,-1},  // down right diag
        {0,-1},  // down
        {-1,-1}, // down left diag
    }

    maxRow := len(ws) - 1
    maxCol := len(ws[0]) - 1

    result := 0
    for _, d := range dir {
        mLoc := []int{row+(d[1]*1), col+(d[0]*1)}
        aLoc := []int{row+(d[1]*2), col+(d[0]*2)}
        sLoc := []int{row+(d[1]*3), col+(d[0]*3)}

        if sLoc[0] < 0 || sLoc[0] > maxRow ||
        sLoc[1] < 0 || sLoc[1] > maxCol {
            continue
        } else if ws[mLoc[0]][mLoc[1]] == 'M' &&
        ws[aLoc[0]][aLoc[1]] == 'A' &&
        ws[sLoc[0]][sLoc[1]] == 'S' {
            result++
        }
    }

    return result
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
    reX := regexp.MustCompile("X")
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
