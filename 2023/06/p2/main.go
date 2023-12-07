package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "strings"
    "strconv"
    "math"
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

func numWaysToBeatRecord(time int, record int) int {
    t := float64(time)
    r := float64(record)
    s1 := 0.5 * (math.Sqrt(math.Pow(t,2.0) - (4*r)) + t)
    s0 := t - s1

    ways := math.Ceil(s1) - math.Ceil(s0)

    if math.Trunc(s1) == s1 {
        return int(ways) - 1
    } else {
        return int(ways) 
    }
}

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")

    timeValue := toInt(strings.ReplaceAll(strings.Split(input_lines[0], ":")[1], " ", ""))
    recordValue := toInt(strings.ReplaceAll(strings.Split(input_lines[1], ":")[1], " ", ""))

    return numWaysToBeatRecord(timeValue, recordValue)
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Print(res)
}
