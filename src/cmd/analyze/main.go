package main

import (
    "log"
    "fmt"
    _ "path/filepath"
	_ "os"

	arg "github.com/alexflint/go-arg"

    "github.com/StefanosChaliasos/evm-smart-contracts-similarity/src/utils"
    "github.com/StefanosChaliasos/evm-smart-contracts-similarity/src/analysis"
)

func main() {
	var args struct {
        Path         string `arg:"positional,required" help:"Path to directory containing bytecodes"` 
		Threads      int64  `default:"16"`
	}

	arg.MustParse(&args)

    log.Println("Read files from:", args.Path)
    bytecodeFilePaths := utils.WalkDirectoryForFiles(args.Path, ".bin")
    totalAddresses := len(bytecodeFilePaths)
    log.Println("Total files detected:", totalAddresses)
    bytecodes := utils.ReadFiles(bytecodeFilePaths)

    withoutEmpty, emptyFiles, clustersSize, clustersNumber := analysis.IdenticalAnalysis(bytecodes, true)
    analysis.PrintResults(totalAddresses, emptyFiles, clustersSize, clustersNumber)
    fmt.Println()

    _, _, opcodesClustersSize, opcodesClustersNumber := analysis.DisassembledWithoutArgumentsAnalysis(withoutEmpty)
    analysis.PrintResults(totalAddresses, emptyFiles, opcodesClustersSize, opcodesClustersNumber)
}
