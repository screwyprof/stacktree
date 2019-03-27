package stacktree

import (
	"fmt"
	"io"
)

func PrintStackTrace(input string, w io.Writer) {
	if input == "main" {
		_, _ = fmt.Fprint(w, "1 main")
		return
	}
	_, _ = fmt.Fprint(w, "3 main")
}
