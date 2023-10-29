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

type direction struct {
    Dir string
    Val int
}

func subPos(input string) int {
    input_lines := strings.Split(input, "\n")

    lines := make([]direction, len(input_lines))

    for i := range input_lines {
        if input_lines[i] != "" {
            line_val := strings.Split(input_lines[i], " ")
            lines[i] = direction{
                Dir: strings.TrimSpace(line_val[0]),
                Val: toInt(strings.TrimSpace(line_val[1])),
            }
        }
    }

    h := 0
    d := 0
    a := 0
    for i := 0; i < len(lines); i++ {
        if lines[i].Dir == "forward" {
            h = h + lines[i].Val
            d = d + (a * lines[i].Val)
        } else if lines[i].Dir == "down" {
            // d = d + lines[i].Val
            a = a + lines[i].Val
        } else if lines[i].Dir == "up" {
            // d = d - lines[i].Val
            a = a - lines[i].Val
        }
        // fmt.Printf("h %d; d %d; a %d:: ", h, d, a)
    }

    // fmt.Printf("h %d; d %d", h, d)
    return h * d
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := subPos(string(file))

    fmt.Print(res)
}
