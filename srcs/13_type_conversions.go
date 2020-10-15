package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y int        = 3, 4;
    var f float64       = math.Sqrt(float64(x * x + y * y));
    var z uint          = uint(f);
    i                   := 42;
    f2                  := float64(i);
    u                   := uint(f);

    fmt.Println(x, y, z);
    fmt.Println(i, f2, u);
}
