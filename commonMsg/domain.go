package commonMsg

type Domain struct {
	Id              uint
	Name            string
	User_id         uint
	Forbidden       bool
	Expiration_time int64

	Updated int64
	Created int64
}

//add api msg
type Msg_Req_AddDomain struct {
	Name            string
	Expiration_time int64
}

type Msg_Resp_AddDomain struct {
	API_META_STATUS
	Domain *Domain
}

//query api msg
type Msg_Req_QueryDomain_Filter struct {
	Id           *uint
	Name_pattern *string
	User_id      *uint
}

type Msg_Req_QueryDomain struct {
	Filter Msg_Req_QueryDomain_Filter
	Limit  int
	Offset int
}

type Msg_Resp_QueryDomain struct {
	API_META_STATUS
	Domain_list []*Domain
	Count       int64
}

//update api msg
type Msg_Req_UpdateDomain_Filter struct {
	Id []uint
}

type Msg_Req_UpdateDomain_To struct {
	Forbidden       *bool
	Expiration_time *int64
}

type Msg_Req_UpdateDomain struct {
	Filter Msg_Req_UpdateDomain_Filter
	Update Msg_Req_UpdateDomain_To
}

//using API_META_STATUS as update response

//delete
type Msg_Req_DeleteDomain_Filter struct {
	Id []uint
}

type Msg_Req_DeleteDomain struct {
	Filter Msg_Req_DeleteDomain_Filter
}

//using API_META_STATUS as delete response
