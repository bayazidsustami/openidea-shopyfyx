package bank_account_model

type BankAccountRequest struct {
	BankName          string `json:"bankName" validate:"required,min=3,max=30"`
	BankAccountName   string `json:"bankAccountName" validate:"required,min=3,max=30"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required,min=10,max=19"`
}
