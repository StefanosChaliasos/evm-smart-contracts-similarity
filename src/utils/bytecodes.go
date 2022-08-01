package utils

import (
    "strings"

    log "github.com/sirupsen/logrus"
)

func RemovePushValues(dissasembled []string) []string {
    var res []string
    for _, line := range dissasembled {
        line = strings.TrimSpace(line)
        fields := strings.Fields(line)
        if len(fields) > 3 {
            log.Panic("Fields should not be > 3")    
        }
        if len(fields) == 3 {
            if ! strings.HasPrefix(fields[1], "PUSH") {
                log.Panic("OPCODE should be PUSH when len is 3")    
            } 
        }
        res = append(res, fields[1])
    }
    return res
}
