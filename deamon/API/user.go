package API

import (
	"github.com/PayLaterCLI/constants"
	"github.com/PayLaterCLI/deamon/db"
	"github.com/PayLaterCLI/models"
	"github.com/sirupsen/logrus"
)

type User interface {
	AddUser(user *models.User) *models.DataTransportResponse
	AddMerchant(user *models.Merchant) *models.DataTransportResponse
}

type userApi struct {
	log *logrus.Logger
	db  db.DB
}

func (u userApi) AddMerchant(merchant *models.Merchant) *models.DataTransportResponse {
	if ok := u.db.InsertData(constants.AddMerchant, merchant); ok {
		return &models.DataTransportResponse{
			Method:constants.AddMerchant,
			Code: constants.Success,
			Message: "Merchant added successfully!",
		}
	}

	return &models.DataTransportResponse{
		Method:constants.AddMerchant,
		Code: constants.Failure,
		Message: "Merchant add failed!",
	}
}

func (u userApi) AddUser(user *models.User) *models.DataTransportResponse {
	if u.db.CheckDataExists(constants.CollectionUser, user.Name) {
		return &models.DataTransportResponse{
			Method:constants.AddUser,
			Code: constants.Failure,
			Message: "ooh! you are already a member.",
		}
	}

	if u.db.InsertData(constants.AddUser, user) {
		return &models.DataTransportResponse{
			Method:constants.AddUser,
			Code: constants.Success,
			Message: "User added successfully!",
		}
	}

	return nil
}

func NewUser(log *logrus.Logger, db db.DB) User {
	return &userApi{
		log: log,
		db:  db,
	}
}
