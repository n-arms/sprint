package sprint

import (
    "path/filepath"
    "os/exec"
    "os/user"
    "strconv"
    "sort"
)

type callable func(cmd []byte)

func getType(tests [][]byte, path string) chan projectType {
    out := make(chan projectType)
    go func() {
        for i, v := range tests {
            result, err := exec.Command("bash", "-c", "cd "+path+" && python3 -c '"+string(v)+"'").CombinedOutput()
            if err != nil {
                continue
            }
            priority, err := strconv.Atoi(string(result[:len(result)-1]))
            if err == nil {
                out <- projectType{index: i, path: path, priority: priority}
            }
        }
        close(out)
    }()
    return out
}

type projectType struct {
    index int;
    path string;
    priority int;
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
    sort.Slice(output, func(i, j int) bool {
        return output[i].priority < output[j].priority
    })
    return output
}

