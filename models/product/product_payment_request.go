package product_model

type ProductPaymentRequest struct {
	BankAccountId        int    `json:"bankAccountId" validate:"required,number"`
	PaymentProofImageUrl string `json:"paymentProofImageUrl" validate:"required,imageurl"`
	Quantity             int    `json:"quantity" validate:"required,number"`
}
