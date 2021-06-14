package sprint

import (
    "io/ioutil"
    "path/filepath"
    "os/user"
    "bytes"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func FindConfigs() []byte {
    output := []byte{}
    current, err := filepath.Abs(".")
    check(err)
    usr, _ := user.Current()
    target := usr.HomeDir
    for {
        file, err := ioutil.ReadFile(filepath.Join(current, "/.sprint"))
        if err == nil {
            file = bytes.TrimRight(file, "\n\t ")
            output = append(output, append(file, byte('\n'))...)
        }

        if current == target {
            return output;
        }

        current = filepath.Clean(filepath.Join(current, "/.."))
    }
    return output;
}
