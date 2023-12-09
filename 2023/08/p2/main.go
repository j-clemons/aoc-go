package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "strings"
    "regexp"
    "slices"
    "math/big"

    "github.com/fxtlabs/primes"
)

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

type Graph struct {
    Nodes map[string]*Node
}

type Node struct {
    Name  string
    Left  *Node
    Right *Node
}

func (g *Graph) AddNode(key string) {
    if _, ok := g.Nodes[key]; !ok {
        g.Nodes[key] = &Node{
            Name: key,
            Left: &Node{},
            Right: &Node{},
        }
    }
}

func (g *Graph) AddLeftRight(srcKey string, leftKey string, rightKey string) {
    if _, ok := g.Nodes[srcKey]; !ok {
        g.AddNode(srcKey)
    }
    if _, ok := g.Nodes[leftKey]; !ok {
        g.AddNode(leftKey)
    }
    if _, ok := g.Nodes[rightKey]; !ok {
        g.AddNode(rightKey)
    }

    g.Nodes[srcKey].Left = g.Nodes[leftKey]
    g.Nodes[srcKey].Right = g.Nodes[rightKey]
}

func (g *Graph) getKeysEnding(str string) []string {
    var matchingSlice []string
    for _, v := range g.Nodes {
        if string(v.Name[len(v.Name)-1]) == str {
            matchingSlice = append(matchingSlice, v.Name)
        }
    }

    return matchingSlice
}

func primeFactors(i int) (int, int) {
    if big.NewInt(int64(i)).ProbablyPrime(0) {
        return 1, i
    }

    primeSlice := primes.Sieve(i)
    for _, v := range primeSlice {
        if i % v == 0 {
            r := i / v

            if big.NewInt(int64(r)).ProbablyPrime(0) {
                return v, r
            } else {
                return primeFactors(r)
            }
        }
    }
    return 1, i
}

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")

    ins := input_lines[0]

    g := &Graph{Nodes: map[string]*Node{}}

    for i := 2; i < len(input_lines); i++ {
        if input_lines[i] != "" {
            re := regexp.MustCompile(`[A-Z0-9]{3}`)
            matches := re.FindAllString(input_lines[i], -1)

            g.AddLeftRight(matches[0], matches[1], matches[2])
        }
    }

    keys := g.getKeysEnding("A")

    primeFactorSlice := []int{}
    for _, v := range keys {
        idx := 0
        steps := 0
        for 1 > 0 {
            if idx == len(ins) {
                idx = 0
            }
            switch string(ins[idx]) {
            case "L":
                v = g.Nodes[v].Left.Name
            case "R":
                v = g.Nodes[v].Right.Name
            }

            idx += 1
            steps += 1

            if string(v[2]) == "Z" {
                p1, p2 := primeFactors(steps)
                if !slices.Contains(primeFactorSlice, p1) {
                    primeFactorSlice = append(primeFactorSlice, p1)
                }
                if !slices.Contains(primeFactorSlice, p2) {
                    primeFactorSlice = append(primeFactorSlice, p2)
                }
                break
            }
        }
    }

    output := 1
    for _, s := range primeFactorSlice {
        output *= s
    }

    return output
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Println(res)
}
