package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type Msg_Req_AddCert struct {
	Name        string
	CertContent string
	KeyContent  string
}

type Msg_Req_UploadCert struct {
	CertContent string
	KeyContent  string
}

type Msg_Req_QueryCert struct {
	NamePattern string
	UserId      uint
	Limit       int
	Offset      int
}

//custom apply cert
type Msg_Req_ApplyCustomCert struct {
	ApplyDomain  string //ex. example.customdomain.com  (cname to pz1.mesoncdn.com)
	PullZoneName string //ex. pz1
	HostedDomain string //ex. mesoncdn.com
}

type Msg_Resp_CertInfo struct {
	api.API_META_STATUS
	Cert Cert
}

type Msg_Resp_QueryCert struct {
	api.API_META_STATUS
	CertList []*Cert
	Count    int64
}

type Msg_Resp_CertContent struct {
	api.API_META_STATUS
	CertContent string
	KeyContent  string
}

type Msg_Resp_CertHash struct {
	api.API_META_STATUS
	Hash string
}

type Cert struct {
	ID             uint
	Name           string
	UserId         uint
	CertContent    string
	KeyContent     string
	ExpirationTime int64
	Hash           string
	RelatedDomain  []string

	Updated int64
	Created int64
}
