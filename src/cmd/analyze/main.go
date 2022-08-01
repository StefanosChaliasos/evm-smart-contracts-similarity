package main

import (
    "fmt"

    //"github.com/ethereum/go-ethereum/core/asm"
	arg "github.com/alexflint/go-arg"
    log "github.com/sirupsen/logrus"

    "github.com/StefanosChaliasos/evm-smart-contracts-similarity/src/utils"
    "github.com/StefanosChaliasos/evm-smart-contracts-similarity/src/analysis"
)

func main() {
	var args struct {
        Path         string `arg:"positional,required" help:"Path to directory containing bytecodes"` 
        Debug        bool   `help:"Print debug information"` 
        SkipProxy    bool   `help:"Skip potential proxy contracts"` 
        Ngram        int    `help:"Select how many elements should an n-gram have" default:"5"` 
        Threshold    int    `help:"Set a similarity threshold" default:"90"` 
	}

	arg.MustParse(&args)

    log.SetFormatter(&log.TextFormatter{
        TimestampFormat : "2006-01-02 15:04:05",
        FullTimestamp: true,
    })

    if args.Debug {
        log.SetLevel(log.DebugLevel)
    }

    log.Println("Read files from:", args.Path)
    bytecodeFilePaths := utils.WalkDirectoryForFiles(args.Path, ".bin")
    totalAddresses := len(bytecodeFilePaths)
    log.Println("Total files detected:", totalAddresses)
    bytecodes := utils.ReadFiles(bytecodeFilePaths)

    log.Println("Pre-processing")
    withoutEmpty, emptyFiles, clustersSize, clustersNumber, potentialProxies := analysis.IdenticalAnalysis(bytecodes, args.SkipProxy)
    if emptyFiles > 0 {
        log.Warn("Total empty files detected: ", emptyFiles)
    }
    if args.SkipProxy && potentialProxies > 0 {
        log.Warn("Total potential proxy files detected: ", potentialProxies)
    }
    totalAddresses = totalAddresses - emptyFiles -potentialProxies
    log.Println("Identical clustering")
    analysis.PrintResults(totalAddresses, clustersSize, clustersNumber)
    fmt.Println()

    log.Println("Disassembled without arguments clustering")
    processedOpcodes, opcodesClustersSize, opcodesClustersNumber := analysis.DisassembledWithoutArgumentsAnalysis(withoutEmpty)
    analysis.PrintResults(totalAddresses, opcodesClustersSize, opcodesClustersNumber)
    fmt.Println()

    log.Printf("Jaccard similarity with %d-gram and %d%% threshold\n", args.Ngram, args.Threshold)
    similarityClustersSize, similarityClustersNumber := analysis.SimilarityAnalysis(processedOpcodes, args.Ngram, args.Threshold)
    analysis.PrintResults(totalAddresses, similarityClustersSize, similarityClustersNumber)

    log.Println("Finish")
}
