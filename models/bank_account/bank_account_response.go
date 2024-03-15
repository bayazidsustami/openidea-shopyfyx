package bank_account_model

type BankAccountData struct {
	BankAccountId     int    `json:"bankAccountId"`
	BankName          string `json:"bankName"`
	BankAccountName   string `json:"bankAccountName"`
	BankAccountNumber string `json:"bankAccountNumber"`
}

type BankAccountsByUserIdResponse struct {
	Message string            `json:"message"`
	Data    []BankAccountData `json:"data"`
}
