package common_msg

type Record struct {
	Id        int64  `json:"id"`
	Domain_id int64  `json:"domain_id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Forbidden bool   `json:"forbidden"`
	TTL       uint32 `json:"ttl"`

	Updated_unixtime int64 `gorm:"autoUpdateTime" json:"updated_unixtime"`
	Created_unixtime int64 `gorm:"autoCreateTime" json:"created_unixtime"`
}

//add api msg
// @Description Msg_Req_AddRecord
type Msg_Req_AddRecord struct {
	Domain_id int64  `json:"domain_id"` //required
	Name      string `json:"name"`      //required
	Type      string `json:"type"`      //required
	TTL       uint32 `json:"ttl"`       //required
}

type Msg_Resp_AddRecord struct {
	API_META_STATUS
	Record *Record `json:"record"`
}

//query
// @Description Msg_Req_QueryRecord_Filter
type Msg_Req_QueryRecord_Filter struct {
	Id           *int64    `json:"id"`           //optional
	Name         *[]string `json:"name"`         //optional //name list
	Name_pattern *string   `json:"name_pattern"` //optional //query name pattern
	Type         *[]string `json:"type"`         //optional
	Domain_id    *int64    `json:"domain_id"`    //optional
}

// @Description Msg_Req_QueryRecord
type Msg_Req_QueryRecord struct {
	Filter Msg_Req_QueryRecord_Filter `json:"filter"` //required
	Limit  int                        `json:"limit"`  //required
	Offset int                        `json:"offset"` //required
}

type Msg_Resp_QueryRecord struct {
	API_META_STATUS
	Records []*Record `json:"records"`
	Count   int64     `json:"count"`
}

//update
// @Description Msg_Req_UpdateRecord_Filter
type Msg_Req_UpdateRecord_Filter struct {
	Id []int64 `json:"id"` //required
}

// @Description Msg_Req_UpdateRecord_To
type Msg_Req_UpdateRecord_To struct {
	TTL       *uint32 `json:"ttl"`       //optional
	Forbidden *bool   `json:"forbidden"` //optional
}

// @Description Msg_Req_UpdateRecord
type Msg_Req_UpdateRecord struct {
	Filter Msg_Req_UpdateRecord_Filter `json:"filter"` //required
	Update Msg_Req_UpdateRecord_To     `json:"update"` //required
}

//using API_META_STATUS as update response

//delete
// @Description Msg_Req_DeleteRecord_Filter
type Msg_Req_DeleteRecord_Filter struct {
	Id []int64 `json:"id"` //required
}

// @Description Msg_Req_DeleteRecord
type Msg_Req_DeleteRecord struct {
	Filter Msg_Req_DeleteRecord_Filter `json:"filter"` //required
}

//using API_META_STATUS as delete response
