package cmd

import (
	"github.com/spf13/cobra"
	"go-es/internal/cmd/make"
)

var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	CmdMake.AddCommand(
		make.CmdMakeModel,
		make.CmdMakeDto,
		make.CmdMakeAPIHandle,
		make.CmdMakeMigration,
	)
}
