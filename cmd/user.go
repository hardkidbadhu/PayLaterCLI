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

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "create a new user",
	Long: `sample : new user u1 u1@email.in 1000
	# name, email, credit-limit`,

	Run: func(cmd *cobra.Command, args []string) {

		credit, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			fmt.Println("enter valid discount percentage")
		}
		userIns := &models.User{
			Name:        args[0],
			Email:       args[1],
			CreditLimit: credit,
		}

		if !userIns.Validate() {
			fmt.Println("ouch! please enter a valid data")
			return
		}
		log := logrus.New()
		tp, err := tcp.Connect(log)
		if err != nil {
			fmt.Println("ohh! try again after some time.")
			return
		}

		response, _ := tp.MakeRequest(&models.DataTransportRequest{
			Method:   constants.AddUser,
			UserData: userIns,
		})

		if response.Code != constants.Success {
			fmt.Println(response.Message)
			return
		}

		fmt.Printf("welcome on-board %s", userIns.Name)
	},
}

func init() {
	newCmd.AddCommand(userCmd)
}
