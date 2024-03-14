package bank_account_model

type BankAccountRequest struct {
	BankName          string `json:"bank_name" validate:"required,min=3,max=30"`
	BankAccountName   string `json:"bank_account_name" validate:"required,min=3,max=15"`
	BankAccountNumber string `json:"bank_account_number" validate:"required,min=10,max=19"`
}
