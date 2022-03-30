package commonMsg

import (
	"github.com/coreservice-io/dns-common/data"
)

type AddRecordMsg struct {
	DomainId uint
	Name     string
	Type     string //enum
	TTL      uint32
}

type AddRecordByDomainNameMsg struct {
	DomainName string
	Name       string
	Type       string //enum
	TTL        uint32
}

type UpdateRecordMsg struct {
	TTL       *uint32
	Forbidden *bool
}

type QueryRecordMsg struct {
	DomainId    uint
	NamePattern string
	Limit       int
	Offset      int
}

type QueryRecordByDomainNameMsg struct {
	DomainName  string
	NamePattern string
	Limit       int
	Offset      int
}

type QueryRecordListMsg struct {
	DomainName     string
	RecordNameList []string
}

type DeleteRecordByNameMsg struct {
	DomainName string
	RecordName string
	RecordType string
}

type UpdateRecordByNameMsg struct {
	DomainName string
	RecordName string
	TTL        *uint32
	Forbidden  *bool
}

type QueryRecordResp struct {
	Records []*data.Record
	Count   int64
}
