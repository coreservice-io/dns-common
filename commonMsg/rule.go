package commonMsg

type AddRuleMsg struct {
	SysVersion    int
	RecordId      uint
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
