package model

type User struct {
	ID          uint   `gorm:"primarykey"`
	Email       string `gorm:"index;unique;type:varchar(200)"`
	Password    string `gorm:"type:varchar(200)"`
	Token       string `gorm:"index;unique;type:varchar(24)"`
	Forbidden   bool
	Roles       StringArray
	Permissions StringArray

	Updated int64 `gorm:"autoUpdateTime"`
	Created int64 `gorm:"autoCreateTime"`
}
