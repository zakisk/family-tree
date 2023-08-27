package cmd

import (
	"github.com/spf13/cobra"
)

// FamilyTreeCmd represents the base command when called without any subcommands
var FamilyTreeCmd = &cobra.Command{
	Use:   "family-tree",
	Short: "family tree utility CLI app",
	Long: `An utility tool to create and view family tree records.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil	
	},
}

func Execute() error {
	err := FamilyTreeCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	FamilyTreeCmd.AddCommand(AddCmd, ConnectRelationshipCmd, CountRelativesCmd, FatherCmd)
}


