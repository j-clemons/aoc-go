package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    "regexp"
)

func toInt(str string) int {
    i, err := strconv.Atoi(str)
    if err != nil {
        log.Fatal(err)
    }

    return i
}

func checkGame(rawGame string) int {
    cubeCount := map[string]int{
        "red": 12,
        "green": 13,
        "blue": 14,
    }

    summaryGames := strings.Split(rawGame, ":")

    summary := summaryGames[0]
    games := summaryGames[1]

    gameResult := strings.Split(games, ";")

    for _, g := range(gameResult) {
        gameLine := strings.Split(g, ",")

        for _, gl := range(gameLine) {
            glDetails := strings.Split(strings.TrimSpace(gl), " ")

            color := strings.TrimSpace(glDetails[1])
            count := toInt(strings.TrimSpace(glDetails[0]))

            if count > cubeCount[color] {
                return 0
            }
        }
    }

    re := regexp.MustCompile("[0-9]+")
    return toInt(re.FindString(summary))
}

func promptFunc(input string) int {
    input_lines := strings.Split(input, "\n")

    returnVal := 0
    for _, ln := range(input_lines) {
        if ln != "" {
            returnVal += checkGame(ln)
        }
    }

    return returnVal
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Print(res)
}
