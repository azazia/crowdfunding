package transaction

import "gorm.io/gorm"

type Repository interface {
	GetCampaignByID(campaignID int) ([]Transaction, error)
	GetUserByID(userID int) ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCampaignByID(campaignID int)([]Transaction, error){
	// agar gorm tahu kita akan mencari data di tabel transactions
	var transactions []Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&transactions).Error
	if err != nil{
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) GetUserByID(userID int) ([]Transaction, error){
	var transactions []Transaction

	// untuk memuat sebuah relasi yang tidak terkait langsung dengan objek transactions
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = true").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error
	if err != nil{
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) Save(transaction Transaction) (Transaction, error){
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error){
	err := r.db.Save(&transaction).Error
	if err != nil{
		return transaction, err
	}
	return transaction, nil
}