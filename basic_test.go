package shuf_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/shuf"
)

func ExampleShuf_basic() {
	// echo "1\n2\n3\n4\n5" | shuf -n 2
	// Note: Output is random, so we just demonstrate usage
	yup.MustRun(
		Shuf(Count(2), strings.NewReader("1\n2\n3\n4\n5")),
	)
}

