package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type Cert struct {
	Id              uint
	User_id         uint
	Related_domain  []string
	Hash            string
	Cert_content    string
	Key_content     string
	Expiration_time int64

	Updated int64
	Created int64
}

//filter
type Msg_Req_Cert_Filter struct {
	Id             *uint
	User_id        *uint
	Related_domain *string
	Hash           *string
}

//add api msg
type Msg_Req_AddCert struct {
	Related_domain []string
	Cert_content   string
	Key_content    string
}

type Msg_Resp_AddCert struct {
	api.API_META_STATUS
	Cert *Cert
}

//update api msg
type Msg_Req_UpdateCert_To struct {
	Cert_content string
	Key_content  string
}

type Msg_Req_UploadCert struct {
	Filter Msg_Req_Cert_Filter
	Update Msg_Req_UpdateCert_To
}

//delete
//delete will use Msg_Req_Cert_Filter as msg

//query api msg
type Msg_Req_QueryCert struct {
	Filter Msg_Req_Cert_Filter
	Limit  int
	Offset int
}

type Msg_Resp_QueryCert struct {
	api.API_META_STATUS
	Cert_list []*Cert
	Count     int64
}

//custom apply cert
type Msg_Req_ApplyCustomCert struct {
	Apply_domain  string //ex. example.customdomain.com  (cname to pz1.mesoncdn.com)
	Txt_name_tag  string //ex. pz1
	Hosted_domain string //ex. mesoncdn.com
}
