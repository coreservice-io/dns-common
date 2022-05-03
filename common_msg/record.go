package common_msg

type Record struct {
	Id        int64
	Domain_id int64
	Name      string
	Type      string
	TTL       uint32
	Forbidden bool

	Updated int64
	Created int64
}

//add api msg
// @Description Msg_Req_AddRecord
type Msg_Req_AddRecord struct {
	Domain_id int64  //required
	Name      string //required
	Type      string //required
	TTL       uint32 //required
}

type Msg_Resp_AddRecord struct {
	API_META_STATUS
	Record *Record
}

//query
// @Description Msg_Req_QueryRecord_Filter
type Msg_Req_QueryRecord_Filter struct {
	Id           *int64    //optional
	Name         *[]string //optional //name list
	Name_pattern *string   //optional //query name pattern
	Type         *[]string //optional
	Domain_id    *int64    //optional
}

// @Description Msg_Req_QueryRecord
type Msg_Req_QueryRecord struct {
	Filter Msg_Req_QueryRecord_Filter //required
	Limit  int                        //required
	Offset int                        //required
}

type Msg_Resp_QueryRecord struct {
	API_META_STATUS
	Records []*Record
	Count   int64
}

//update
// @Description Msg_Req_UpdateRecord_Filter
type Msg_Req_UpdateRecord_Filter struct {
	Id []int64 //required
}

// @Description Msg_Req_UpdateRecord_To
type Msg_Req_UpdateRecord_To struct {
	TTL       *uint32 //optional
	Forbidden *bool   //optional
}

// @Description Msg_Req_UpdateRecord
type Msg_Req_UpdateRecord struct {
	Filter Msg_Req_UpdateRecord_Filter //required
	Update Msg_Req_UpdateRecord_To     //required
}

//using API_META_STATUS as update response

//delete
// @Description Msg_Req_DeleteRecord_Filter
type Msg_Req_DeleteRecord_Filter struct {
	Id []int64 //required
}

// @Description Msg_Req_DeleteRecord
type Msg_Req_DeleteRecord struct {
	Filter Msg_Req_DeleteRecord_Filter //required
}

//using API_META_STATUS as delete response
