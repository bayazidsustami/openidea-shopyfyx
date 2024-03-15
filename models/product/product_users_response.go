package product_model

import bank_account_model "openidea-shopyfyx/models/bank_account"

type ProductUsersResponse struct {
	Product ProductResponse `json:"product"`
	Seller  Seller          `json:"seller"`
}

type Seller struct {
	Name          string                               `json:"name"`
	PurchaseTotal int                                  `json:"productSoldTotal"`
	BankAccounts  []bank_account_model.BankAccountData `json:"bankAccounts"`
}
