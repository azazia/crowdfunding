package payment

import (
	"strconv"
	"website-crowdfunding/campaign"
	"website-crowdfunding/transaction"
	"website-crowdfunding/user"

	"github.com/veritrans/go-midtrans"
)

type service struct {
	transactionRepository transaction.Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
	PaymentProcess(input transaction.TransactionNotificationInput) error
}

func NewService(transactionRepository transaction.Repository, campaignRepository campaign.Repository) *service{
	return &service{transactionRepository, campaignRepository}
}

// melakukan integrasi dengan sistem midtrans
func (s *service)GetPaymentURL(transaction Transaction, user user.User) (string, error){
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-O8N9Iq06pkcEj42AUn5v_-lD"
	midclient.ClientKey = "SB-Mid-client-spE79BCeFwkxj50F"
	midclient.APIEnvType = midtrans.Sandbox  

	snapeGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapeGateway.GetToken(&snapReq)
	if err != nil{
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}

func (s *service) PaymentProcess(input transaction.TransactionNotificationInput) error{
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.transactionRepository.GetByID(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" &&  input.FraudStatus == "accept" && input.TransactionStatus == "capture" {
		transaction.Status = "paid"
	}else if input.TransactionStatus == "settlement"{
		transaction.Status = "paid"
	}else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel"{
		transaction.Status = "cancelled"
	}

	updateTransaction, err := s.transactionRepository.Update(transaction)
	if err != nil {
		return err
	}

	// buat update backer count dan amount di campaign
	campaign, err := s.campaignRepository.FindByCampaignID(updateTransaction.CampaignID)
	if err != nil {
		return err
	}

	if updateTransaction.Status == "paid"{
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updateTransaction.Amount

		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}