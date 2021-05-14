package API

import (
	"fmt"
	"github.com/PayLaterCLI/constants"
	"github.com/PayLaterCLI/deamon/db"
	"github.com/PayLaterCLI/models"
	"github.com/sirupsen/logrus"
)

type Transaction interface {
	NewTxn(user, merchant string, amount float64) *models.DataTransportResponse
	PayBack(user string, amount float64) *models.DataTransportResponse
}

type transaction struct {
	db  db.DB
	log *logrus.Logger
}

func (t transaction) PayBack(user string, amount float64) *models.DataTransportResponse {
	userData, ok := db.UserCollection[user]
	if !ok {
		return &models.DataTransportResponse{
			Method: constants.Transaction,
			Code: constants.Failure,
			Error: "user not found",
		}
	}

	userData.Lock()
	db.UserCollection[user] = userData

	defer func() {
		userData.Unlock()
		db.UserCollection[user] = userData
	}()

	if !userData.CalculateDues(amount, 0, constants.PayBack) {
		return &models.DataTransportResponse{
			Method: constants.PayBack,
			Code: constants.Failure,
			Error: "please try after sometime",
		}
	}

	return &models.DataTransportResponse{
		Method: constants.PayBack,
		Code: constants.Success,
		Message: fmt.Sprintf("%s (dues: %f)", userData.Name, userData.Dues),
	}
}

func (t transaction) NewTxn(user, merchant string, amount float64) *models.DataTransportResponse {
	userData, ok := db.UserCollection[user]
	if !ok {
		return &models.DataTransportResponse{
			Method: constants.Transaction,
			Code: constants.Failure,
			Error: "user not found",
		}
	}

	merchantData, ok := db.MerchantCollection[merchant]
	if !ok {
		userData.Unlock()
		return &models.DataTransportResponse{
			Method: constants.Transaction,
			Code: constants.Failure,
			Error: "merchant not found",
		}
	}

	userData.Lock()
	merchantData.Lock()

	db.UserCollection[user] = userData
	db.MerchantCollection[merchant] = merchantData
	t.log.Infof("Userdata Before transaction - (Dues - %f, CreditLimit - %f)", userData.Dues, userData.CreditLimit)

	defer func() {
		userData.Unlock()
		merchantData.Unlock()

		t.log.Infof("Userdata After transaction - (Dues - %f, CreditLimit - %f)", userData.Dues, userData.CreditLimit)
		//updates the same in map
		db.UserCollection[user] = userData
		db.MerchantCollection[merchant] = merchantData
	}()


	if !userData.CalculateDues(amount, merchantData.DiscountPercent, constants.Transaction) {
		return &models.DataTransportResponse{
			Method: constants.Transaction,
			Code: constants.Failure,
			Error: "specified amount is higher than available credit",
			Message: "rejected! (reason: credit limit)",
		}
	}

	return &models.DataTransportResponse{
		Method: constants.Transaction,
		Code: constants.Success,
		Message: "Success!",
	}
}

func NewTransaction(db  db.DB, log *logrus.Logger) Transaction {
	return &transaction {
		db:  db,
		log: log,
	}
}
