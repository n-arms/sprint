package sprint

import (
    "path/filepath"
    "os/exec"
    "os/user"
    "strconv"
    "sort"
    "sync"
    "os"
    "io/ioutil"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func pyExec(command []byte, current string) ([]byte, error) {
    file, err := ioutil.TempFile("", "sprint_test*.py")
    check(err)
    defer os.Remove(file.Name())

    err = ioutil.WriteFile(file.Name(), command, 0644)
    check(err)

    return exec.Command("bash", "-c", "cd "+current+" && python3 "+file.Name()).CombinedOutput()
}

func getType(tests [][]byte, path string, types *[]projectType, m *sync.Mutex) {
    for i, v := range tests {
        result, err := pyExec(v, path)
        if err != nil {
            continue
        }
        priority, err := strconv.Atoi(string(result[:len(result)-1]))
        if err == nil {
            m.Lock()
            *types = append(*types, projectType{index: i, path: path, priority: priority})
            m.Unlock()
        }
    }
}

type projectType struct {
    index int;
    path string;
    priority int;
}


func DetectType(tests [][]byte) []projectType {
    types := new([]projectType)
    current, _ := filepath.Abs(".")
    usr, _ := user.Current()
    var wg sync.WaitGroup
    var m sync.Mutex
    for {
        wg.Add(1)
        go func(current string){
            getType(tests, current, types, &m)
            wg.Done()
        }(current)

        if current == usr.HomeDir {
            wg.Wait()
            return *types
        }

        current = filepath.Clean(filepath.Join(current, "/.."))
    }
}

type Runnable func(path string, command []byte)
func Run(commands [][]byte, types []projectType, rw Runnable) {
    sort.Slice(types, func(i, j int) bool {
        return types[i].priority < types[j].priority
    })
    if len(types) == 0 {
        panic("no configs found")
    }
    rw(types[0].path, commands[types[0].index])
}
