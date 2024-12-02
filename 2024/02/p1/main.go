package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "strings"
    "strconv"
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

func checkDirection(v0 int, v1 int) (int, int) {
    if v0 == v1 {
        return 0, 0
    } else if v0 > v1 {
        return (v0 - v1), -1
    } else {
        return (v1 - v0), 1
    }
}

func checkDiff(diff int) bool {
    if diff > 0 && diff < 4 {
        return true
    }
    return false
}



func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")

    safe_ct := len(input_lines)
    for _, l := range input_lines {
        if len(l) == 0 {
            safe_ct = safe_ct - 1
            continue
        }
        values := strings.Split(l, " ")

        direction := 0
        for i, v := range values {
            if i == 0 {
                continue
            } else if i == 1 {
                diff, tmp_dir := checkDirection(toInt(values[i-1]), toInt(v))
                direction = tmp_dir
                if !checkDiff(diff) {
                    safe_ct = safe_ct - 1
                    break
                }
            } else {
                diff, tmp_dir := checkDirection(toInt(values[i-1]), toInt(v))
                if tmp_dir != direction || !checkDiff(diff) {
                    safe_ct = safe_ct - 1
                    break
                }
            }
        }
    }

    return safe_ct
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Println(res)
}
