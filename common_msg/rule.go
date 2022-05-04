package common_msg

type Rule struct {
	Id             int64  `json:"id"`
	Record_id      int64  `json:"record_id"`
	Continent_code string `json:"continent_code"`
	Country_code   string `json:"country_code"`
	Destination    string `json:"destination"`
	Weight         int64  `json:"weight"`

	Updated_unixtime int64 `gorm:"autoUpdateTime" json:"updated_unixtime"`
	Created_unixtime int64 `gorm:"autoCreateTime" json:"created_unixtime"`
}

//add
// @Description Msg_Req_AddRule
type Msg_Req_AddRule struct {
	Record_id      int64  `json:"record_id"`      //required
	Continent_code string `json:"continent_code"` //required
	Country_code   string `json:"country_code"`   //required
	Destination    string `json:"destination"`    //required
	Weight         int64  `json:"weight"`         //required
}

type Msg_Resp_AddRule struct {
	API_META_STATUS
	Rule *Rule `json:"rule"`
}

//query
// @Description Msg_Req_QueryRule_Filter
type Msg_Req_QueryRule_Filter struct {
	Record_id int64 `json:"record_id"` //required
}

// @Description Msg_Req_QueryRule
type Msg_Req_QueryRule struct {
	Filter Msg_Req_QueryRule_Filter `json:"filter"` //required
}

type Msg_Resp_QueryRules struct {
	API_META_STATUS
	Rules []*Rule `json:"rules"`
}

//update
// @Description Msg_Req_UpdateRule_Filter
type Msg_Req_UpdateRule_Filter struct {
	Id []int64 `json:"id"` //required
}

// @Description Msg_Req_UpdateRule_To
type Msg_Req_UpdateRule_To struct {
	Continent_code *string `json:"continent_code"` //optional
	Country_code   *string `json:"country_code"`   //optional
	Destination    *string `json:"destination"`    //optional
	Weight         *int64  `json:"weight"`         //optional
}

// @Description Msg_Req_UpdateRule
type Msg_Req_UpdateRule struct {
	Filter Msg_Req_UpdateRule_Filter `json:"filter"` //required
	Update Msg_Req_UpdateRule_To     `json:"update"` //required
}

//using API_META_STATUS as update response

//delete
// @Description Msg_Req_DeleteRule_Filter
type Msg_Req_DeleteRule_Filter struct {
	Id []int64 `json:"id"` //required
}

// @Description Msg_Req_DeleteRule
type Msg_Req_DeleteRule struct {
	Filter Msg_Req_DeleteRule_Filter `json:"filter"` //required
}

//using API_META_STATUS as delete response
