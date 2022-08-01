package main

import (
    "log"
    _ "path/filepath"
	_ "os"
    "fmt"

	arg "github.com/alexflint/go-arg"

    "github.com/StefanosChaliasos/evm-smart-contracts-similarity/src/utils"
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
    fmt.Println("Total files detected:", totalAddresses)
    bytecodes := utils.ReadFiles(bytecodeFilePaths)

    // clusters
    clustersNumber := 0
    var clustersSize []int
    emptyFiles := 0
    clusters := utils.FindIdentical(bytecodes) 
    clustersWithoutEmpty := make(map[string][]string)
    for k, v := range clusters {
        // filter out empty strings
        if k == "" {
            fmt.Println("Addresses with empty values: ", v)
            emptyFiles = len(v)
            continue
        }
        clustersWithoutEmpty[k] = v
        if len(v) > 1 {
            clustersNumber += 1
            clustersSize = append(clustersSize, len(v))
            fmt.Println("Cluster items: ", v)
        }
    }

    clusteredAddresses := utils.SumIntSlice(clustersSize)
    nonClusteredAddresses := totalAddresses - emptyFiles - clusteredAddresses
    fmt.Println()
    fmt.Println("==========Identical Results==========")
    fmt.Println("Total addresses               ", totalAddresses)
    fmt.Println("Number of empty files         ", emptyFiles)
    fmt.Println("Number of identical clusters  ", clustersNumber)
    fmt.Println("Total clustered addresses     ", clusteredAddresses)
    fmt.Println("Total unclustered addresses   ", nonClusteredAddresses)
    fmt.Println("=====================================")
}
