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

func createListThenSort(lines []string) int {
    var l0 []int
    l1 := map[int]int{}

    for _, l := range lines {
        if len(l) == 0 {
            continue
        }
        values := strings.Split(l, "   ")

        l0 = append(l0, toInt(values[0]))
        l1Val := toInt(values[1])
        if val, ok := l1[l1Val]; ok {
            l1[l1Val] = val + 1
        } else {
            l1[l1Val] = 1
        }
    }

    result := 0
    for _, v := range l0 {
        result = result + (v * l1[v])
    }

    return int(result)

}

func PromptFunc(input string) int {
    defer duration(track("PromptFunc"))
    input_lines := strings.Split(input, "\n")

    return createListThenSort(input_lines)
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := PromptFunc(string(file))

    fmt.Println(res)
}
