package cmd

import (
	"fmt"
	"github.com/PayLaterCLI/cmd/tcp"
	"github.com/PayLaterCLI/constants"
	"github.com/PayLaterCLI/models"
	"github.com/sirupsen/logrus"
	"strconv"

	"github.com/spf13/cobra"
)

// paybackCmd represents the payback command
var paybackCmd = &cobra.Command{
	Use:   "payback",
	Short: "repay the dues",
	Long: `Pay back the dues 
		payback user3 400
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			fmt.Println("enter user and amount")
			return
		}

		txAmount, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println("ohh!!invalid amount")
		}
		reqData := &models.DataTransportRequest{
			Method: constants.PayBack,
			Transaction: &models.Transaction{
				User: args[0],
				Amount: txAmount,
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
			fmt.Println(response.Error)
			return
		}

		fmt.Println(response.Message)
		return

	},
}

func init() {
	rootCmd.AddCommand(paybackCmd)
}
