package commonMsg

type Record struct {
	Id        uint
	Domain_id uint
	Name      string
	Type      string
	TTL       uint32
	Forbidden bool

	Updated int64
	Created int64
}

//add api msg
type Msg_Req_AddRecord struct {
	Domain_id uint
	Name      string
	Type      string
	TTL       uint32
}

type Msg_Resp_AddRecord struct {
	API_META_STATUS
	Record *Record
}

//query
type Msg_Req_QueryRecord_Filter struct {
	Id           *uint
	Name         *[]string //name list
	Name_pattern *string   //query name pattern
	Type         *string
	Domain_id    *uint
}

type Msg_Req_QueryRecord struct {
	Filter Msg_Req_QueryRecord_Filter
	Limit  int
	Offset int
}

type Msg_Resp_QueryRecord struct {
	API_META_STATUS
	Records []*Record
	Count   int64
}

//update
type Msg_Req_UpdateRecord_Filter struct {
	Id []uint
}

type Msg_Req_UpdateRecord_To struct {
	TTL       *uint32
	Forbidden *bool
}

type Msg_Req_UpdateRecord struct {
	Filter Msg_Req_UpdateRecord_Filter
	Update Msg_Req_UpdateRecord_To
}

//using API_META_STATUS as update response

//delete
type Msg_Req_DeleteRecord_Filter struct {
	Id []uint
}

type Msg_Req_DeleteRecord struct {
	Filter Msg_Req_DeleteRecord_Filter
}

//using API_META_STATUS as delete response
