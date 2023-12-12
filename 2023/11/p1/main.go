package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "strings"
    "strconv"
    "regexp"
    "slices"
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

func findEmptyColumns(col map[int][]int, availCols []int) []int {
    galaxyCols := []int{}
    for _, v := range col {
        for i := range v {
            galaxyCols = append(galaxyCols, v[i])
        }
    }

    emptyCols := []int{}
    for _, c := range availCols {
        if !slices.Contains(galaxyCols, c) {
            emptyCols = append(emptyCols, c)
        }
    }

    return emptyCols
}

func galaxyList(galMap map[int][]int) [][]int {
    output := [][]int{}
    for k, v := range galMap {
        for i := range v {
            output = append(output, []int{k, v[i]})
        }
    }

    return output
}

func calcDistance(g1 []int, g2 []int) int {
    x1 := float64(g1[0])
    y1 := float64(g1[1])
    x2 := float64(g2[0])
    y2 := float64(g2[1])

    xDiff := math.Abs(x2 - x1)
    yDiff := math.Abs(y2 - y1)

    if xDiff < yDiff {
        return int(yDiff - xDiff + (xDiff * 2))
    } else {
        return int(xDiff - yDiff + (yDiff * 2))
    }
}

func countSliceLessThan(raw []int, maxVal int) int {
    ct := 0
    for i := range raw {
        if raw[i] < maxVal {
            ct += 1
        }
    }

    return ct
}

func offsetGalaxies(galaxies [][]int, rows []int, cols []int) [][]int {
    updatedGalaxies := [][]int{}
    for _, g := range galaxies {
        g[0] += countSliceLessThan(rows, g[0])
        g[1] += countSliceLessThan(cols, g[1])
        updatedGalaxies = append(updatedGalaxies, g)
    }

    return updatedGalaxies
}

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")

    availCols := []int{}
    for i := 0; i < len(input_lines[0]); i++ {
        availCols = append(availCols, i)
    }

    emptyRows := []int{}
    rowGalaxies := map[int][]int{}
    re := regexp.MustCompile("#")
    for i, ln := range input_lines {
        matches := re.FindAllStringIndex(ln, -1)
        if len(matches) == 0 {
            emptyRows = append(emptyRows, i)
        } else {
            for _, m := range matches {
                rowGalaxies[i] = append(rowGalaxies[i], m[0])
            }
        }
    }

    emptyColumns := findEmptyColumns(rowGalaxies, availCols)

    galaxies := offsetGalaxies(galaxyList(rowGalaxies), emptyRows, emptyColumns)
    distTotal := 0
    for i, g := range galaxies {
        for j := i+1; j < len(galaxies); j++ {
            dist := calcDistance(g, galaxies[j])
            distTotal += dist
        }
    }
         
    return distTotal
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Println(res)
}
