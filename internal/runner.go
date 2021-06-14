package sprint

import (
    "path/filepath"
    "os/exec"
    "os/user"
    "bytes"
    "fmt"
)

type callable func(cmd []byte)

func getType(tests [][]byte, path string) chan int {
    out := make(chan int)
    go func() {
        for i, v := range tests {
            result, err := exec.Command("bash", "-c", "cd "+path+" && python3 -c '"+string(v)+"'").Output()
            if err == nil && bytes.Equal(result, []byte("True\n")) {
                out <- i
            }
        }
        close(out)
    }()
    return out
}

func DetectType(tests [][]byte) []int {
    output := []int{}
    current, _ := filepath.Abs(".")
    usr, _ := user.Current()
    target := usr.HomeDir
    chans := []chan int{}
    for {
        chans = append(chans, getType(tests, current))
        if current == target {
            break
        }

        current = filepath.Clean(filepath.Join(current, "/.."))
    }

    for _, c := range chans {
        for i := range c {
            output = append(output, i)
        }
    }

    return output
}
