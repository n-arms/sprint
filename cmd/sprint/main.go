package main

import (
    "github.com/n-arms/sprint/internal"
    "fmt"
    "os/exec"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ExecPrint(path string, command []byte) {
    file, err := ioutil.TempFile("", "sprint*.py")
    check(err)
    defer os.Remove(file.Name())

    err = ioutil.WriteFile(file.Name(), command, 0644)
    check(err)

    result, err := exec.Command("bash", "-c", "cd "+path+" && python3 "+file.Name()).CombinedOutput()
    check(err)

    fmt.Println(string(result))
}

func main() {
    configs := sprint.FindConfigs()
    detectors := [][]byte{}
    runners := [][]byte{}
    for _, config := range configs {
        detect, run := sprint.SplitConfig(config)
        detectors = append(detectors, detect)
        runners = append(runners, run)
    }
    types := sprint.DetectType(detectors)
    sprint.Run(runners, types, ExecPrint)
}
