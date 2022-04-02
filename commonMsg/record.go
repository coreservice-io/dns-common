package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type Msg_Req_AddRecord struct {
	Domain_id uint
	Name      string
	Type      string //enum
	TTL       uint32
}

type Msg_Req_AddRecordByDomainName struct {
	Domain_name string
	Name        string
	Type        string //enum
	TTL         uint32
}

type Msg_Req_UpdateRecord struct {
	TTL       *uint32
	Forbidden *bool
}

type Msg_Req_QueryRecord struct {
	Domain_id    uint
	Name_pattern string
	Record_id    uint
	Record_type  string
	Limit        int
	Offset       int
}

type Msg_Req_QueryRecordByDomainName struct {
	Domain_name  string
	Name_pattern string
	Record_id    uint
	Record_type  string
	Limit        int
	Offset       int
}

type Msg_Req_QueryRecordByGivenName struct {
	Domain_name      string
	Record_name_list []string
	Record_type      string
}

type Msg_Req_DeleteRecordByName struct {
	Domain_name string
	Record_name string
	Record_type string
}

type Msg_Req_UpdateRecordByName struct {
	Domain_name string
	Record_name string
	Record_type string
	TTL         *uint32
	Forbidden   *bool
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
	Domain_id uint
	Name      string
	Type      string //enum
	Forbidden bool
	TTL       uint32

	Updated int64
	Created int64
}
