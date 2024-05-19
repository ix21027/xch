package debug

import (
	"fmt"
	"strings"

	"github.com/ztrue/tracerr"
)


func Err(err error) {
	fmt.Println("ERROR MSG: ", err)
	wrappedErr := tracerr.Wrap(err)
	msg := tracerr.SprintSourceColor(wrappedErr)
	frames := strings.Split(msg, "\n\n")
	fmt.Println(frames[2])
}
