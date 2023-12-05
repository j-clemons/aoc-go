package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    "regexp"
    "unicode"
)

func toInt(str string) int {
    i, err := strconv.Atoi(str)
    if err != nil {
        log.Fatal(err)
    }

    return i
}

func schematicProc(currentLn string, prevLn string, nextLn string) int {
    re := regexp.MustCompile("[0-9]+")
    match := re.FindAllStringIndex(currentLn, -1)

    valTotal := 0
    for _, indices := range(match) {
        val := toInt(currentLn[indices[0]:indices[1]])

        s := indices[0]
        e := indices[1]

        if s != 0 {
            s -= 1
        }
        if e < len(currentLn) {
            e += 1
        }

        pln := ""
        if prevLn != "" {
            pln = prevLn[s:e]
        }
        nln := ""
        if nextLn != "" {
            nln = nextLn[s:e]
        }

        cln := string(currentLn[s]) + string(currentLn[e-1])

        check := []rune(pln + nln + cln)
        for _, r := range(check) {
            if unicode.IsDigit(r) == false && r != []rune(".")[0] {
                valTotal += val
                break
            }
        }
    }

    return valTotal
}

func promptFunc(input string) int {
    input_lines := strings.Split(input, "\n")

    returnVal := 0
    for i, ln := range(input_lines) {
        if ln != "" {
            if i == 0 {
                returnVal += schematicProc(ln, "", input_lines[i+1])
            } else if i == len(input_lines)-1 {
                returnVal += schematicProc(ln, input_lines[i-1], "")
            } else {
                returnVal += schematicProc(ln, input_lines[i-1], input_lines[i+1])
            }

        }
    }

    return returnVal
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Print(res)
}
