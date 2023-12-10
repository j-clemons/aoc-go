package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "strings"
    "strconv"
    "slices"
    "math"
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

type Direction int

const (
    North Direction = iota
    East
    South
    West
)

type Location struct {
    Row int
    Col int
}

type Pipe struct {
    D1 Direction
    D2 Direction
}

var PipeTypes = []rune{
    '|',
    '-',
    'L',
    'J',
    '7',
    'F',
}

var Pipes = map[rune]Pipe{
    '|': {North, South},
    '-': {East, West},
    'L': {North, East},
    'J': {North, West},
    '7': {South, West},
    'F': {South, East},
}

var LocationOffset = map[Direction]Location{
    North: {-1,0},
    East: {0,1},
    South: {1,0},
    West: {0,-1},
}

func (l *Location) addLocation(l2 Location) Location {
    return Location{
        l.Row + l2.Row,
        l.Col + l2.Col,
    }
}

func swapOutputDirection(d Direction) Direction {
    switch d {
    case North:
        return South
    case South:
        return North
    case East:
        return West
    case West:
        return East
    default:
        return d
    }
}

func getStartPipe(l Location, input_lines []string) (rune, Direction, Location) {
    opts := []Direction{
        North,
        East,
        South,
        West,
    }

    for _, v := range opts {
        ofs := l.addLocation(LocationOffset[v])
        if ofs.Row < 0 || ofs.Col < 0 {
            continue
        }

        checkVal := rune(input_lines[ofs.Row][ofs.Col])
        if slices.Contains(PipeTypes, checkVal) {
            return checkVal, swapOutputDirection(v), ofs
        }
    }

    return rune('_'), North, l
}

func getNextPipeLoc(p rune, loc Location, prevDir Direction) (Direction, Location) {
    dirs := Pipes[p]

    if prevDir == dirs.D1 {
        // swapping output direction so it is correct input for next pipe
        // but using the original output direction for the offset
        return swapOutputDirection(dirs.D2), loc.addLocation(LocationOffset[dirs.D2])
    } else {
        return swapOutputDirection(dirs.D1), loc.addLocation(LocationOffset[dirs.D1])
    }
}

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")

    start := Location{}
    for i, ln := range input_lines {
        startIdx := strings.Index(ln, "S")
        if startIdx != -1 {
            start = Location{Row: i, Col: startIdx}
            break
        }
    }

    pipe, dir, loc := getStartPipe(start, input_lines)

    steps := 1
    for 0<1 {
        dir, loc = getNextPipeLoc(pipe, loc, dir)
        pipe = rune(input_lines[loc.Row][loc.Col])

        if pipe == 'S' {
            break
        } else {
            steps += 1
        }
    }

    if steps % 2 == 1 {
        return int(math.Ceil(float64(steps)/2))
    } else {
        return steps / 2
    }
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Println(res)
}
