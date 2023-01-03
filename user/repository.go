package user

import "gorm.io/gorm"

type Repository interface {
	//jika awal huruf besar bisa diakses package lain
	Save(user User) (User, error)
}

type repository struct {
	//jika awal huruf besar tidak bisa diakses package lain
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
