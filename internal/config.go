package sprint

import (
    "io/ioutil"
    "path/filepath"
    "os/user"
    "sync"
    "bytes"
)

func findConfig(current string, configs *[][]byte, m *sync.Mutex) {
    files, err := ioutil.ReadDir(current)
    if err != nil {
        panic(err)
    }
    localConfigs := [][]byte{}
    for _, f := range files {
        if len(f.Name()) >= 7 && bytes.Equal([]byte(f.Name())[len(f.Name())-7:], []byte(".sprint")) {
            text, err := ioutil.ReadFile(filepath.Join(current, f.Name()))
            if err != nil {
                panic(err)
            }
            localConfigs = append(localConfigs, text)
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
