package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
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

type Hand struct {
    Cards    string
    Bid      int
    Strength int
    C0       int
    C1       int
    C2       int
    C3       int
    C4       int
}

func (h *Hand) rankCards() {
    cardRank := map[string]int {
        "A":14,
        "K":13,
        "Q":12,
        "T":10,
        "9":9,
        "8":8,
        "7":7,
        "6":6,
        "5":5,
        "4":4,
        "3":3,
        "2":2,
        "J":1,
    }

    h.C0 = cardRank[string(h.Cards[0])]
    h.C1 = cardRank[string(h.Cards[1])]
    h.C2 = cardRank[string(h.Cards[2])]
    h.C3 = cardRank[string(h.Cards[3])]
    h.C4 = cardRank[string(h.Cards[4])]
}

func (h *Hand) calcStrength() {
    handMap := map[string]int{}

    for i := 0; i < len(h.Cards); i++ {
        handMap[string(h.Cards[i])] += 1
    }

    hasJ := false
    maxKey := ""
    for k, v := range handMap{
       if k == "J" {
           hasJ = true
       } else if maxKey == "" {
           maxKey = k
       } else if v > handMap[maxKey] {
           maxKey = k
       }
    }

    if hasJ {
        handMap[maxKey] += handMap["J"]
        delete(handMap, "J")
    }

    mapLen := len(handMap)
    if mapLen == 1 {
        h.Strength = 6
    } else if mapLen == 2 {
        vals := []int{}
        for _, v := range handMap {
            vals = append(vals, v)
        }
        if slices.Contains(vals, 4) {
            h.Strength = 5
        } else {
            h.Strength = 4
        }
    } else if mapLen == 3 {
        vals := []int{}
        for _, v := range handMap {
            vals = append(vals, v)
        }
        if slices.Contains(vals, 3) {
            h.Strength = 3
        } else {
            h.Strength = 2
        }
    } else if mapLen == 4 {
        h.Strength = 1
    } else {
        h.Strength = 0
    }
}

func sortHands(h []Hand) {
    sort.SliceStable(h,  func(i, j int) bool {
        hi, hj := h[i], h[j]
        switch {
        case hi.Strength != hj.Strength:
            return hi.Strength < hj.Strength
        case hi.C0 != hj.C0:
            return hi.C0 < hj.C0
        case hi.C1 != hj.C1:
            return hi.C1 < hj.C1
        case hi.C2 != hj.C2:
            return hi.C2 < hj.C2
        case hi.C3 != hj.C3:
            return hi.C3 < hj.C3
        case hi.C4 != hj.C4:
            return hi.C4 < hj.C4
        default:
            return hi.Strength < hj.Strength
        }
    })
}

func promptFunc(input string) int {
    defer duration(track("promptFunc"))
    input_lines := strings.Split(input, "\n")

    handSlice := []Hand{}
    for _, ln := range(input_lines) {
        if ln != "" {
            raw := strings.Fields(ln)
            hand := Hand{
                Cards: raw[0],
                Bid:   toInt(raw[1]),
            }
            hand.rankCards()
            hand.calcStrength()
            handSlice = append(handSlice, hand)
        }
    }

    sortHands(handSlice)

    totalWinnings := 0
    for i, h := range(handSlice) {
        totalWinnings += ((i+1) * h.Bid)
    }

    return totalWinnings
}

func main() {
    file, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    res := promptFunc(string(file))

    fmt.Println(res)
}
