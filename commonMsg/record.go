package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type Record struct {
	Id        uint
	Domain_id uint
	Name      string
	Type      string
	TTL       uint32
	Forbidden bool

	Updated int64
	Created int64
}

type Msg_Req_Record_Filter struct {
	Id           *uint
	Name         *[]string //name list
	Name_pattern *string   //query name pattern
	Type         *string
	Domain_id    *uint
}

//add api msg
type Msg_Req_AddRecord_Domain_Filter struct {
	Domain_id   *uint
	Domain_name *string
}

type Msg_Req_AddRecord struct {
	//Filter use domain_id in api path
	Name string
	Type string
	TTL  uint32
}

type Msg_Resp_AddRecord struct {
	api.API_META_STATUS
	Record *Record
}

//update
type Msg_Req_UpdateRecord_To struct {
	TTL       *uint32
	Forbidden *bool
}

type Msg_Req_UpdateRecord struct {
	//Filter use id in api path
	Update Msg_Req_UpdateRecord_To
}

//delete
//delete will use Msg_Req_Record_Filter as msg

//query
type Msg_Req_QueryRecord struct {
	Filter Msg_Req_Record_Filter
	Limit  int
	Offset int
}

type Msg_Resp_QueryRecord struct {
	api.API_META_STATUS
	Records []*Record
	Count   int64
}
