package payment

import (
	"strconv"
	"website-crowdfunding/user"
	"github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service{
	return &service{}
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

