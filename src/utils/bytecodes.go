package utils

import (
    "log"
    "strings"
)

func RemovePushValues(dissasembled []string) []string {
    var res []string
    for _, line := range dissasembled {
        line = strings.TrimSpace(line)
        fields := strings.Fields(line)
        if len(fields) > 3 {
            log.Println("Fields should not be > 3")    
        }
        if len(fields) == 3 {
            if ! strings.HasPrefix(fields[1], "PUSH") {
                log.Println("OPCODE should be PUSH when len is 3")    
            } 
        }
        res = append(res, fields[1])
    }
    return res
}
