package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "strings"
    "strconv"
    "slices"
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

func checkIndexes(lines []string) int {
    var originalLoc int
    // get first location by processing without fixing smudge
    for i := 1; i < len(lines); i++ {
        compResult := compareRows(lines, i, false)
        if compResult == true {
            originalLoc = i
        }
    }

    // look for reflection line but skip known first location
    // fix smudge
    for i := 1; i < len(lines); i++ {
        if i == originalLoc {
            continue
        }
        compResult := compareRows(lines, i, true)
        if compResult == true {
            return i
        }
    }

    return -1
}

func checkOffByOne(s1 string, s2 string) bool {
    // check if strings only differ by one character
    diffCount := 0
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            diffCount += 1
        }
    }

    if diffCount == 1 {
        return true
    } else {
        return false
    }
}

func compareRows(lines []string, idx int, smudge bool) bool {
    offset := 1
    for i := idx; i < len(lines); i++ {
        if idx-offset < 0 || i > len(lines) {
            return true
        }
        if lines[idx-offset] != lines[i] {
            if checkOffByOne(lines[idx-offset], lines[i]) && smudge {
            } else {
                return false
            }
        }
        offset += 1
    }
    return true
}

func transposeSlice(strSlice []string) []string {
    tMap := map[int]string{}
    for _, i := range strSlice {
        for j, s := range i {
            tMap[j] = tMap[j] + string(s)
        }
    }

    keys := []int{}
    for k := range tMap {
        keys = append(keys, k)
    }

    slices.Sort(keys)

    transposed := []string{}
    for _, key := range keys {
        transposed = append(transposed, tMap[key])
    }

    return transposed
}

func printSlice(sl []string) {
    for _, s := range sl {
        fmt.Println(s)
    }
    fmt.Println("")
}

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")

    splits := []int{-1}
    for i, v := range input_lines {
        if v == "" {
            splits = append(splits, i)
        }
    }

    processLines := [][]string{}
    // break up the input into chunks of strings
    for i := 1; i < len(splits); i++ {
        processLines = append(processLines, input_lines[splits[i-1]+1:splits[i]])
    }

    result := 0
    for _, p := range processLines {
        // check for horizontal matches
        checkResult := checkIndexes(p)
        if checkResult != -1 {
            result += (checkResult * 100)
        } else {
            // transpose p columns into rows
            transposed := transposeSlice(p)
            checkResult = checkIndexes(transposed)
            if checkResult != -1 {
                result += checkResult
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
