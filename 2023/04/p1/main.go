package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    "slices"
    "math"
)

func toInt(str string) int {
    i, err := strconv.Atoi(str)
    if err != nil {
        log.Fatal(err)
    }

    return i
}

func checkPoints(pick []string, winners []string) int {
    winCount := 0
    for _, v := range(pick) {
        if slices.Contains(winners, v) {
            winCount += 1
        }
    }
    
    if winCount == 0 {
        return 0
    } else {
        return int(math.Pow(2, float64(winCount-1)))
    }
}

func promptFunc(input string) int {
    input_lines := strings.Split(input, "\n")

    points := 0
    for _, ln := range(input_lines) {
        if ln != "" {
            picksWinners := strings.Split(ln, "|")

            winners := strings.Fields(picksWinners[1])
            picks := strings.Fields(strings.Split(picksWinners[0], ":")[1])

            points += checkPoints(picks, winners)
        }
    }

    return points
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Print(res)
}
