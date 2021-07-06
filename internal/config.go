package sprint

import (
    "io/ioutil"
    "path/filepath"
    "os/user"
    "sync"
    "bytes"
    "log"
    "os"
    "fmt"
)

func findConfig(current string, configs *[][]byte, m *sync.Mutex) {
    files, err := ioutil.ReadDir(current)
    if err != nil {
        log.Fatal("error", err, "at internal/config.go:15")
        return
    }
    localConfigs := [][]byte{}
    for _, f := range files {
        if len(f.Name()) >= 7 && bytes.Equal([]byte(f.Name())[len(f.Name())-7:], []byte(".sprint")) {
            text, err := ioutil.ReadFile(filepath.Join(current, f.Name()))
            if err == nil {
                localConfigs = append(localConfigs, text)
            }
        }
    }
    m.Lock()
    *configs = append(*configs, localConfigs...)
    m.Unlock()
}

func FindConfigs() [][]byte {
    configs := [][]byte{}
    current, _ := filepath.Abs(".")
    usr, _ := user.Current()
    var wg sync.WaitGroup
    var m sync.Mutex
    config, err := os.UserConfigDir()
    if err != nil {
        log.Fatal("failed to find homedir with error", err)
    }
    wg.Add(1)
    go func(){
        fmt.Println("searching for configs at path", filepath.Join(config, "sprint"))
        findConfig(filepath.Join(config, "sprint"), &configs, &m)
        wg.Done()
    }()

    for {
        wg.Add(1)
        go func(current string){
            findConfig(current, &configs, &m)
            wg.Done()
        }(current)
        
        if current == usr.HomeDir {
            wg.Wait()
            return configs
        }

        current = filepath.Clean(filepath.Join(current, "/.."))
    }


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
