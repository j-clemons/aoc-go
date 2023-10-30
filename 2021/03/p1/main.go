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

type Diagnostic struct {
    Zero int `default:0`
    One  int `default:0`
}

func (diag *Diagnostic) valueFrequency() (int, int) {
    if diag.Zero > diag.One {
        return 0, 1
    } else {
        return 1, 0
    }
}

func (diag *Diagnostic) parseValue(val byte) {
    if val == '0' {
        diag.Zero += 1
    } else if val == '1' {
        diag.One += 1
    }
}

func powerConsumption(input string) int64 {
    input_lines := strings.Split(input, "\n")

    binaryLen := len(input_lines[0])

    diagList:= make([]Diagnostic, binaryLen)

    for i := range input_lines {
        if input_lines[i] != "" {

            for j := 0; j < binaryLen; j++ {
                diagList[j].parseValue(input_lines[i][j])
            }
        }
    }

    gamma_str := ""
    epsilon_str := ""

    for i := range diagList {
        g, e := diagList[i].valueFrequency()
        gamma_str += strconv.Itoa(g)
        epsilon_str += strconv.Itoa(e)
    }

    gamma, err := strconv.ParseInt(gamma_str, 2, 64)
    if err != nil {
        log.Fatal(err)
    }
    epsilon, err := strconv.ParseInt(epsilon_str, 2, 64)
    if err != nil {
        log.Fatal(err)
    }

    return gamma * epsilon
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := powerConsumption(string(file))

    fmt.Print(res)
}
