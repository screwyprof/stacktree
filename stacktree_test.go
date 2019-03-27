package stacktree_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/screwyprof/stacktree"
)

func TestPrintStackTrace(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		output string
	}{
		{"ItShouldPrintAnInvocationOfAFunction", "main", "1 main\n"},
		{"ItShouldPrintAnInvocationOfAnotherFunction", "foo", "1 foo\n"},
		{"ItShouldPrintInvocationsOfAFunction", "main\nmain\nmain", "3 main\n"},
		{"ItShouldPrintAnInvocationOfANestedFunction", "main, workloop", "1 main\n\t1 workloop\n"},
		{
			"ItShouldPrintInvocationsOfADeeplyNestedFunction",
			"main, workloop, parse_data, parse_entry",
			"1 main\n\t1 workloop\n\t\t1 parse_data\n\t\t\t1 parse_entry\n",
		},

		{
			"ItShouldPrintInvocationsOfTwoDifferentStackTraceLines",
			"main, workloop, select\nmain, parse_args",
			"2 main\n\t1 workloop\n\t\t1 select\n\t1 parse_args\n",
		},
		{
			"ItShouldPrintInvocationsOfTwoEqualStackTraceLines",
			"main, workloop, select\nmain, workloop, select",
			"2 main\n\t2 workloop\n\t\t2 select\n",
		},
		{
			"ItShouldPrintInvocationsOfArbitraryStackTraceLines",
			"main, workloop, select\nmain, parse_args\nmain, workloop, parse_data, parse_entry\nmain, workloop, select",
			"4 main\n\t3 workloop\n\t\t1 parse_data\n\t\t\t1 parse_entry\n\t\t2 select\n\t1 parse_args\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buffer := bytes.Buffer{}

			stacktree.PrintStackTrace(tc.input, &buffer)

			equals(t, tc.output, buffer.String())
		})
	}

}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Errorf("\033[31m\n\n\texp:\n\t%#+v\n\n\tgot:\n\t%#+v\033[39m", exp, act)
	}
}
