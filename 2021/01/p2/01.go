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

func depthCheck(input string) int {
    input_lines := strings.Split(input, "\n")

    lines := make([]int, len(input_lines))

    for i := range input_lines {
        str := strings.TrimSpace(input_lines[i])
        if str != "" {
            lines[i] = toInt(str)
        }
    }

    result := 0
    for i := 1; i < len(lines) - 2; i++ {
        a := lines[i-1] + lines[i] + lines[i+1]
        b := lines[i] + lines[i+1] + lines[i+2]
        if b > a {
            result = result + 1
        }
    }

    return result
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := depthCheck(string(file))

    fmt.Print(res)
}
