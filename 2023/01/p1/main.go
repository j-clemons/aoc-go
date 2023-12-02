package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    "unicode"
)

func toInt(str string) int {
    i, err := strconv.Atoi(str)
    if err != nil {
        log.Fatal(err)
    }

    return i
}

func intFromString(str string) int {
    var il []int

    r := []rune(str)
    for i, v:= range(r) {
        if unicode.IsDigit(v) {
            il = append(il, i)
        }
    }

    p1 := string(str[il[0]])
    p2 := string(str[il[len(il)-1]])

    outputStr := p1 + p2

    return toInt(outputStr)
}

func promptFunc(input string) int {
    input_lines := strings.Split(input, "\n")

    total := 0
    for _, v := range(input_lines) {
        if len(v) > 0 {
            val := intFromString(v)
            total += val
        }
    }

    return total
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Print(res)
}
