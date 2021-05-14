package cmd

import (
	"fmt"
	"github.com/PayLaterCLI/cmd/tcp"
	"github.com/PayLaterCLI/constants"
	"github.com/PayLaterCLI/models"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// totalDuesCmd represents the totalDues command
var totalDuesCmd = &cobra.Command{
	Use:   "total-dues",
	Short: "get total dues",
	Long: `report total-dues`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("no args required usage report total-dues")
			return
		}

		reqData := &models.DataTransportRequest{
			Method: constants.TotalDues,
			Reports: &models.Report{
				TotalDues: true,
			},
		}

		log := logrus.New()

		tp, err := tcp.Connect(log)
		if err != nil {
			fmt.Println("ohh! try again after some time.")
			return
		}

		response, _ := tp.MakeRequest(reqData)
		if response.Error != ""{
			fmt.Println(response.Message)
			return
		}

		fmt.Println(response.Data)
		return
	},
}

func init() {
	rootCmd.AddCommand(totalDuesCmd)
}
