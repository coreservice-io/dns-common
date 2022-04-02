package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type Msg_Req_AddDomain struct {
	Domain          string
	Expiration_time int64
}

type Msg_Req_UpdateDomain struct {
	Forbidden       *bool
	Expiration_time *int64
}

type Msg_Req_QueryDomain struct {
	Domain_pattern string
	Domain_id      uint
	User_id        uint
	Limit          int
	Offset         int
}

type Msg_Resp_DomainInfo struct {
	api.API_META_STATUS
	Domain Domain
}

type Msg_Resp_QueryDomain struct {
	api.API_META_STATUS
	Domain_list []*Domain
	Count       int64
}

type Domain struct {
	ID              uint
	Name            string
	Expiration_time int64
	Forbidden       bool
	User_id         uint

	Updated int64
	Created int64
}
