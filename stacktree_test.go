package stacktree_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/screwyprof/stacktree"
)

//1. Parse a list of stack traces
//2. Parse stack trace
//3. Parse functions
//4. Add function as a Tree Node
//5. Count function invocations
//6. Print an intended diagram to Stdout

// Input:
// main, workloop, select
// main, parse_args
// main, workloop, parse_data, parse_entry
// main, workloop, select

// Output:
// 4 main
//    3 workLoop
//        1 parse_data
//            1 parse_entry
//        2 select
//    1 parse_args

func TestPrintStackTrace(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		output string
	}{
		{"ItShouldPrintAnInvocationOfAFunction", "main", "1 main"},
		{"ItShouldPrintAnInvocationOfAnotherFunction", "foo", "1 foo"},
		{"ItShouldPrintInvocationsOfAFunction", "main\nmain\nmain", "3 main"},
		//{"ItShouldPrintAnInvocationOfANestedFunction", "main workloop", "1 main\n\tworkloop"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buffer := bytes.Buffer{}

			stacktree.PrintStackTrace(tc.input, &buffer)

			equals(t, tc.output, buffer.String())
		})
	}

}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatalf("\033[31munexpected error: %v\033[39m\n\n", err)
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Errorf("\033[31m\n\n\texp:\n\t%#+v\n\n\tgot:\n\t%#+v\033[39m", exp, act)
	}
}
