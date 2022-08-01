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
		Threads      int64  `default:"16"`
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

    log.Println("Identical clustering")
    withoutEmpty, emptyFiles, clustersSize, clustersNumber := analysis.IdenticalAnalysis(bytecodes, true)
    if emptyFiles > 0 {
        log.Warn("Total empty files detected:", emptyFiles)
    }
    totalAddresses = totalAddresses - emptyFiles
    analysis.PrintResults(totalAddresses, clustersSize, clustersNumber)
    fmt.Println()

    log.Println("Disassembled without arguments clustering")
    processedOpcodes, opcodesClustersSize, opcodesClustersNumber := analysis.DisassembledWithoutArgumentsAnalysis(withoutEmpty)
    analysis.PrintResults(totalAddresses, opcodesClustersSize, opcodesClustersNumber)
    fmt.Println()

    log.Printf("Jaccard similarity with %d-gram and %d%% threshold\n", 5, 90)
    similarityClustersSize, similarityClustersNumber := analysis.SimilarityAnalysis(processedOpcodes)
    analysis.PrintResults(totalAddresses, similarityClustersSize, similarityClustersNumber)
}
