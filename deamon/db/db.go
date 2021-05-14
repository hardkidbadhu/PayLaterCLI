package db

import (
	"github.com/PayLaterCLI/constants"
	"github.com/PayLaterCLI/models"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	Once               sync.Once
	UserCollection     map[string]*models.User
	MerchantCollection map[string]*models.Merchant
)

type DB interface {
	InsertData(collection constants.Method, data interface{}) bool
	CheckDataExists(collection constants.Method, key string) bool
	GetUserData(key string) *models.User
	GetMerchantData(key string) *models.Merchant
}

type db struct{
	log *logrus.Logger
}

func (d db) GetMerchantData(key string) *models.Merchant {
	m := &models.Merchant{}

	merchantData, ok := MerchantCollection[key]
	if !ok {
		return nil
	}

	defer func() {
		merchantData.Unlock()
		MerchantCollection[key] = merchantData
	}()

	merchantData.Lock()
	MerchantCollection[key] = merchantData

	m = merchantData
	return m
}

func (d db) GetUserData(key string) *models.User {
	k := &models.User{}

	userData, ok := UserCollection[key]
	if !ok {
		return nil
	}

	defer func() {
		userData.Unlock()
		UserCollection[key] = userData
	}()

	userData.Lock()
	UserCollection[key] = userData

	k = userData
	return k
}

func (d db) CheckDataExists(collection constants.Method, key string) bool {
	switch collection {
	case constants.CollectionUser:
		if _, ok := UserCollection[key]; ok {
			return true
		}
		return false
	}
	return false
}

func (d db) InsertData(collection constants.Method, data interface{}) bool {
	switch collection {
	case constants.AddUser:
		if val, ok := data.(*models.User); ok {
			UserCollection[val.Name] = val
		}
		d.log.Infof("User data - %+v", UserCollection)
		return true

	case constants.AddMerchant:
		if val, ok := data.(*models.Merchant); ok {
			MerchantCollection[val.Name] = val
		}
		d.log.Infof("Merchant data - %+v", MerchantCollection)
		return true
	}

	return false
}

func (d db) UpdateData() {

}

func NewDB(	log *logrus.Logger) DB {
	return &db{
		log : log,
	}
}

func InitDB() {
	Once.Do(func() {
		UserCollection = make(map[string]*models.User)
		MerchantCollection = make(map[string]*models.Merchant)
	})
}
