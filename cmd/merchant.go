package cmd

import (
	"fmt"
	"github.com/PayLaterCLI/cmd/tcp"
	"github.com/PayLaterCLI/constants"
	"github.com/PayLaterCLI/models"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// merchantCmd represents the merchant command
var merchantCmd = &cobra.Command{
	Use:   "merchant",
	Short: "adds a new merchant",
	Long:  `Adds a new merchant new merchant m1 2%`,
	Run: func(cmd *cobra.Command, args []string) {

		dis, err := strconv.ParseFloat(TrimPercentage(args[2]), 64)
		if err != nil {
			fmt.Println("oouch! Invalid percentage")
			return
		}

		merchant := &models.Merchant{
			Name:            args[0],
			Email:           args[1],
			DiscountPercent: dis,
		}

		log := logrus.New()
		tp, err := tcp.Connect(log)
		if err != nil {
			fmt.Println("ohh! try again after some time.")
			return
		}

		response, _ := tp.MakeRequest(&models.DataTransportRequest{
			Method:       constants.AddMerchant,
			MerchantData: merchant,
		})

		if response.Code != constants.Success {
			fmt.Println(response.Message)
			return
		}

		fmt.Printf("merchant information added / updated %s", merchant.Name)

	},
}

func init() {
	newCmd.AddCommand(merchantCmd)
}

func TrimPercentage(discount string) string {
	if strings.Contains(discount, "%") {
		return strings.Replace(discount, "%", "", -1)
	}
	return discount
}
