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

func swapInvalidValue(idx int, val int, rules []int, update []int) []int {

    newUpdate := make([]int, len(update))
    copy(newUpdate, update)
    for i := 0; i < len(update); i++ {
        if slices.Contains(rules, update[i]) {
            newUpdate[i] = val
            newUpdate[idx] = update[i]
            break
        }
    }
    return newUpdate
}

func checkUpdate(update []int, rules map[int][]int) bool {
    return true
}

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")

    result := 0
    ruleMap := make(map[int][]int)
    ascRuleMap := make(map[int][]int)
    for _, l := range input_lines {
        if len(l) == 0 {
            continue
        } else if l[2] == '|' {
            rules := splitIntSlice(l, "|")
            current := ruleMap[rules[1]]
            ruleMap[rules[1]] = append(current, rules[0])
            currentAsc := ascRuleMap[rules[0]]
            ascRuleMap[rules[0]] = append(currentAsc, rules[1])
        } else if l[2] == ',' {
            update := splitIntSlice(l, ",")

            validUpdate := true
            invalidValues := []int{}
            fixedUpdate := make([]int, len(update))
            copy(fixedUpdate, update)
            for i := 0; i < len(fixedUpdate); i++ {
               if slices.Contains(invalidValues, fixedUpdate[i]) {
                   validUpdate = false
                   fixedUpdate = swapInvalidValue(i, fixedUpdate[i], ascRuleMap[fixedUpdate[i]], fixedUpdate)
                   i = -1
                   invalidValues = []int{}
               } else {
                   invalidValues = append(invalidValues, ruleMap[fixedUpdate[i]]...)
               }
            }
            if !validUpdate {
                result = result + fixedUpdate[int(math.Floor(float64(len(fixedUpdate))/2))]
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
