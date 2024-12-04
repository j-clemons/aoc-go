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

func testSlice(sl []string) (bool, int) {
    direction := 0
    for i, v := range sl {
        if i == 0 {
            continue
        } else if i == 1 {
            diff, tmp_dir := checkDirection(toInt(sl[i-1]), toInt(v))
            direction = tmp_dir
            if !checkDiff(diff) {
                return false, i
            }
        } else {
            diff, tmp_dir := checkDirection(toInt(sl[i-1]), toInt(v))
            if tmp_dir != direction || !checkDiff(diff) {
                return false, i
            }
        }
    }

    return true, -1
}



func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")

    safe_ct := 0
    for _, l := range input_lines {
        if len(l) == 0 {
            continue
        }
        values := strings.Split(l, " ")

        result, _ := testSlice(values)
        if !result {

            for i := 0; i < len(values); i++ {
                sl := make([]string, len(values)-1)
                copy(sl[:i], values[:i])
                copy(sl[i:], values[i+1:])

                addtlPass, _ := testSlice(sl)
                if addtlPass {
                    safe_ct++
                    break
                }
            }
        } else if result {
            safe_ct++
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
