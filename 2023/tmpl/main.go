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

func promptFunc(input string) int {
    input_lines := strings.Split(input, "\n")

    return nil
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Print(res)
}
