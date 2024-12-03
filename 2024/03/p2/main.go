package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

func calcMul(mul string) int {
    re := regexp.MustCompile(`[0-9]+`)
    factors := re.FindAllString(mul, -1)

    return toInt(factors[0]) * toInt(factors[1])
}

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    updatedInput := strings.ReplaceAll(input, "\n", "")
    enableInstructions := regexp.MustCompile(`(^|do\(\))(.*?)(don't\(\)|$)`)
    enabled := enableInstructions.FindAllString(updatedInput, -1)

    re := regexp.MustCompile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)`)
    result := 0
    for _, e := range enabled {
        mul := re.FindAllString(e, -1)

        for _, m := range mul {
            result = result + calcMul(m)
        }
    }

    return result
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Println(res)
}
