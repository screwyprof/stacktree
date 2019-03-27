package stacktree_test

import (
	"os"

	"github.com/screwyprof/stacktree"
)

func Example() {
	const input = `main, workloop, select
main, parse_args
main, workloop, parse_data, parse_entry
main, workloop, select`

	stacktree.PrintStackTrace(input, os.Stdout)

	// Output:
	// 4 main
	//	3 workloop
	//		1 parse_data
	//			1 parse_entry
	//		2 select
	//	1 parse_args
	//
}
