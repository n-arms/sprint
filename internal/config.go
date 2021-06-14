package sprint

import (
    "io/ioutil"
    "path/filepath"
    "os/user"
)
func findConfig(current string, out chan<- []byte) {
    file, err := ioutil.ReadFile(filepath.Join(current, "/.sprint"))
    if err == nil {
        out <- file
    }
    out <- []byte{}
}

func FindConfigs() [][]byte {
    output := [][]byte{}
    current, _ := filepath.Abs(".")
    usr, _ := user.Current()
    target := usr.HomeDir
    result := make(chan []byte)
    total := 0
    for {
        go findConfig(current, result)
        total += 1

        if current == target {
            break
        }

        current = filepath.Clean(filepath.Join(current, "/.."))
    }

    for i := 0; i < total; i ++ {
        output = append(output, <-result)
    }
    return output
}

func SplitConfig(file []byte) (detect []byte, run []byte) {
    detect = append(file, []byte("\nprint(detect())")...)
    run = append(file, []byte("\nprint(run())")...)
    return
}
