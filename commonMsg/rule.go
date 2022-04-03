package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type Rule struct {
	Id              uint
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

//filter
type Msg_Req_Rule_Filter struct {
	Record_id   *uint
	Domain_name *string
	Record_name *string
	Record_type *string
}

//add
type Msg_Req_AddRule struct {
	Filter         Msg_Req_Rule_Filter
	Sys_version    int
	Continent_code string
	Country_code   string
	Start_time     string //15:08:08
	End_time       string //15:08:08
	Destination    string
	Weight         int
}

type Msg_Resp_AddRule struct {
	api.API_META_STATUS
	Rule *Rule
}

//delete
//delete will use Msg_Req_Rule_Filter as msg

//query
type Msg_Req_QueryRules struct {
	Filter Msg_Req_Rule_Filter
}

type Msg_Resp_QueryRules struct {
	api.API_META_STATUS
	Rules []*Rule
}

//update
type Msg_Req_UpdateRule_To struct {
	Continent_code *string
	Country_code   *string
	Start_time     *string //15:08:08
	End_time       *string //15:08:08
	Destination    *string
	Weight         *int
}

type Msg_Req_UpdateRule struct {
	Filter Msg_Req_Rule_Filter
	Update Msg_Req_UpdateRule_To
}
