package page_model_test_tools

import (
	"fmt"

	pgm "github.com/terryhay/dolly/man_style_help/page_model"
	ri "github.com/terryhay/dolly/man_style_help/row_iterator"
)

// PrintTerminalContentState inputs debug info
func PrintTerminalContentState(pgm *pgm.PageModel) {
	ts := pgm.GetUsingTermSize().GetWidthLimit()
	fmt.Printf("\n\033[36mTERMINAL: w=%d, h=%d\033[0m",
		ts.Max().Int(), pgm.GetUsingTermSize().GetHeight().Int(),
	)
	fmt.Println("\033[36m  |0         |10        |20        |30        |40        |50\033[0m")

	counter := 0
	for it := ri.MakeRowIterator(pgm); !it.End(); it.Next() {
		counter++
		s := it.RowModel().String()

		if counter < 10 {
			fmt.Println(fmt.Sprintf("\u001B[36m %d|\u001B[0m", counter) + s)
			continue
		}

		fmt.Println(fmt.Sprintf("\u001B[36m%d|\u001B[0m", counter) + s)
	}
	fmt.Println()
}
