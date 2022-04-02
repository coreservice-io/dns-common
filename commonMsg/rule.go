package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type Msg_Req_AddRule struct {
	Record_id      uint
	Sys_version    int
	Continent_code string
	Country_code   string
	Start_time     string //15:08:08
	End_time       string //15:08:08
	Destination    string
	Weight         int
}

type Msg_Req_UpdateRule struct {
	Continent_code *string
	Country_code   *string
	Start_time     *string //15:08:08
	End_time       *string //15:08:08
	Destination    *string
	Weight         *int
}

type Msg_Req_AddRuleByRecordName struct {
	Domain_name    string
	Record_name    string
	Record_type    string
	Sys_version    int
	Continent_code string
	Country_code   string
	Start_time     string //15:08:08
	End_time       string //15:08:08
	Destination    string
	Weight         int
}

type Msg_Req_QueryRulesByRecordName struct {
	Domain_name string
	Record_name string
	Record_type string
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
	ID              uint
	Sys_version     int
	Record_id       uint
	Continent_code  string
	Country_code    string
	Start_time_secs int64 //range 0-86399
	End_time_secs   int64 //range 0-86399 End>Start
	Destination     string
	Weight          int

	Updated int64
	Created int64
}
