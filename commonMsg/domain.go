package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type Domain struct {
	Id              uint
	Name            string
	User_id         uint
	Forbidden       bool
	Expiration_time int64

	Updated int64
	Created int64
}

//filter
type Msg_Req_Domain_Filter struct {
	Id      *uint
	Name    *string
	User_id *uint
}

//add api msg
type Msg_Req_AddDomain struct {
	Name            string
	Expiration_time int64
}

type Msg_Resp_AddDomain struct {
	api.API_META_STATUS
	Domain *Domain
}

//update api msg
type Msg_Req_UpdateDomain_To struct {
	Forbidden       *bool
	Expiration_time *int64
}

type Msg_Req_UpdateDomain struct {
	//Filter use id in path
	Update Msg_Req_UpdateDomain_To
}

//delete
//delete will use Msg_Req_Domain_Filter as msg

//query api msg
type Msg_Req_QueryDomain struct {
	Filter Msg_Req_Domain_Filter
	Limit  int
	Offset int
}

type Msg_Resp_QueryDomain struct {
	api.API_META_STATUS
	Domain_list []*Domain
	Count       int64
}
