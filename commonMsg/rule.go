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
// @Description Msg_Req_AddRule
type Msg_Req_AddRule struct {
	Record_id      uint   //required
	Sys_version    int    //required
	Continent_code string //required
	Country_code   string //required
	Start_time     string //required //15:08:08
	End_time       string //required //15:08:08
	Destination    string //required
	Weight         int    //required
}

type Msg_Resp_AddRule struct {
	API_META_STATUS
	Rule *Rule
}

//query
// @Description Msg_Req_QueryRule_Filter
type Msg_Req_QueryRule_Filter struct {
	Record_id uint //required
}

// @Description Msg_Req_QueryRule
type Msg_Req_QueryRule struct {
	Filter Msg_Req_QueryRule_Filter //required
}

type Msg_Resp_QueryRules struct {
	API_META_STATUS
	Rules []*Rule
}

//update
// @Description Msg_Req_UpdateRule_Filter
type Msg_Req_UpdateRule_Filter struct {
	Id []uint //required
}

// @Description Msg_Req_UpdateRule_To
type Msg_Req_UpdateRule_To struct {
	Continent_code *string //optional
	Country_code   *string //optional
	Start_time     *string //optional //15:08:08
	End_time       *string //optional //15:08:08
	Destination    *string //optional
	Weight         *int    //optional
}

// @Description Msg_Req_UpdateRule
type Msg_Req_UpdateRule struct {
	Filter Msg_Req_UpdateRule_Filter //required
	Update Msg_Req_UpdateRule_To     //required
}

//using API_META_STATUS as update response

//delete
// @Description Msg_Req_DeleteRule_Filter
type Msg_Req_DeleteRule_Filter struct {
	Id []uint //required
}

// @Description Msg_Req_DeleteRule
type Msg_Req_DeleteRule struct {
	Filter Msg_Req_DeleteRule_Filter //required
}

//using API_META_STATUS as delete response
