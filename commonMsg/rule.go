package commonMsg

type AddRuleMsg struct {
	RecordId      uint
	SysVersion    int
	ContinentCode string
	CountryCode   string
	StartTime     string //15:08:08
	EndTime       string //15:08:08
	Destination   string
	Weight        int
}

type UpdateRuleMsg struct {
	ContinentCode *string
	CountryCode   *string
	StartTime     *string //15:08:08
	EndTime       *string //15:08:08
	Destination   *string
	Weight        *int
}

type AddRuleByRecordNameMsg struct {
	DomainName    string
	RecordName    string
	RecordType    string
	SysVersion    int
	ContinentCode string
	CountryCode   string
	StartTime     string //15:08:08
	EndTime       string //15:08:08
	Destination   string
	Weight        int
}

type QueryRulesByRecordNameMsg struct {
	DomainName string
	RecordName string
	RecordType string
}
