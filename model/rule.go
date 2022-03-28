package model

type Rule struct {
	ID            uint `gorm:"primarykey"`
	SysVersion    int  `gorm:"index"`
	RecordId      uint `gorm:"index"`
	ContinentCode string
	CountryCode   string
	StartTimeSecs int64 //range 0-86399 //sqldb.DaySecondTime
	EndTimeSecs   int64 //range 0-86399 End>Start //sqldb.DaySecondTime
	Destination   string
	Weight        int

	Updated int64 `gorm:"autoUpdateTime"`
	Created int64 `gorm:"autoCreateTime"`
}
