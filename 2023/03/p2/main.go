package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    "regexp"
)

func toInt(str string) int {
    i, err := strconv.Atoi(str)
    if err != nil {
        log.Fatal(err)
    }

    return i
}

func checkLine(str string, start int) []int {
    r := regexp.MustCompile(`\*`)
    match := r.FindAllStringIndex(str, -1)

    var locations []int
    for _, m := range(match) {
       locations = append(locations, (m[0]+start)) 
    }

    return locations
}

func schematicProc(
    currentLn string,
    idx int, 
    prevLn string, 
    nextLn string, 
    locMap map[string][]int) map[string][]int {

    locMapCopy := locMap
    re := regexp.MustCompile("[0-9]+")
    match := re.FindAllStringIndex(currentLn, -1)

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

        clns := string(currentLn[s])
        clne := string(currentLn[e-1])

        plnMatch := checkLine(pln, s)
        for _, p := range(plnMatch) {
            loc := fmt.Sprintf("%d,%d", idx-1, p)
            lmcVals := locMapCopy[loc]
            locMapCopy[loc] = append(lmcVals, val)
        }
        nlnMatch := checkLine(nln, s)
        for _, p := range(nlnMatch) {
            loc := fmt.Sprintf("%d,%d", idx+1, p)
            lmcVals := locMapCopy[loc]
            locMapCopy[loc] = append(lmcVals, val)
        }
        if clns == "*" {
            loc := fmt.Sprintf("%d,%d", idx, s)
            lmcVals := locMapCopy[loc]
            locMapCopy[loc] = append(lmcVals, val)
        }
        if clne == "*" {
            loc := fmt.Sprintf("%d,%d", idx, e-1)
            lmcVals := locMapCopy[loc]
            locMapCopy[loc] = append(lmcVals, val)
        }
    }

    return locMapCopy
}

func calcGearRatioTotal(mp map[string][]int) int {
    grTotal := 0
    for _, v := range mp {
        if len(v) == 2 {
            gr := v[0] * v[1]
            grTotal += gr
        }
    }

    return grTotal
}

func promptFunc(input string) int {
    input_lines := strings.Split(input, "\n")

    outMap := make(map[string][]int)
    for i, ln := range(input_lines) {
        if ln != "" {
            if i == 0 {
                outMap = schematicProc(ln, i, "", input_lines[i+1], outMap)
            } else if i == len(input_lines)-1 {
                outMap = schematicProc(ln, i, input_lines[i-1], "", outMap)
            } else {
                outMap = schematicProc(ln, i, input_lines[i-1], input_lines[i+1], outMap)
            }
        }
    }

    return calcGearRatioTotal(outMap)
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Print(res)
}
