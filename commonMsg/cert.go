package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type Msg_Req_AddCert struct {
	Domains      []string
	Cert_content string
	Key_content  string
}

type Msg_Req_UploadCert struct {
	Cert_content string
	Key_content  string
}

type Msg_Req_QueryCert struct {
	Name_pattern string
	User_id      uint
	Limit        int
	Offset       int
}

//custom apply cert
type Msg_Req_ApplyCustomCert struct {
	Apply_domain   string //ex. example.customdomain.com  (cname to pz1.mesoncdn.com)
	Pull_zone_name string //ex. pz1
	Hosted_domain  string //ex. mesoncdn.com
}

type Msg_Resp_CertInfo struct {
	api.API_META_STATUS
	Cert Cert
}

type Msg_Resp_QueryCert struct {
	api.API_META_STATUS
	Cert_list []*Cert
	Count     int64
}

type Msg_Resp_CertContent struct {
	api.API_META_STATUS
	Cert_content string
	Key_content  string
}

type Msg_Resp_CertHash struct {
	api.API_META_STATUS
	Hash string
}

type Cert struct {
	ID              uint
	Name            string
	User_id         uint
	Cert_content    string
	Key_content     string
	Expiration_time int64
	Hash            string
	Related_domain  []string

	Updated int64
	Created int64
}
