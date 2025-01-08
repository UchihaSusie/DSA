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

    if key < 2 {
        fmt.Printf("The term %d of the fibonacci series is: %d\n", key, d[key])
        return
    }

    fmt.Println(fibo(key, d))
}

func fibo(n int, d map[int]int) int {
    if val, exists := d[n]; exists {
        return val
    }
    d[n] = fibo(n-1, d) + fibo(n-2, d)
    return d[n]
}
