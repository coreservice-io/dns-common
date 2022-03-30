package commonMsg

import (
	"github.com/coreservice-io/dns-common/data"
)

type AddDomainMsg struct {
	Domain         string
	ExpirationTime int64
}

type UpdateDomainMsg struct {
	Forbidden      *bool
	ExpirationTime *int64
}

type QueryDomainMsg struct {
	DomainPattern string
	UserId        uint
	Limit         int
	Offset        int
}

type QueryDomainByNameMsg struct {
	Domain string
}

type QueryDomainResp struct {
	DomainList []*data.Domain
	Count      int64
}
