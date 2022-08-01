package utils

import (
	"path/filepath"
	"os"
    "strings"
    "io/ioutil"

    log "github.com/sirupsen/logrus"
)

func WalkDirectoryForFiles(directory string, extension string) []string {
    var res []string
    err := filepath.Walk(directory,
        func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if strings.HasSuffix(path, extension) {
            res = append(res, path)
        }
        return nil
    })
    if err != nil {
        log.Panic(err)
    }
    return res
}

func ReadFiles(files []string) map[string]string {
    res := make(map[string]string)
    for _, file := range files {
        fileBuf, err := ioutil.ReadFile(file)
        if err != nil {
            log.Panic(err)
        }
        fileContents := string(fileBuf)
        basename := filepath.Base(file)
        basename = strings.TrimSuffix(basename, filepath.Ext(basename))
        res[basename] = fileContents
    }
    return res
}
