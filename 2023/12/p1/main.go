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

func strIntSlice(strSlice []string) []int {
    intSlice := make([]int, len(strSlice))
    for i, s := range strSlice {
        intSlice[i], _ = strconv.Atoi(s)
    }

    return intSlice
}

func genRegexpMatch(str string) []int {
    s := strings.Split(str, ",")
    output := []int{}
    for i := range s {
        output = append(output, toInt(s[i]))
    }

    return output
}

func matchString(testStr string, criteria []int) bool {
    re := regexp.MustCompile("[#]+")
    matches := re.FindAllStringIndex(testStr, -1)

    for m := range criteria {
        if m >= len(matches) || m >= len(criteria) {
            return false
        }
        if (matches[m][1] - matches[m][0]) != criteria[m] {
            return false
        }
    }

    if len(matches) > len(criteria) {
        return false
    }

    return true
}

func binaryOfLen(i int64, l int) []int {
	s := fmt.Sprintf("%b", i)


	if len(s) < l {
		pre := l - len(s)
		for j := 0; j < pre; j++ {
			s = "0" + s
		}
	}

	binarySlice := []string{}
	for b := 0; b < len(s); b++ {
		binarySlice = append(binarySlice, string(s[b]))
	}

	return strIntSlice(binarySlice)
}

func replaceStringChars(str string, chars []int, changeMap []int) string {
    mp := map[int]string{0:".", 1:"#"}
    rStr := str

    for i := range chars {
        idx := chars[i]
        rStr = rStr[:idx] + mp[changeMap[i]] + rStr[idx+1:]
    }

    return rStr
}

func bruteForce(inputStr string, criteria []int) int {
    re := regexp.MustCompile("[?]")
    matches := re.FindAllStringIndex(inputStr, -1)

    charLocs := []int{}
    for i := range matches {
        charLocs = append(charLocs, matches[i][0])
    }

    maxBinary := ""
    for k := 0; k < len(charLocs); k++ {
        maxBinary = maxBinary + "1"
    }

    maxInt, _ := strconv.ParseInt(maxBinary, 2, 64)

    matchCount := 0
    for j := 0; j <= int(maxInt); j++ {
        bin := binaryOfLen(int64(j), len(charLocs))

        compStr := replaceStringChars(inputStr, charLocs, bin)

        if matchString(compStr, criteria) {
            matchCount += 1
        }
    }

    return matchCount
}

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")

    total := 0
    for _, ln := range input_lines {
        if ln == "" {
            continue
        }
        lnSplit := strings.Split(ln, " ")

        exp := genRegexpMatch(lnSplit[1])

        arrangements := bruteForce(lnSplit[0], exp)

        total += arrangements
    }

    return total
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Println(res)
}
