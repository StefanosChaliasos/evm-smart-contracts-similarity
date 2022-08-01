package analysis

import (
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

func IdenticalAnalysis (bytecodes map[string]string, skipProxy bool) (map[string][]string, map[string]string, int, []int, int, int) {
    clustersNumber := 0
    var clustersSize []int
    emptyFiles := 0
    potentialProxies := 0

    clusters := utils.FindIdentical(bytecodes) 
    finalClusters := make(map[string][]string)

    withoutEmpty := make(map[string]string)

    for bytecode, addresses := range clusters {
        if bytecode == "" {
            log.Debug("Addresses with empty values:", addresses)
            emptyFiles = len(addresses)
            continue
        }
        if len(bytecode) <= 100 {
            log.Warn("Potentially proxy contracts:", addresses)
            if skipProxy {
                potentialProxies += len(addresses)
                continue 
            }
        }
        // filter out empty strings
        finalClusters[bytecode] = addresses
        for _, address := range addresses {
            withoutEmpty[address] = bytecode
        }
        if len(addresses) > 1 {
            clustersNumber += 1
            clustersSize = append(clustersSize, len(addresses))
            log.Debug("Cluster items:", addresses)
        }
    }
    return finalClusters, withoutEmpty, emptyFiles, clustersSize, clustersNumber, potentialProxies
}

func DisassembledWithoutArgumentsAnalysis (bytecodes map[string]string) (map[string][]string, map[string]string, []int, int) {
    // Process Opcodes
    processedOpcodes := make(map[string]string)
    for address, bytecodeString := range bytecodes {
        addressOpcodes := utils.DissasembleWithoutPushValues(address, bytecodeString)
        processedOpcodes[address] = addressOpcodes
    }

    clusters := utils.FindIdentical(processedOpcodes) 
    clustersSize, clustersNumber := analyseClusters(clusters)
    return clusters, processedOpcodes, clustersSize, clustersNumber
}

func SimilarityAnalysis (opcodes map[string]string, ngram int, threshold int) (map[string][]string, []int, int) {
    clusters := make(map[string][]string)

    j := metrics.NewJaccard()
    j.CaseSensitive = false
    j.NgramSize = ngram

    for address, addressOpcodes := range opcodes {
        hasSimilar := false
        for clusterOpcodes := range clusters {
            similarity := strutil.Similarity(addressOpcodes, clusterOpcodes, j) 
            if similarity > float64(threshold) * 0.01 {
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
    return clusters, clustersSize, clustersNumber
}
