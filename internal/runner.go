package sprint

import (
    "fmt"
    "path/filepath"
    "os/exec"
    "os/user"
)

type callable func(cmd []byte)

func DetectType(detect []byte) []int {
    output := []int{}
    current, _ := filepath.Abs(".")
    usr, _ := user.Current()
    target := usr.HomeDir
    for {
        for i, v := range detect {
            result, err := exec.Command("python3", string(v)).Output()
            fmt.Println(i, v, result)
            if err != nil {
                panic(err)
            }
        }

        if current == target {
            return output
        }

        current = filepath.Clean(filepath.Join(current, "/.."))
    }
}
