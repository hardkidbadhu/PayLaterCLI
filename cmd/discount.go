package cmd

import (
	"fmt"
	"github.com/PayLaterCLI/cmd/tcp"
	"github.com/PayLaterCLI/constants"
	"github.com/PayLaterCLI/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// discountCmd represents the discount command
var discountCmd = &cobra.Command{
	Use:   "discount",
	Short: "fetches the merchant discount",
	Long: `use like : report discount m1`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			fmt.Println("please enter valid args usage report discount m1")
			return
		}

		reqData := &models.DataTransportRequest{
			Method: constants.Discount,
			Reports: &models.Report{
				Merchant: args[0],
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
	reportCmd.AddCommand(discountCmd)
}
