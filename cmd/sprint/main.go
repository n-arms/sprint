package main

import (
    "github.com/n-arms/sprint/internal"
    "fmt"
    "os/exec"
    "io/ioutil"
    "os"
    "log"
)

func ExecPrint(path string, command []byte) {
    file, err := ioutil.TempFile("", "sprint*.py")
    if err != nil {
        log.Fatal("failed to create temp file for running final executable, error: ", err)
    }
    defer os.Remove(file.Name())

    err = ioutil.WriteFile(file.Name(), command, 0644)
    if err != nil {
        log.Fatal("failed to write to temp file, error: ", err)
    }

    result, err := exec.Command("bash", "-c", "cd "+path+" && python3 "+file.Name()).CombinedOutput()
    if err != nil {
        log.Fatal("failed to exec python command with error ", err, " and output ", string(result))
    }

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
