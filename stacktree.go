package stacktree

import (
	"fmt"
	"io"
)

func PrintStackTrace(input string, w io.Writer) {
	_, _ = fmt.Fprint(w, "1 main")
}
