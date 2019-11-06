package main

import (
    "fmt"
    "os"
    "strconv"
    "time"
)

// recursive if-else
func f1(n int) int {
    if n == 0 {
        return 0
    }

    if n == 1 {
        return 1
    }

    return f1(n-1) + f1(n-2)
}

// optimization
func f2(n int) int {
    a, b := 0, 1

    for i := 2; i < n + 1; i++ {
        a, b = b, a + b
    }

    return b
}

func measure(f func(arg1 int) int, n int) int {
    start := time.Now()
    result := f(n)
    fmt.Printf("%.8f\n", time.Since(start).Seconds())

    return result
}

func start(f string, n int) int {
    switch f {
        case "f1": return measure(f1, n)
        case "f2": return measure(f2, n)
        default: return 0
    }
}

func main() {
    f := os.Args[1]
    n, _ := strconv.Atoi(os.Args[2])

    fmt.Println(start(f, n))
}
