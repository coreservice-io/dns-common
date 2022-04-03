package commonMsg

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

//add
type Msg_Req_AddRule struct {
	Record_id      int
	Sys_version    int
	Continent_code string
	Country_code   string
	Start_time     string //15:08:08
	End_time       string //15:08:08
	Destination    string
	Weight         int
}

type Msg_Resp_AddRule struct {
	API_META_STATUS
	Rule *Rule
}

//query
type Msg_Req_QueryRule_Filter struct {
	Record_id uint
}

type Msg_Req_QueryRule struct {
	Filter Msg_Req_QueryRule_Filter
}

type Msg_Resp_QueryRules struct {
	API_META_STATUS
	Rules []*Rule
}

//update
type Msg_Req_UpdateRule_Filter struct {
	Id []uint
}

type Msg_Req_UpdateRule_To struct {
	Continent_code *string
	Country_code   *string
	Start_time     *string //15:08:08
	End_time       *string //15:08:08
	Destination    *string
	Weight         *int
}

type Msg_Req_UpdateRule struct {
	Filter Msg_Req_UpdateRule_Filter
	Update Msg_Req_UpdateRule_To
}

//using API_META_STATUS as update response

//delete
type Msg_Req_DeleteRule_Filter struct {
	Id []uint
}

type Msg_Req_DeleteRule struct {
	Filter Msg_Req_DeleteRule_Filter
}

//using API_META_STATUS as delete response
