package upgrade

import (
	"fmt"

	"kratosx/cmd/kratosx/internal/base"
	"github.com/spf13/cobra"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the kratosx tools",
	Long:  "Upgrade the kratosx tools. Example: kratosx upgrade",
	Run:   Run,
}

// Run upgrade the kratos tools.
func Run(_ *cobra.Command, _ []string) {
	err := base.GoInstall(
		"kratosx/cmd/kratosx@latest",
		"kratosx/cmd/protoc-gen-go-httpx@latest",
		"kratosx/cmd/protoc-gen-go-errorsx@latest",
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
		"github.com/google/gnostic/cmd/protoc-gen-openapi@latest",
		"github.com/envoyproxy/protoc-gen-validate@latest",
	)
	if err != nil {
		fmt.Println(err)
	}
}
