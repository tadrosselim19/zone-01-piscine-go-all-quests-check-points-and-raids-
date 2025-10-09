package main

import "fmt"

func main() {
    w, h, ok := validateArgs()
    if !ok {
        return
    }
    drawQuadA(w, h)
}

func drawQuadA(w, h int) {
    for row := 0; row < h; row++ {
        for col := 0; col < w; col++ {
            switch {
            case (col == 0 || col == w-1) && (row == 0 || row == h-1):
                fmt.Print("o")
            case col == 0 || col == w-1:
                fmt.Print("|")
            case row == 0 || row == h-1:
                fmt.Print("-")
            default:
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}
