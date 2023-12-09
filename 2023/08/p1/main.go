package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "strings"
    "regexp"
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

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")

    ins := input_lines[0]

    g := &Graph{Nodes: map[string]*Node{}}

    for i := 2; i < len(input_lines); i++ {
        if input_lines[i] != "" {
            re := regexp.MustCompile(`[A-Z]{3}`)
            matches := re.FindAllString(input_lines[i], -1)

            g.AddLeftRight(matches[0], matches[1], matches[2])
        }
    }

    key := "AAA"
    idx := 0
    steps := 0
    for ok := true; ok; ok = (key != "ZZZ") {
        if idx == len(ins) {
            idx = 0
        }
        switch string(ins[idx]) {
        case "L":
            key = g.Nodes[key].Left.Name
        case "R":
            key = g.Nodes[key].Right.Name
        }

        idx += 1
        steps += 1
        if key == "ZZZ" {
            break
        }
    }

    return steps
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Println(res)
}
