package model

type Record struct {
	ID        uint   `gorm:"primarykey"`
	DomainId  uint   `gorm:"index"`
	Name      string `gorm:"index"`
	Type      string //enum
	Forbidden bool
	TTL       uint32

	Updated int64 `gorm:"autoUpdateTime"`
	Created int64 `gorm:"autoCreateTime"`
}
