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

func reverse(s string) string {
    var str string
    for i := len(s)-1; i > -1; i-- {
        str += string(s[i])
    }

    return str
}

func intFromString(str string) int {
    reStr := "one|two|three|four|five|six|seven|eight|nine"
    re := regexp.MustCompile("[0-9]|"+reStr)
    match := re.FindString(str)

    numMap := map[string]string{
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
    }

    reverseNumMap := map[string]string{}
    for k, v := range numMap {
        reverseNumMap[reverse(k)] = v
    }
    reReverse := regexp.MustCompile("[0-9]|"+reverse(reStr))
    matchReverse := reReverse.FindString(reverse(str))

    
    v0, ok := numMap[match]
    if !ok {
        v0 = match
    }
    v1, ok := reverseNumMap[matchReverse]
    if !ok {
        v1 = matchReverse
    }

    outputStr := v0 + v1

    return toInt(outputStr)
}

func promptFunc(input string) int {
    inputLines := strings.Split(input, "\n")

    total := 0
    for _, v := range(inputLines) {
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
