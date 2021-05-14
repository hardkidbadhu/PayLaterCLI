package cmd

import (
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "on-boards new user/merchant or perform transactions",
	Long: `used to create new user or merchant with the sub command user / merchant or perform transactions`,

}

func init() {
	rootCmd.AddCommand(newCmd)
}
