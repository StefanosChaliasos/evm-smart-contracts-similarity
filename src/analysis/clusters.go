package analysis

import (
    "log"
    "fmt"
    "strings"
    "encoding/hex"

    "github.com/ethereum/go-ethereum/core/asm"

    "github.com/StefanosChaliasos/evm-smart-contracts-similarity/src/utils"
)

func IdenticalAnalysis (bytecodes map[string]string, checkProxy bool) (map[string]string, int, []int, int) {
    log.Println("Identical clustering")

    clustersNumber := 0
    var clustersSize []int
    emptyFiles := 0

    clusters := utils.FindIdentical(bytecodes) 

    withoutEmpty := make(map[string]string)

    for bytecode, addresses := range clusters {
        if bytecode == "" {
            fmt.Println("Addresses with empty values:", addresses)
            emptyFiles = len(addresses)
            continue
        }
        // filter out empty strings
        for _, address := range addresses {
            withoutEmpty[address] = bytecode
        }
        if checkProxy && len(bytecode) <= 100 {
            fmt.Println("Potentially proxy contracts:", addresses)
        }
        if len(addresses) > 1 {
            clustersNumber += 1
            clustersSize = append(clustersSize, len(addresses))
            fmt.Println("Cluster items:", addresses)
        }
    }
    return withoutEmpty, emptyFiles, clustersSize, clustersNumber
}

func DisassembledWithoutArgumentsAnalysis (bytecodes map[string]string) (map[string]string, int, []int, int) {
    log.Println("Disassembled without arguments clustering")

    // Process Opcodes
    processedOpcodes := make(map[string]string)
    for address, bytecode := range bytecodes {
        script, scriptErr := hex.DecodeString(bytecode)
        if scriptErr != nil {
            log.Println(scriptErr)
        }
        dis, disErr := asm.Disassemble(script)
        if disErr != nil {
            log.Println(disErr)
        }
        processed := strings.Join(utils.RemovePushValues(dis), " ")
        processedOpcodes[address] = processed
    }

    return IdenticalAnalysis(processedOpcodes, false)
}
