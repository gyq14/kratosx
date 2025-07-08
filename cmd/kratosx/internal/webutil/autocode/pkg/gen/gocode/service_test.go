package gocode

import (
	"encoding/json"
	"fmt"
	"kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
	"kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/types"
	"os"
	"testing"
)

func TestGenService(t *testing.T) {
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

	builder := NewServiceBuilder(gen.NewBuilder(nil, initTable()))
	code, err := builder.GenService()
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
}
