package domain

import "errors"

var (
	// ErrAccountNotFound is returned when an account is not found
	ErrAccountNotFound = errors.New("account not found")
	// ErrDuplicatedApiKey is returned when an account with the same API key already exists
	ErrDuplicatedApiKey = errors.New("duplicated api key")
	// ErrInvoiceNotFound is returned when an invoice is not found
	ErrInvoiceNotFound = errors.New("invoice not found")
	// ErrUnauthorizedAccess is returned when an unauthorized access is attempted
	ErrUnauthorizedAccess = errors.New("unauthorized access")
	// ErrInsufficientBalance is returned when a transaction is attempted with insufficient balance
	ErrInsufficientBalance = errors.New("insufficient balance")
	// ErrAmountMustBeGreaterThanZero is returned when the amount is not greater than 0
	ErrAmountMustBeGreaterThanZero = errors.New("amount must be greater than 0")
	// ErrDescriptionIsRequired is returned when the description is not provided
	ErrDescriptionIsRequired = errors.New("description is required")
	// ErrInvalidAmount is returned when the amount is not valid
	ErrInvalidAmount = errors.New("invalid amount")
	// ErrInvalidStatus is returned when the status is not valid
	ErrInvalidStatus = errors.New("invalid status")
	// ErrInvalidPaymentType is returned when the payment type is not valid
	ErrInvalidPaymentType = errors.New("invalid payment type")
	// ErrInvalidCreditCard is returned when the credit card is not valid
	ErrInvalidCreditCard = errors.New("invalid credit card")
	// ErrInvalidCreditCardNumber is returned when the credit card number is not valid
	ErrInvalidCreditCardNumber = errors.New("invalid credit card number")
	// ErrInvalidCreditCardCVV is returned when the credit card CVV is not valid
	ErrInvalidCreditCardCVV = errors.New("invalid credit card CVV")
)
