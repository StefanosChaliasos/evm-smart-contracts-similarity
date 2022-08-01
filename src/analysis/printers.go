package analysis

import (
    "fmt"

    "github.com/StefanosChaliasos/evm-smart-contracts-similarity/src/utils"
)

func PrintResults (totalAddresses int, clustersSize []int, clustersNumber int) {
    clusteredAddresses := utils.SumIntSlice(clustersSize)
    nonClusteredAddresses := totalAddresses - clusteredAddresses
    fmt.Println()
    fmt.Println("===============Results===============")
    fmt.Println("Total addresses               ", totalAddresses)
    fmt.Println("Number of clusters            ", clustersNumber)
    fmt.Println("Total clustered addresses     ", clusteredAddresses)
    fmt.Println("Total unclustered addresses   ", nonClusteredAddresses)
    fmt.Println("=====================================")
}
