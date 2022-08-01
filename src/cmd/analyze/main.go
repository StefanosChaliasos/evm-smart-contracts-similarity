package main

import (
    "fmt"

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

    log.SetFormatter(&utils.MyFormatter{log.TextFormatter{
        TimestampFormat : "2006-01-02 15:04:05",
        FullTimestamp: true,
    }})

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
    analysis.PrintResults(totalAddresses, emptyFiles, clustersSize, clustersNumber)
    fmt.Println()

    log.Println("Disassembled without arguments clustering")
    _, _, opcodesClustersSize, opcodesClustersNumber := analysis.DisassembledWithoutArgumentsAnalysis(withoutEmpty)
    analysis.PrintResults(totalAddresses, emptyFiles, opcodesClustersSize, opcodesClustersNumber)
}
