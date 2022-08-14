package main

import ( 
    "fmt"
    "sort"
)

func main() {
    fmt.Println("Binary Search from Hardvard CS50. Written in Go.")
    a := []int{4, 2, 0, 1, 3}
    sort.Ints(a)
    bsearch(a, 3)
}

func bsearch(a []int, key int) {
    
    midIdx := ((len(a))/2)
    midElm := a[midIdx]

    if (key == midElm) {
        fmt.Println("found key at mid!")
        return;

    } else if (key <= midElm) {
        fmt.Println("search to the left of mid")
        b := a[0:midIdx]
        bsearch(b, key)
    } else if (key > midElm) {
        b := a[midIdx+1 : len(a)]
        fmt.Println("search to the right of mid!")
        bsearch(b, key)
    }
}

