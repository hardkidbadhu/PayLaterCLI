package constants

const (
	Regexp = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

type (
	Method string
	Code   int
)

const (

	//action requests
	AddUser        Method = `createUser`
	CollectionUser        = `user`
	CollectionMerchant    = `merchant`
	AddMerchant           = `addMerchant`
	Transaction           = `txn`
	PayBack               = `payback`
	Report                = `report`
	Discount              = `discount`
	Dues                  = `dues`
	CreditLimit           = `usersCreditLimit`
	TotalDues             = `total_dues`
)

const (
	Success             Code = 200
	Failure                  = 400
	InternalServerError      = 500
)
