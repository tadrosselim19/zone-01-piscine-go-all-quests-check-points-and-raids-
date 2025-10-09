package main

import "fmt"

func main() {
    w, h, ok := validateArgs()
    if !ok {
        return
    }
    drawQuadC(w, h)
}

func drawQuadC(w, h int) {
    for row := 0; row < h; row++ {
        for col := 0; col < w; col++ {
            switch {
            case row == 0 && (col == 0 || col == w-1):
                fmt.Print("A")
            case row == h-1 && (col == 0 || col == w-1):
                fmt.Print("C")
            case col == 0 || col == w-1 || row == 0 || row == h-1:
                fmt.Print("B")
            default:
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}
