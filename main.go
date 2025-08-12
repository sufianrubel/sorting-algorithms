package main

import (
    "fmt"
    "sorting-algorithms/algorithms"
)

func main() {
    arr := []int{42, 25, 65, 15, 100}

    fmt.Print("Original array: ")
    for _, v := range arr {
        fmt.Printf("%d ", v)
    }
    fmt.Println()

    algorithms.BubbleSort(arr)

    fmt.Print("Sorted array (ascending order): ")
    for _, v := range arr {
        fmt.Printf("%d ", v)
    }
    fmt.Println()
}
