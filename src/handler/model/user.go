package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `grom:"type:varchar(30)"`
	Password string `grom:"type:varchar(32)"`
}

func (u *User) FindUser(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&u).Error
}

func (u *User) UserCreate(db *gorm.DB) error {
	return db.Debug().Create(&u).Error
}

type Sp struct {
	gorm.Model
	Name  string `grom:"type:varchar(30)"`
	SpLen string `grom:"type:varchar(30)"`
	SpNum string `grom:"type:varchar(30)"`
}

func (s *Sp) FingSp(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&s).Error
}

func (s *Sp) SpAdd(db *gorm.DB) error {
	return db.Debug().Create(&s).Error
}
