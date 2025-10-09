package main

import (
    "fmt"
    "os"
    "strconv"
)

func validateArgs() (int, int, bool) {
    if len(os.Args) != 3 {
        fmt.Print("Not a quad function")
        return 0, 0, false
    }
    w, err1 := strconv.Atoi(os.Args[1])
    h, err2 := strconv.Atoi(os.Args[2])
    if err1 != nil || err2 != nil || w <= 0 || h <= 0 {
        fmt.Print("Not a quad function")
        return 0, 0, false
    }
    return w, h, true
}
