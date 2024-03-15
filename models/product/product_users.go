package product_model

import bank_account_model "openidea-shopyfyx/models/bank_account"

type ProductUsers struct {
	Product       Product
	Name          string
	PurchaseTotal int
	BankAccounts  []bank_account_model.BankAccount
}
