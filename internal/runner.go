package sprint

import (
    "path/filepath"
    "os/exec"
    "os/user"
    "bytes"
)

type callable func(cmd []byte)

func getType(tests [][]byte, path string) chan projectType {
    out := make(chan projectType)
    go func() {
        for i, v := range tests {
            result, err := exec.Command("bash", "-c", "cd "+path+" && python3 -c '"+string(v)+"'").CombinedOutput()
            if err == nil && bytes.Equal(result, []byte("True\n")) {
                out <- projectType{index: i, path: path}
            }
        }
        close(out)
    }()
    return out
}

type projectType struct {
    index int;
    path string;
}

func DetectType(tests [][]byte) []projectType {
    output := []projectType{}
    current, _ := filepath.Abs(".")
    usr, _ := user.Current()
    target := usr.HomeDir
    chans := []chan projectType{}
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

