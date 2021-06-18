package sprint

import (
    "path/filepath"
    "os/exec"
    "os/user"
    "strconv"
    "sort"
)

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
    return output
}

type Runnable func(path string, command []byte)
func Run(commands [][]byte, types []projectType, rw Runnable) {
    sort.Slice(types, func(i, j int) bool {
        return types[i].priority < types[j].priority
    })

    rw(types[0].path, commands[types[0].index])
}
