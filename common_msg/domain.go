package common_msg

type Domain struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	Expiration_time int64  `json:"expiration_time"`
	Forbidden       bool   `json:"forbidden"`
	User_id         int64  `json:"user_id"`

	Updated_unixtime int64 `gorm:"autoUpdateTime" json:"updated_unixtime"`
	Created_unixtime int64 `gorm:"autoCreateTime" json:"created_unixtime"`
}

//add api msg
// @Description Msg_Req_AddDomain
type Msg_Req_AddDomain struct {
	Name            string `json:"name"`            //required
	Expiration_time int64  `json:"expiration_time"` //required
}

type Msg_Resp_AddDomain struct {
	API_META_STATUS
	Domain *Domain `json:"domain"`
}

//query api msg
// @Description Msg_Req_QueryDomain_Filter
type Msg_Req_QueryDomain_Filter struct {
	Id           *int64  `json:"id"`           //optional
	Name_pattern *string `json:"name_pattern"` //optional
	Name         *string `json:"name"`         //optional
	User_id      *int64  `json:"user_id"`      //optional
}

// @Description Msg_Req_QueryDomain
type Msg_Req_QueryDomain struct {
	Filter Msg_Req_QueryDomain_Filter `json:"filter"` //required
	Limit  int                        `json:"limit"`  //required
	Offset int                        `json:"offset"` //required
}

type Msg_Resp_QueryDomain struct {
	API_META_STATUS
	Domain_list []*Domain `json:"domain_list"`
	Count       int64     `json:"count"`
}

//update api msg
// @Description Msg_Req_UpdateDomain_Filter
type Msg_Req_UpdateDomain_Filter struct {
	Id []int64 `json:"id"` //required
}

// @Description Msg_Req_UpdateDomain_To
type Msg_Req_UpdateDomain_To struct {
	Forbidden       *bool  `json:"forbidden"`       //optional
	Expiration_time *int64 `json:"expiration_time"` //optional
}

// @Description Msg_Req_UpdateDomain
type Msg_Req_UpdateDomain struct {
	Filter Msg_Req_UpdateDomain_Filter `json:"filter"` //required
	Update Msg_Req_UpdateDomain_To     `json:"update"` //required
}

//using API_META_STATUS as update response

//delete
// @Description Msg_Req_DeleteDomain_Filter
type Msg_Req_DeleteDomain_Filter struct {
	Id []int64 `json:"id"` //required
}

// @Description Msg_Req_DeleteDomain
type Msg_Req_DeleteDomain struct {
	Filter Msg_Req_DeleteDomain_Filter `json:"filter"` //required
}

//using API_META_STATUS as delete response
