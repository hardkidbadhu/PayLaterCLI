package cmd

import (
	"github.com/spf13/cobra"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "simple reporting cli",
	Long: `used to fetch dues, discount and total reports`,
}

func init() {
	rootCmd.AddCommand(reportCmd)

}
