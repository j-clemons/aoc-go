package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")

    result := 0
    ruleMap := make(map[int][]int)
    for _, l := range input_lines {
        if len(l) == 0 {
            continue
        } else if l[2] == '|' {
            rules := splitIntSlice(l, "|")
            current := ruleMap[rules[1]]
            ruleMap[rules[1]] = append(current, rules[0])
        } else if l[2] == ',' {
            update := splitIntSlice(l, ",")

            validUpdate := true
            invalidValues := []int{}
            for _, u := range update {
               if slices.Contains(invalidValues, u) {
                   validUpdate = false
                   break
               }
               invalidValues = append(invalidValues, ruleMap[u]...)
            }
            if validUpdate {
                result = result + update[int(math.Floor(float64(len(update))/2))]
            }
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
