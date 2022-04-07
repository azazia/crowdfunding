package transaction

type Service interface{
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput)([]Transaction, error)
}

type service struct{
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput)([]Transaction, error){
	transactions, err := s.repository.GetCampaignByID(input.ID)
	if err != nil{
		return transactions, err
	}
	return transactions, nil
}