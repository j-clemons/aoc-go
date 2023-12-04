package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

func toInt(str string) int {
    i, err := strconv.Atoi(str)
    if err != nil {
        log.Fatal(err)
    }

    return i
}

type Game struct {
    Red   int
    Blue  int
    Green int
}

func checkGame(rawGame string) int {
    minGameReq := Game{
        Red: 0,
        Blue: 0,
        Green: 0,
    }

    summaryGames := strings.Split(rawGame, ":")

    // summary := summaryGames[0]
    games := summaryGames[1]

    gameResult := strings.Split(games, ";")

    for _, g := range(gameResult) {
        gameLine := strings.Split(g, ",")

        game := Game{
            Red: 0,
            Blue: 0,
            Green: 0,
        }

        for _, gl := range(gameLine) {
            glDetails := strings.Split(strings.TrimSpace(gl), " ")

            color := strings.TrimSpace(glDetails[1])
            count := toInt(strings.TrimSpace(glDetails[0]))

            switch color {
                case "red":
                    game.Red = count
                case "blue":
                    game.Blue = count
                case "green":
                    game.Green = count
            }
        }

        if game.Red > minGameReq.Red {
            minGameReq.Red = game.Red
        }
        if game.Blue > minGameReq.Blue {
            minGameReq.Blue = game.Blue
        }
        if game.Green > minGameReq.Green {
            minGameReq.Green = game.Green
        }
    }

    powerSum := minGameReq.Red * minGameReq.Blue * minGameReq.Green

    return powerSum
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
