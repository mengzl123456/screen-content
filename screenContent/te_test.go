package screenContent

import (
	"fmt"
	"strings"
	"testing"
)

func TestSpilt(t *testing.T) {
	str := "test12314516"
	fmt.Println(strings.SplitN(str, "1", 3))
}
