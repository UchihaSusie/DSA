// algorithms/Go/dynamic_programming/fibonacci_series_sum.go
package main

import (
    "fmt"
)

func main() {
    key := 10

    if key < 0 {
        fmt.Println("Please enter a valid term.")
        return
    }

    d := make(map[int]int)
    d[0] = 0
    d[1] = 1

    if key == 0 {
        fmt.Printf("Sum up to term %d of fibonacci series: %d\n", key, 0)
        return
    }

    if key == 1 {
        fmt.Printf("Sum up to term %d of fibonacci series: %d\n", key, 1)
        return
    }

    sum := 1

    for i := 2; i <= key; i++ {
        sum += fibo(i, d)
    }

    fmt.Printf("Sum up to term %d of fibonacci series is: %d\n", key, sum)
}

func fibo(n int, d map[int]int) int {
    if val, exists := d[n]; exists {
        return val
    }
    d[n] = fibo(n-1, d) + fibo(n-2, d)
    return d[n]
}