package bank_account_model

type BankAccountData struct {
	BankAccountId     int    `json:"bank_account_id"`
	BankName          string `json:"bank_name"`
	BankAccountName   string `json:"bank_account_name"`
	BankAccountNumber string `json:"bank_account_number"`
}

type BankAccountResponse struct {
	Data    BankAccountData `json:"data"`
	Message string          `json:"string"`
}
