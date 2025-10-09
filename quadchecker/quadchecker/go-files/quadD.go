package main

import "fmt"

func main() {
    w, h, ok := validateArgs()
    if !ok {
        return
    }
    drawQuadD(w, h)
}

func drawQuadD(w, h int) {
    for row := 0; row < h; row++ {
        for col := 0; col < w; col++ {
            switch {
            case col == 0 && (row == 0 || row == h-1):
                fmt.Print("A")
            case col == w-1 && (row == 0 || row == h-1):
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
