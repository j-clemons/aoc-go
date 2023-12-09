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

func strIntSlice(strSlice []string) []int {
    intSlice := make([]int, len(strSlice))
    for i, s := range strSlice {
        intSlice[i], _ = strconv.Atoi(s)
    }

    return intSlice
}

func sliceOnlyZero(slc []int) bool {
    for _, v := range slc {
        if v != 0 {
            return false
        }
    }
    return true
}

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")

    endVal := 0
    for _, ln := range input_lines {
        if ln == "" {
            continue
        }
        seq := strIntSlice(strings.Fields(ln))

        lnSlice := [][]int{seq}
        for 1 < 2 {
            lnSl := []int{}
            for i := 1; i < len(seq); i++ {
                diff := seq[i] - seq[i-1]
                lnSl = append(lnSl, diff)
            }
            if sliceOnlyZero(lnSl) {
                break
            } else {
                lnSlice = append(lnSlice, lnSl)
                seq = lnSl
            }
        }

        for i := 0; i < len(lnSlice); i++ {
            fs := lnSlice[i]
            endVal += fs[len(fs)-1]
        }
    }

    return endVal
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Println(res)
}
