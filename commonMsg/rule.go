package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type Msg_Req_AddRule struct {
	RecordId      uint
	SysVersion    int
	ContinentCode string
	CountryCode   string
	StartTime     string //15:08:08
	EndTime       string //15:08:08
	Destination   string
	Weight        int
}

type Msg_Req_UpdateRule struct {
	ContinentCode *string
	CountryCode   *string
	StartTime     *string //15:08:08
	EndTime       *string //15:08:08
	Destination   *string
	Weight        *int
}

type Msg_Req_AddRuleByRecordName struct {
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

type Msg_Req_QueryRulesByRecordName struct {
	DomainName string
	RecordName string
	RecordType string
}

type Msg_Resp_RuleInfo struct {
	api.API_META_STATUS
	Rule Rule
}

type Msg_Resp_Rules struct {
	api.API_META_STATUS
	Rules []*Rule
}

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
