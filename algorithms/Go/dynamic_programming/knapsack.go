// author: @UchihaSusie	
// algorithms/Go/dynamic_programming/knapsack.go
package main

import (
    "fmt"
)

func knapsack(items [][]int, capacity int) int {
    dp := make([][]int, len(items)+1)
    for i := range dp {
        dp[i] = make([]int, capacity+1)
    }

    for row := 1; row < len(dp); row++ {
        for col := 1; col < len(dp[row]); col++ {
            currentWeight := items[row-1][1]
            currentValue := items[row-1][0]

            if currentWeight > col {
                dp[row][col] = dp[row-1][col]
            } else {
                dp[row][col] = max(dp[row-1][col], dp[row-1][col-currentWeight]+currentValue)
            }
        }
    }
    return dp[len(items)][capacity]
}


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    items := [][]int{{1, 2}, {4, 3}, {5, 6}, {6, 7}}
    capacity := 10

    knapsackValue := knapsack(items, capacity)

    fmt.Println(knapsackValue) // expected output: 10
}