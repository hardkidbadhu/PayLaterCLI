package cmd

import (
	"fmt"
	"github.com/PayLaterCLI/cmd/tcp"
	"github.com/PayLaterCLI/constants"
	"github.com/PayLaterCLI/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

// txnCmd represents the txn command
var txnCmd = &cobra.Command{
	Use:   "txn",
	Short: "perform a transaction",
	Long: `make user transaction with respective to merchant
	sample : new txn u1 m2 400
	where u1 - user 
	m2 - merchant 
	400 - amount of transaction
	`,
	Run: func(cmd *cobra.Command, args []string) {

		txAmount, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			fmt.Println("ohh!!invalid amount")
		}
		reqData := &models.DataTransportRequest{
			Method: constants.Transaction,
			Transaction: &models.Transaction{
				User: args[0],
				Merchant: args[1],
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
			fmt.Println(response.Message)
			return
		}

		fmt.Println(response.Message)
		return
	},
}

func init() {
	newCmd.AddCommand(txnCmd)
}
