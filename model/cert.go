package model

type Cert struct {
	ID             uint   `gorm:"primarykey"`
	Name           string `gorm:"index;unique;type:varchar(200)"`
	UserId         uint   `gorm:"index"`
	CertContent    string
	KeyContent     string
	ExpirationTime int64 `gorm:"index"`
	Hash           string
	RelatedDomain  StringArray

	Updated int64 `gorm:"autoUpdateTime"`
	Created int64 `gorm:"autoCreateTime"`
}
