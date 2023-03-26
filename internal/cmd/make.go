package cmd

import (
	"github.com/spf13/cobra"
	make2 "go-es/internal/cmd/make"
)

var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	CmdMake.AddCommand(
		make2.CmdMakeModel,
	)
}
