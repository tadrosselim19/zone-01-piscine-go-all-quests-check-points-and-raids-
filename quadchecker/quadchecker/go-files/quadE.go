package main

import "fmt"

func main() {
    w, h, ok := validateArgs()
    if !ok {
        return
    }
    drawQuadE(w, h)
}

func drawQuadE(w, h int) {
    for row := 0; row < h; row++ {
        for col := 0; col < w; col++ {
            switch {
            case col == 0 && row == 0, col == w-1 && row == h-1 && w > 1 && h > 1:
                fmt.Print("A")
            case col == 0 && row == h-1, col == w-1 && row == 0:
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
