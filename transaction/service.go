package transaction

import (
	"errors"
	"website-crowdfunding/campaign"
	"website-crowdfunding/payment"
)

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository         	Repository
	campaignRepository 	campaign.Repository
	paymentService		payment.Service
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByCampaignID(input.ID)
	if err != nil{
		return []Transaction{}, err 
	}

	if campaign.ID != input.User.ID{
		return []Transaction{}, errors.New("not an owner of the campaign")
	}

	transactions, err := s.repository.GetCampaignByID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionByUserID(userID int) ([]Transaction, error){
	transactions, err := s.repository.GetUserByID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error){
	transaction := Transaction{}
	transaction.Amount		= input.Amount
	transaction.CampaignID	= input.CampaignID
	transaction.UserID		= input.User.ID
	transaction.Status		= "pending"
	transaction.Code		= ""

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	// mapping ke struct transaction di package payment
	paymentTransaction := payment.Transaction{
		ID: newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	// dapatkan data payment url di package payment
	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil
}