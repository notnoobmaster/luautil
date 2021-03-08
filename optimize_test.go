package beautifier

import (
	"os"
	"strings"
	"testing"
	_ "embed"

	"github.com/yuin/gopher-lua/parse"
)

//go:embed test.lua
var test string

func TestOptimize(t *testing.T) {
	chunk, err := parse.Parse(strings.NewReader(test), "")
	if err != nil {
		t.Fatal(err)
	}
	
	Optimize(chunk)
	
	file, err := os.Create("optimized.lua")

	file.WriteString(Beautify(chunk))
}