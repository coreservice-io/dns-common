package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type Msg_Req_AddDomain struct {
	Domain         string
	ExpirationTime int64
}

type Msg_Req_UpdateDomain struct {
	Forbidden      *bool
	ExpirationTime *int64
}

type Msg_Req_QueryDomain struct {
	DomainPattern string
	DomainId      uint
	UserId        uint
	Limit         int
	Offset        int
}

type Msg_Resp_DomainInfo struct {
	api.API_META_STATUS
	Domain Domain
}

type Msg_Resp_QueryDomain struct {
	api.API_META_STATUS
	DomainList []*Domain
	Count      int64
}

type Domain struct {
	ID             uint
	Name           string
	ExpirationTime int64
	Forbidden      bool
	UserId         uint

	Updated int64
	Created int64
}
