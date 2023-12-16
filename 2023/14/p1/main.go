package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "strings"
)

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

type RockStack struct {
    MaxValue  int
    RockCount int
}

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")

    lenInputLines := 0
    for i := range input_lines {
        if input_lines[i] != "" {
            lenInputLines += 1
        }
    }

    colMap := map[int]int{}
    for i := 0; i < len(input_lines[0]); i++ {
        colMap[i] = lenInputLines
    }

    resultMap := map[string]RockStack{}
    for i, ln := range input_lines[:lenInputLines] {
        rowValue := lenInputLines - i
        for k, v := range colMap {
            if ln[k] == 'O' {
                key := fmt.Sprint(k)+"-"+fmt.Sprint(v)
                copy := resultMap[key]
                copy.MaxValue = v
                copy.RockCount += 1
                resultMap[key] = copy
            } else if ln[k] == '#' {
                colMap[k] = rowValue - 1
            }
        }
    }


    results := 0
    for _, v := range resultMap {
        topValue := v.MaxValue
        for i := v.RockCount; i > 0; i-- {
            value := (topValue)
            results += value
            topValue -= 1
        }
    }

    return results
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Println(res)
}
