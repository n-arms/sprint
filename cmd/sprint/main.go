package main

import (
    "github.com/n-arms/sprint/internal"
    "fmt"
)

func main() {
    configs := sprint.FindConfigs()
    detectors := [][]byte{}
    runners := [][]byte{}
    for _, config := range configs {
        detect, run := sprint.SplitConfig(config)
        detectors = append(detectors, detect)
        runners = append(runners, run)
    }
    for _, i := range sprint.DetectType(detectors) {
        fmt.Println(i)
    }
}
