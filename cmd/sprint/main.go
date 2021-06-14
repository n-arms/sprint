package main

import (
    "github.com/n-arms/sprint/internal"
    "fmt"
)

func main() {
    fmt.Println(sprint.DetectType([][]byte{[]byte("print(\"True\")")}))
}
