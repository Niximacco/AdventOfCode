package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
)

func main() {
    f, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Print(err)
    }

    input := string(f)
    inputNum, err := strconv.Atoi(input)
    if err != nil {
        panic(err)
    }

    directions := [4]string{"U", "L", "D", "R"}
    currentDirection := 0

    inputSquare := 0
    fmt.Println(inputNum)
    for square := 1; inputNum > (square*square); square+=2 {
        inputSquare = square
    }

    horizontalDistance := (inputSquare+1)/2
    fmt.Printf("Input Square: %d\n", inputSquare)
    verticalDistance := -((inputSquare-1)/2)

    numMoves := inputSquare

    for current:= (inputSquare*inputSquare)+1; current < inputNum; current++ {
        //fmt.Printf("%d: v: %d\th:%d\tnumMoves:%d\tDirection: %s\n", current, verticalDistance, horizontalDistance, numMoves, directions[currentDirection])

        if numMoves == 0 {
            currentDirection++
            if currentDirection > 3 {
                currentDirection = 0
                inputSquare+=2
            }
            numMoves = inputSquare+1
        }

        switch directions[currentDirection] {
        case "U":
            verticalDistance++
        case "L":
            horizontalDistance--
        case "D":
            verticalDistance--
        case "R":
            horizontalDistance++
        }

        numMoves--
    }

    if horizontalDistance < 0 {
        horizontalDistance = -horizontalDistance
    }

    if verticalDistance < 0 {
        verticalDistance= -verticalDistance
    }

    manhattanDistance :=  horizontalDistance + verticalDistance
    fmt.Printf("Horizontal: %d\nVertical: %d\nManhattan: %d\n", horizontalDistance, verticalDistance, manhattanDistance)

}