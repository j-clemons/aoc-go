package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

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

func createMap(returnMap *map[int]int, source int, dest int, rng int) {
    for i := 0; i < rng; i++ {
        (*returnMap)[source] = dest
    }
}

func getMapVal(val int, mapInput [][]int) int {
    for _, v := range(mapInput) {
        if val >= v[1] && val < v[1]+v[2] {
            return v[0] + (val - v[1])
        }
    }

    return val
}

func promptFunc(input string) int {
    input_lines := strings.Split(input, "\n")

    seeds := []int{}
    mapToBuild := ""

    seedSoilMapInputs := [][]int{}
    soilFertMapInputs := [][]int{}
    fertWatrMapInputs := [][]int{}
    watrLghtMapInputs := [][]int{}
    lghtTempMapInputs := [][]int{}
    tempHumdMapInputs := [][]int{}
    humdLocaMapInputs := [][]int{}
    for _, ln := range(input_lines) {
        if ln == "" {
            continue
        }
        if strings.Contains(ln, "seeds") {
            seeds = strIntSlice(strings.Fields(strings.Split(ln, ":")[1]))
            continue
        }
        if strings.Contains(ln, "map") {
            mapToBuild = strings.Fields(ln)[0]
            continue
        }

        mapInput := strIntSlice(strings.Fields(ln))
        switch mapToBuild {
        case "seed-to-soil":
            seedSoilMapInputs = append(seedSoilMapInputs, mapInput)
        case "soil-to-fertilizer":
            soilFertMapInputs = append(soilFertMapInputs, mapInput)
        case "fertilizer-to-water":
            fertWatrMapInputs = append(fertWatrMapInputs, mapInput)
        case "water-to-light":
            watrLghtMapInputs = append(watrLghtMapInputs, mapInput)
        case "light-to-temperature":
            lghtTempMapInputs = append(lghtTempMapInputs, mapInput)
        case "temperature-to-humidity":
            tempHumdMapInputs = append(tempHumdMapInputs, mapInput)
        case "humidity-to-location":
            humdLocaMapInputs = append(humdLocaMapInputs, mapInput)
        }
    }

    finalLocation := -1

    for i, v := range(seeds) {
        if i % 2 == 0 {
            for j := 0; j < seeds[i+1]; j++ {
                nextVal := v+j
                nextVal = getMapVal(nextVal, seedSoilMapInputs)
                nextVal = getMapVal(nextVal, soilFertMapInputs)
                nextVal = getMapVal(nextVal, fertWatrMapInputs)
                nextVal = getMapVal(nextVal, watrLghtMapInputs)
                nextVal = getMapVal(nextVal, lghtTempMapInputs)
                nextVal = getMapVal(nextVal, tempHumdMapInputs)
                nextVal = getMapVal(nextVal, humdLocaMapInputs)

                if finalLocation == -1 {
                    finalLocation = nextVal
                } else if nextVal < finalLocation {
                    finalLocation = nextVal
                }
            }
        }
    }

    return finalLocation
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Print(res)
}
