package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

// function agar db di struct repository terhubung dan terisi 
// data dari database/ db yang di main.go
func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) Save(user User)(User, error){

	// menulis nilai ke database
	err := r.db.Create(&user).Error
	if err != nil{
		return user, err
	}
	return user, nil
}