package commonMsg

type Domain struct {
	Id              int64
	Name            string
	User_id         int64
	Forbidden       bool
	Expiration_time int64

	Updated int64
	Created int64
}

//add api msg
// @Description Msg_Req_AddDomain
type Msg_Req_AddDomain struct {
	Name            string //required
	Expiration_time int64  //required
}

type Msg_Resp_AddDomain struct {
	API_META_STATUS
	Domain *Domain
}

//query api msg
// @Description Msg_Req_QueryDomain_Filter
type Msg_Req_QueryDomain_Filter struct {
	Id           *int64  //optional
	Name_pattern *string //optional
	Name         *string //optional
	User_id      *int64  //optional
}

// @Description Msg_Req_QueryDomain
type Msg_Req_QueryDomain struct {
	Filter Msg_Req_QueryDomain_Filter //required
	Limit  int64                      //required
	Offset int64                      //required
}

type Msg_Resp_QueryDomain struct {
	API_META_STATUS
	Domain_list []*Domain
	Count       int64
}

//update api msg
// @Description Msg_Req_UpdateDomain_Filter
type Msg_Req_UpdateDomain_Filter struct {
	Id []int64 //required
}

// @Description Msg_Req_UpdateDomain_To
type Msg_Req_UpdateDomain_To struct {
	Forbidden       *bool  //optional
	Expiration_time *int64 //optional
}

// @Description Msg_Req_UpdateDomain
type Msg_Req_UpdateDomain struct {
	Filter Msg_Req_UpdateDomain_Filter //required
	Update Msg_Req_UpdateDomain_To     //required
}

//using API_META_STATUS as update response

//delete
// @Description Msg_Req_DeleteDomain_Filter
type Msg_Req_DeleteDomain_Filter struct {
	Id []int64 //required
}

// @Description Msg_Req_DeleteDomain
type Msg_Req_DeleteDomain struct {
	Filter Msg_Req_DeleteDomain_Filter //required
}

//using API_META_STATUS as delete response
