package commonMsg

type Rule struct {
	Id             int
	Sys_version    int
	Record_id      int
	Continent_code string
	Country_code   string
	Destination    string
	Weight         int

	Updated int64
	Created int64
}

//add
// @Description Msg_Req_AddRule
type Msg_Req_AddRule struct {
	Record_id      int    //required
	Sys_version    int    //required
	Continent_code string //required
	Country_code   string //required
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
	Record_id int //required
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
	Id []int //required
}

// @Description Msg_Req_UpdateRule_To
type Msg_Req_UpdateRule_To struct {
	Continent_code *string //optional
	Country_code   *string //optional
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
	Id []int //required
}

// @Description Msg_Req_DeleteRule
type Msg_Req_DeleteRule struct {
	Filter Msg_Req_DeleteRule_Filter //required
}

//using API_META_STATUS as delete response
