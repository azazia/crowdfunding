package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByCampaignID(ID int) (Campaign, error)
	Save(campaign Campaign)(Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// mendapatkan semua campaign
func (r *repository) FindAll() ([]Campaign, error){
	var campaign []Campaign
	
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaign).Error
	if err != nil{
		return campaign, err
	}

	return campaign, nil
}

// mendapatkan campaign yang dibuat user tertentu
func (r *repository) FindByUserID(userID int) ([]Campaign, error){
	var campaign []Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaign).Error
	if err != nil{
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) FindByCampaignID(ID int) (Campaign, error){
	var campaign Campaign

	err := r.db.Where("id = ?", ID).Preload("User").Preload("CampaignImages").Find(&campaign).Error
	if err != nil{
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error){
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}