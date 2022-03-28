package model

type Domain struct {
	ID             uint   `gorm:"primarykey"`
	Name           string `gorm:"index;unique"`
	ExpirationTime int64  `gorm:"index"`
	Forbidden      bool
	UserId         uint `gorm:"index"`

	Updated int64 `gorm:"autoUpdateTime"`
	Created int64 `gorm:"autoCreateTime"`
}
