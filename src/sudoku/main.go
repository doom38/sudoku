package main

import (
	"fmt"
	"sudoku/objects"
	"sudoku/solverV1"
)

func main() {
	grid, _ := objects.LoadGrid9x9("016400000200009000400000062070230100100000003003087040960000005000800007000006820")
	// grid, _ := objects.LoadGrid9x9("165794038407002050930006004810405002576239400200601075301507849690000527050028103")

	solverV1 := solverV1.New(grid)
	solverV1.Dump()
	solverV1.Solve()

	fmt.Println("\n\n" + grid.String())
}
