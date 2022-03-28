package commonMsg

import (
	"github.com/coreservice-io/dns-common/model"
)

type AddRecordMsg struct {
	DomainId uint
	Name     string
	Type     string //enum
	TTL      uint32
}

type UpdateRecordMsg struct {
	Name      *string
	TTL       *uint32
	Forbidden *bool
}

type QueryRecordMsg struct {
	DomainId    uint
	NamePattern string
	Limit       int
	Offset      int
}

type QueryRecordListMsg struct {
	DomainId       uint
	RecordNameList []string
}

type QueryRecordResp struct {
	Records []*model.Record
	Count   int64
}
