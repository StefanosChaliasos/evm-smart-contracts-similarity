package analysis

import (
    "strings"
    "encoding/hex"

    "github.com/ethereum/go-ethereum/core/asm"
    log "github.com/sirupsen/logrus"
    "github.com/adrg/strutil"
    "github.com/adrg/strutil/metrics"

    "github.com/StefanosChaliasos/evm-smart-contracts-similarity/src/utils"
)

func analyseClusters (clusters map[string][]string) ([]int, int) {
    var clustersSize []int
    clustersNumber := 0
    for _, addresses := range clusters {
        if len(addresses) > 1 {
            clustersNumber += 1
            clustersSize = append(clustersSize, len(addresses))
            log.Debug("Cluster items:", addresses)
        }
    }
    return clustersSize, clustersNumber
}

func IdenticalAnalysis (bytecodes map[string]string, checkProxy bool) (map[string]string, int, []int, int) {
    clustersNumber := 0
    var clustersSize []int
    emptyFiles := 0

    clusters := utils.FindIdentical(bytecodes) 

    withoutEmpty := make(map[string]string)

    for bytecode, addresses := range clusters {
        if bytecode == "" {
            log.Debug("Addresses with empty values:", addresses)
            emptyFiles = len(addresses)
            continue
        }
        // filter out empty strings
        for _, address := range addresses {
            withoutEmpty[address] = bytecode
        }
        if checkProxy && len(bytecode) <= 100 {
            log.Debug("Potentially proxy contracts:", addresses)
        }
        if len(addresses) > 1 {
            clustersNumber += 1
            clustersSize = append(clustersSize, len(addresses))
            log.Debug("Cluster items:", addresses)
        }
    }
    return withoutEmpty, emptyFiles, clustersSize, clustersNumber
}

func DisassembledWithoutArgumentsAnalysis (bytecodes map[string]string) (map[string]string, []int, int) {
    // Process Opcodes
    processedOpcodes := make(map[string]string)
    for address, bytecode := range bytecodes {
        script, scriptErr := hex.DecodeString(bytecode)
        if scriptErr != nil {
            log.Panic(scriptErr)
        }
        dis, disErr := asm.Disassemble(script)
        if disErr != nil {
            log.Panic(disErr)
        }
        processed := strings.Join(utils.RemovePushValues(dis), " ")
        processedOpcodes[address] = processed
    }

    clusters := utils.FindIdentical(processedOpcodes) 
    clustersSize, clustersNumber := analyseClusters(clusters)
    return processedOpcodes, clustersSize, clustersNumber
}

func SimilarityAnalysis (opcodes map[string]string) ([]int, int) {
    clusters := make(map[string][]string)

    j := metrics.NewJaccard()
    j.CaseSensitive = false
    j.NgramSize = 5

    for address, addressOpcodes := range opcodes {
        hasSimilar := false
        for clusterOpcodes, _ := range clusters {
            similarity := strutil.Similarity(addressOpcodes, clusterOpcodes, j) 
            if similarity > 0.90 {
                clusters[clusterOpcodes] = append(clusters[clusterOpcodes], address)
                hasSimilar = true
                break
            }
        }
        if !hasSimilar {
            clusters[addressOpcodes] = append(clusters[addressOpcodes], address)
        }
    }

    clustersSize, clustersNumber := analyseClusters(clusters)
    return clustersSize, clustersNumber
}
