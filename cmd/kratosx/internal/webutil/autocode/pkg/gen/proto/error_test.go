package proto

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
	"kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/types"
)

func TestGenError(t *testing.T) {
	initTable := func() *types.Table {
		var table types.Table
		content, err := os.ReadFile("internal/webutil/autocode/pkg/gen/auto.json")
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(content, &table); err != nil {
			panic(err)
		}
		return &table
	}

	builder := NewErrorBuilder(gen.NewBuilder(nil, initTable()))
	code, err := builder.GenError()
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
}
