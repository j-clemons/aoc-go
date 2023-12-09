package main

import (
    "testing"
)

func TestPrimeFactors(t *testing.T) {
     i := 40
     e1 := 2
     e2 := 5

     a1, a2 := primeFactors(i)

     if a1 != e1 || a2 != e2 {
         t.Errorf("Expected %d %d, Got %d %d \n", e1, e2, a1, a2)
     }

     i = 17
     e1 = 1
     e2 = 17

     a1, a2 = primeFactors(i)

     if a1 != e1 || a2 != e2 {
         t.Errorf("Expected %d %d, Got %d %d \n", e1, e2, a1, a2)
     }
}

func TestPromptFunc(t *testing.T) {

    i := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

    var expected int = 6

    actual := promptFunc(i)

    if actual != expected {
        t.Errorf("got: %d; want: %d", actual, expected)
    }
}
