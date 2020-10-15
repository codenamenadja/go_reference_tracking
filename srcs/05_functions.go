package main

import (
    "fmt"
)

func main() {
    fmt.Println(add(42, 13));
}

func add(x int, y int) int {
    return (x + y);
} // no forward declareations in go.

