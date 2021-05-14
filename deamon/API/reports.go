package API

import (
	"fmt"
	"github.com/PayLaterCLI/deamon/db"
	"github.com/sirupsen/logrus"
)

type Reports interface {
	GetDiscount(merchant string) (float64, bool)
	UsersAtCreditLimit() []string
	ReportTotalDues() []string
	ReportUserDues(user string) (float64, bool)
}

type report struct {
	db db.DB
	log *logrus.Logger
}

func (r report) GetDiscount(merchant string) (float64,  bool) {

	if data := r.db.GetMerchantData(merchant); data != nil {
		return data.DiscountPercent, true
	}

	return 0.0, false
}

func (r report) UsersAtCreditLimit() []string {
	var users []string

	for key, val := range db.UserCollection {
		val.Lock()
		db.UserCollection[key] = val


		if !val.UserAtCreditLimit() {
			val.Unlock()
			db.UserCollection[key] = val
			continue
		}

		users = append(users, key)
		val.Unlock()
		db.UserCollection[key] = val

	}

	r.log.Println(users)
	return users
}

func (r report) ReportTotalDues() []string {
	var totalDues []string

	for key, val := range db.UserCollection {
		val.Lock()
		db.UserCollection[key] = val

		dueString := fmt.Sprintf("%s: %f", key, val.Dues)
		totalDues = append(totalDues, dueString)

		val.Unlock()
		db.UserCollection[key] = val

	}

	return totalDues
}

func (r report) ReportUserDues(user string) (float64, bool) {
	if data := r.db.GetUserData(user); data != nil {
		return data.Dues, true
	}

	return 0.0, false
}

func NewReports(db db.DB, log *logrus.Logger) Reports {
	return &report{db: db, log: log}
}