package data

type Rule struct {
	ID            uint
	SysVersion    int
	RecordId      uint
	ContinentCode string
	CountryCode   string
	StartTimeSecs int64 //range 0-86399
	EndTimeSecs   int64 //range 0-86399 End>Start
	Destination   string
	Weight        int

	Updated int64
	Created int64
}
