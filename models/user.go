package models

import (
	"github.com/PayLaterCLI/constants"
	"log"
	"regexp"
	"sync"
)

type User struct {
	sync.RWMutex
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	CreditLimit float64 `json:"credit_limit"`
	Dues        float64 `json:"dues"`
}

/**
 * @Description: Validate returns the user is valid or not.
 * @receiver *User
 * @return bool
 */
func (u *User) Validate() bool {
	if len(u.Name) <= 0 {
		return false
	}

	rxp := regexp.MustCompile(constants.Regexp)

	if len(u.Email) < 3 && len(u.Email) > 100 {
		return false
	}

	if !rxp.MatchString(u.Email) {
		return false
	}

	return true
}

/**
 * @Description: CalculateDues calculates the remaining dues in terms of transactions of payback.
 * @receiver *User
 * @param txnAmt
 * @param txnType constants.Method
 * @param discount
 */
func (u *User) CalculateDues(txnAmt, discount float64, txn constants.Method) bool {

	switch txn {
	case constants.Transaction:
		discountAmt := txnAmt * (discount / 100)
		amt := txnAmt - discountAmt

		log.Println("discount", discountAmt)
		log.Println("amount", amt)

		avaliableCredit := u.CreditLimit - u.Dues

		log.Println("avaliableCredit", avaliableCredit)

		//checks the amount exceeds the limit
		if txnAmt > u.CreditLimit || avaliableCredit < amt {
			return false
		}

		u.Dues += amt
		return true

	case constants.PayBack:
		u.Dues -= txnAmt
		return true
	}
	return false
}

func (u *User) UserAtCreditLimit() bool {
	if u.CreditLimit > u.Dues || u.Dues == 0.0 {
		return true
	}

	return false
}