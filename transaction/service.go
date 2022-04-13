package transaction

import (
	"errors"
	"website-crowdfunding/campaign"
)

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
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
	return newTransaction, nil
}