package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type Msg_Req_AddRecord struct {
	DomainId uint
	Name     string
	Type     string //enum
	TTL      uint32
}

type Msg_Req_AddRecordByDomainName struct {
	DomainName string
	Name       string
	Type       string //enum
	TTL        uint32
}

type Msg_Req_UpdateRecord struct {
	TTL       *uint32
	Forbidden *bool
}

type Msg_Req_QueryRecord struct {
	DomainId    uint
	NamePattern string
	RecordId    uint
	RecordType  string
	Limit       int
	Offset      int
}

type Msg_Req_QueryRecordByDomainName struct {
	DomainName  string
	NamePattern string
	RecordId    uint
	RecordType  string
	Limit       int
	Offset      int
}

type Msg_Req_QueryRecordByGivenName struct {
	DomainName     string
	RecordNameList []string
	RecordType     string
}

type Msg_Req_DeleteRecordByName struct {
	DomainName string
	RecordName string
	RecordType string
}

type Msg_Req_UpdateRecordByName struct {
	DomainName string
	RecordName string
	RecordType string
	TTL        *uint32
	Forbidden  *bool
}

type Msg_Resp_RecordInfo struct {
	api.API_META_STATUS
	Record Record
}

type Msg_Resp_QueryRecordByGivenName struct {
	api.API_META_STATUS
	Records []*Record
}

type Msg_Resp_QueryRecord struct {
	api.API_META_STATUS
	Records []*Record
	Count   int64
}

type Record struct {
	ID        uint
	DomainId  uint
	Name      string
	Type      string //enum
	Forbidden bool
	TTL       uint32

	Updated int64
	Created int64
}
