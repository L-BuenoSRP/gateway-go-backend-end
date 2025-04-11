package dto

import (
	"time"

	"github.com/L-BuenoSRP/imersao25/go-gateway/internal/domain"
)

const (
	InvoiceStatusPending  = string(domain.StatusPending)
	InvoiceStatusApproved = string(domain.StatusApproved)
	InvoiceStatusRejected = string(domain.StatusRejected)
)

type CreateInvoiceInput struct {
	APIKey          string
	Amount          float64 `json:"amount"`
	Description     string  `json:"description"`
	PaymentType     string  `json:"payment_type"`
	CardNumber      string  `json:"card_number"`
	CVV             string  `json:"cvv"`
	ExpirationMonth int     `json:"expiration_month"`
	ExpirationYear  int     `json:"expiration_year"`
	CardHolderName  string  `json:"card_holder_name"`
}

type CreateInvoiceOutput struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
	PaymentType    string    `json:"payment_type"`
	CardLastDigits string    `json:"card_last_digits"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ToInvoice(input CreateInvoiceInput, accountID string) (*domain.Invoice, error) {
	card := domain.CreditCard{
		Number:         input.CardNumber,
		CVV:            input.CVV,
		ExpiryMonth:    input.ExpirationMonth,
		ExpiryYear:     input.ExpirationYear,
		CardholderName: input.CardHolderName,
	}
	return domain.NewInvoice(accountID, input.Amount, input.Description, input.PaymentType, &card)
}

func FromInvoice(invoice *domain.Invoice) *CreateInvoiceOutput {
	return &CreateInvoiceOutput{
		ID:             invoice.ID,
		AccountID:      invoice.AccountID,
		Amount:         invoice.Amount,
		Status:         string(invoice.Status),
		Description:    invoice.Description,
		PaymentType:    invoice.PaymentType,
		CardLastDigits: invoice.CardLastDigits,
		CreatedAt:      invoice.CreatedAt,
		UpdatedAt:      invoice.UpdatedAt,
	}
}
