package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    "slices"
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

    return winCount
}

func promptFunc(input string) int {
    input_lines := strings.Split(input, "\n")

    cards := make(map[int]int)
    for i, ln := range(input_lines) {
        if ln != "" {
            picksWinners := strings.Split(ln, "|")

            winners := strings.Fields(picksWinners[1])
            picks := strings.Fields(strings.Split(picksWinners[0], ":")[1])

            points := checkPoints(picks, winners)
            cards[i] += 1
            if points > 0 {
                for k := 0; k < cards[i]; k++ {
                    for j := 1; j <= points; j++ {
                        cards[i+j] += 1
                    }
                }
            }
        }
    }

    totalCards := 0
    for _, v := range cards {
        totalCards += v
    }

    return totalCards
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Print(res)
}
