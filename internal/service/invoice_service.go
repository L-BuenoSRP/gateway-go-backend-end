package service

import (
	"github.com/L-BuenoSRP/imersao25/go-gateway/internal/domain"
	"github.com/L-BuenoSRP/imersao25/go-gateway/internal/dto"
)

type InvoiceService struct {
	invoiceRepository domain.InvoiceRepository
	accountService    AccountService
}

func NewInvoiceService(invoiceRepository domain.InvoiceRepository, accountService AccountService) *InvoiceService {
	return &InvoiceService{
		invoiceRepository: invoiceRepository,
		accountService:    accountService,
	}
}

func (s *InvoiceService) Create(input dto.CreateInvoiceInput) (*dto.CreateInvoiceOutput, error) {
	accountOutput, err := s.accountService.FindByApiKey(input.APIKey)
	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoice(input, accountOutput.ID)
	if err != nil {
		return nil, err
	}

	if err := invoice.Process(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusApproved {
		_, err := s.accountService.UpdateBalance(input.APIKey, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}

	if err := s.invoiceRepository.Save(invoice); err != nil {
		return nil, err
	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) FindByAccountId(accountId string) ([]*dto.CreateInvoiceOutput, error) {
	invoices, err := s.invoiceRepository.FindByAccountId(accountId)
	if err != nil {
		return nil, err
	}

	output := make([]*dto.CreateInvoiceOutput, len(invoices))
	for i, invoice := range invoices {
		output[i] = dto.FromInvoice(invoice)
	}

	return output, nil
}

func (s *InvoiceService) FindById(id string, apiKey string) (*dto.CreateInvoiceOutput, error) {
	invoice, err := s.invoiceRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	accountOutput, err := s.accountService.FindByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	if invoice.AccountID != accountOutput.ID {
		return nil, domain.ErrUnauthorizedAccess
	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) FindByAccountApiKey(apiKey string) ([]*dto.CreateInvoiceOutput, error) {
	accountOutput, err := s.accountService.FindByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	invoices, err := s.invoiceRepository.FindByAccountId(accountOutput.ID)
	if err != nil {
		return nil, err
	}

	output := make([]*dto.CreateInvoiceOutput, len(invoices))
	for i, invoice := range invoices {
		output[i] = dto.FromInvoice(invoice)
	}

	return output, nil
}

func (s *InvoiceService) UpdateStatus(id string, status domain.Status) error {
	invoice, err := s.invoiceRepository.FindById(id)
	if err != nil {
		return err
	}

	if err := invoice.UpdateStatus(status); err != nil {
		return err
	}

	return s.invoiceRepository.UpdateStatus(invoice)
}
