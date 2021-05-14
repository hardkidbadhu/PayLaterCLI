package cmd

import (
	"fmt"
	"github.com/PayLaterCLI/cmd/tcp"
	"github.com/PayLaterCLI/constants"
	"github.com/PayLaterCLI/models"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// usersAtCreditLimitCmd represents the usersAtCreditLimit command
var usersAtCreditLimitCmd = &cobra.Command{
	Use:   "users-at-credit-limit",
	Short: "A brief description of your command",
	Long: `usage: report users-at-credit-limit`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("no args required usage report total-dues")
			return
		}

		reqData := &models.DataTransportRequest{
			Method: constants.CreditLimit,
			Reports: &models.Report{
				UsersAtCreditLimit: true,
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
	rootCmd.AddCommand(usersAtCreditLimitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usersAtCreditLimitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usersAtCreditLimitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
