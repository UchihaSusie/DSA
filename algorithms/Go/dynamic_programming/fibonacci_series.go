// algorithms/Go/dynamic_programming/fibonacci_series.go
package main

import (
	"fmt"
)

func main() {
    key := 20

    if key < 0 {
        fmt.Println("Please enter a valid term.")
        return
    }

    d := make(map[int]int)
    d[0] = 0
    d[1] = 1

    fmt.Printf("The fibonacci series up to term %d: ", key)

    if key == 0 {
        fmt.Println(0)
        return
    }
    if key == 1 {
        fmt.Println(0, "\t", 1)
        return
    }

    fmt.Print(0, "\t", 1)

    for i := 2; i <= key; i++ {
        fmt.Print("\t", fibo(i, d))
    }

    fmt.Println()
}

func fibo(n int, d map[int]int) int {
    if val, exists := d[n]; exists {
        return val
    }
    d[n] = fibo(n-1, d) + fibo(n-2, d)
    return d[n]
}