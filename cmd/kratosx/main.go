package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/gyq14/kratosx/cmd/kratosx/internal/change"
	"github.com/gyq14/kratosx/cmd/kratosx/internal/project"
	"github.com/gyq14/kratosx/cmd/kratosx/internal/proto"
	"github.com/gyq14/kratosx/cmd/kratosx/internal/run"
	"github.com/gyq14/kratosx/cmd/kratosx/internal/upgrade"
	"github.com/gyq14/kratosx/cmd/kratosx/internal/webutil"
)

var rootCmd = &cobra.Command{
	Use:     "kratosx",
	Short:   "Kratosx: An elegant toolkit for Go micro services.",
	Long:    `Kratosx: An elegant toolkit for Go micro services.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(change.CmdChange)
	rootCmd.AddCommand(run.CmdRun)
	rootCmd.AddCommand(webutil.CmdWebUtil)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
