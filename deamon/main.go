package main

import (
	"encoding/json"
	"fmt"
	"github.com/PayLaterCLI/constants"
	"github.com/PayLaterCLI/deamon/API"
	"github.com/PayLaterCLI/deamon/db"
	"github.com/PayLaterCLI/models"
	"net"
	"os"

	"github.com/sirupsen/logrus"
)

type APIDep struct {
	userAPI     API.User
	transaction API.Transaction
	reports     API.Reports
}

func init() {
	//init DB collections
	db.InitDB()
}

func main() {

	logger := logrus.New()

	arguments := os.Args
	if len(arguments) == 1 {
		logger.Fatalln("oops: invalid port number")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", arguments[1]))
	if err != nil {
		logger.Fatalf("daemon failed to start - %s", err.Error())
	}

	defer listener.Close()

	logger.Infoln("server listening on 4444")
	stopSrv := make(chan struct{})

	db := db.NewDB(logger)
	userAPI := API.NewUser(logger, db)
	txn := API.NewTransaction(db, logger)
	report := API.NewReports(db, logger)

	dep := APIDep{
		userAPI:     userAPI,
		transaction: txn,
		reports:     report,
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			stopSrv <- struct{}{}
			panic(err)
		}

		go handleTCPRequest(conn, logger, stopSrv, dep)

		select {
		case <-stopSrv:
			logrus.Println("bye..")
			os.Exit(1)
		default:
			continue
		}

	}
}

func handleTCPRequest(conn net.Conn, logger *logrus.Logger, stopSrv chan struct{}, dep APIDep) {

	for {
		request := models.DataTransportRequest{}
		if err := json.NewDecoder(conn).Decode(&request); err != nil {
			logger.Printf("error: handleConnection - %s", err.Error())
			stopSrv <- struct{}{}
		}

		response := &models.DataTransportResponse{
			Method: request.Method,
			Code:   constants.InternalServerError,
			Error:  "something went wrong, please try after sometime",
		}


		switch request.Method {
		case constants.AddUser:
			logger.Infoln("journey started for adding user - ", request.UserData.Name)
			response = dep.userAPI.AddUser(request.UserData)
			handleResponse(conn, response)
			logger.Infoln("journey ended for adding user", request.UserData.Name)
		case constants.AddMerchant:
			logger.Infoln("journey started for adding merchant - ", request.MerchantData.Name)
			response = dep.userAPI.AddMerchant(request.MerchantData)
			handleResponse(conn, response)
			logger.Infoln("journey ended for adding merchant", request.MerchantData.Name)
		case constants.Transaction:
			logger.Infoln("journey started for txn - ", request.Transaction.User, request.Transaction.Merchant)
			response = dep.transaction.NewTxn(request.Transaction.User, request.Transaction.Merchant, request.Transaction.Amount)
			logger.Infof("response - %+v", response)
			handleResponse(conn, response)
			logger.Infoln("journey ended for txn", request.Transaction.User, request.Transaction.Merchant)
		case constants.PayBack:
			logger.Infoln("journey started for PayBack txn - ", request.Transaction.User)
			response = dep.transaction.PayBack(request.Transaction.User, request.Transaction.Amount)
			handleResponse(conn, response)
			logger.Infoln("journey ended for PayBack txn", request.Transaction.User)
		case constants.Dues:
			logger.Infoln("journey started for getting user dues - ", request.Reports.User)
			val, ok  := dep.reports.ReportUserDues(request.Reports.User)
			if !ok {
				response := &models.DataTransportResponse{
					Method: constants.Dues,
					Code:   constants.Failure,
					Error:  "user not found",
				}
				handleResponse(conn, response)
				break
			}

			response := &models.DataTransportResponse{
				Method: constants.Dues,
				Code:   constants.Success,
				Message:  fmt.Sprintf("dues: %f", val),
			}
			handleResponse(conn, response)
			logger.Infoln("journey ended for getting user dues", request.Reports.User)
		case constants.CreditLimit:
			logger.Infoln("journey started for getting all user credit limit ")

			info := dep.reports.UsersAtCreditLimit()
			response := &models.DataTransportResponse{
				Method:constants.CreditLimit,
				Data: info,
			}
			handleResponse(conn, response)
			logger.Infoln("journey started for getting all user credit limit ")
		case constants.TotalDues:
			logger.Infoln("journey started for getting total-dues ")

			info := dep.reports.ReportTotalDues()
			response := &models.DataTransportResponse{
				Method:constants.CreditLimit,
				Data: info,
			}
			handleResponse(conn, response)
			logger.Infoln("journey started for getting total-dues ")
		}
		break
	}

	conn.Close()
}

func handleResponse(conn net.Conn, response *models.DataTransportResponse) {
	if err := json.NewEncoder(conn).Encode(response); err != nil {
		logrus.Errorf("handleResponse:", err.Error())
	}
}
