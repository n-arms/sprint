package sprint

import (
    "io/ioutil"
    "path/filepath"
    "os/user"
    //"fmt"
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
        r := <-result
        if len(r) > 0 {
            output = append(output, r)
        }
    }
    return output
}

func SplitConfig(file []byte) (detect []byte, run []byte) {
    detect = make([]byte, len(file))
    run = make([]byte, len(file))
    copy(detect, file)
    copy(run, file)

    run = append(run, []byte("\nprint(run())")...)
    detect = append(detect, []byte("\nprint(detect())")...)

    return
}
