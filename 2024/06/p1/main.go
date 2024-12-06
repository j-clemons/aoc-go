package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

func splitIntSlice(str string, delim string) []int {
    strSlice := strings.Split(str, delim)

    intSlice := []int{}
    for _, v := range strSlice {
        intSlice = append(intSlice, toInt(v))
    }
    return intSlice
}

func getAllIndexes(str string, char rune) []int {
    var indexes []int
    for i, r := range str {
        if r == char {
            indexes = append(indexes, i)
        }
    }
    return indexes
}

type Loc struct {
    X int
    Y int
}

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")
    maxX := len(input_lines[0])
    maxY := len(input_lines)

    dir := []Loc{
        {X:0,  Y:1},  // up
        {X:1,  Y:0},  // right
        {X:0,  Y:-1}, // down
        {X:-1, Y:0}, //left
    }
    objectMap := make(map[Loc]int)
    var startLoc Loc
    for i, l := range input_lines {
        objLocs := getAllIndexes(l, '#')
        startLocX := getAllIndexes(l, '^')

        for _, o := range objLocs {
            objectMap[Loc{X:o, Y:len(input_lines)-i-1}] = 1
        }

        if len(startLocX) > 0 {
            startLoc = Loc{X: startLocX[0], Y:len(input_lines)-i-1}
        }
    }

    visitedMap := make(map[Loc]int)
    guardLoc := startLoc
    validNext := true
    for i:=0; i>=0; i++ {
        for j:=1; j>=0; j=j+0 {
            visitedMap[guardLoc] = 1
            nextLoc := Loc{
                X: guardLoc.X + dir[i].X,
                Y: guardLoc.Y + dir[i].Y,
            }
            if nextLoc.X > -1 && nextLoc.X < maxX &&
            nextLoc.Y > -1 && nextLoc.Y < maxY {
                if objectMap[nextLoc] == 1 {
                    break
                } else {
                    guardLoc = nextLoc
                }
            } else {
                validNext = false
                break
            }
        }
        if !validNext {
            break
        } else {
            if i == 3 {
                i = -1
            }
        }
    }

    return len(visitedMap)
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Println(res)
}
