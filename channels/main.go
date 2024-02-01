package main

import (
    "fmt"
)

func input() chan int {
    ch := make(chan int)
    go func() {
        for i := 2; ; i++ {
            ch <- i
        }
    }()
    return ch
}

func primeNum(in <-chan int, prime int) chan int {
    out := make(chan int)
    go func() {
        for {
            if i := <-in; i%prime != 0 {
                out <- i
            }
        }
    }()
    return out
}

func main() {
    ch := input()
    fmt.Println()
    for i := 0; i < 10; i++ {
        prime := <-ch
        fmt.Printf("%v\n", prime)
        ch = primeNum(ch, prime)
    }
}