package cmd

import (
	"fmt"
	"github.com/PayLaterCLI/cmd/tcp"
	"github.com/PayLaterCLI/constants"
	"github.com/PayLaterCLI/models"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// duesCmd represents the dues command
var duesCmd = &cobra.Command{
	Use:   "dues",
	Short: "fetches the dues of user",
	Long: `sample: report dues u1`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("please enter valid args usage `report dues u1`")
			return
		}

		reqData := &models.DataTransportRequest{
			Method: constants.Dues,
			Reports: &models.Report{
				User: args[0],
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

		fmt.Println(response.Message)
		return
	},
}

func init() {
	reportCmd.AddCommand(duesCmd)
}
