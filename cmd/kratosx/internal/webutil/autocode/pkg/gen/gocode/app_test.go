package gocode

import (
	"encoding/json"
	"fmt"
	"github.com/gyq14/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
	"github.com/gyq14/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/types"
	"os"
	"testing"
)

func TestGenApp(t *testing.T) {
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

	builder := NewAppBuilder(gen.NewBuilder(nil, initTable()))
	code, err := builder.GenApp()
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
}
