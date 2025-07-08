package main

import (
	"log"

	"github.com/spf13/cobra"

	"kratosx/cmd/kratosx/internal/change"
	"kratosx/cmd/kratosx/internal/project"
	"kratosx/cmd/kratosx/internal/proto"
	"kratosx/cmd/kratosx/internal/run"
	"kratosx/cmd/kratosx/internal/upgrade"
	"kratosx/cmd/kratosx/internal/webutil"
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
