package commonMsg

import (
	"github.com/coreservice-io/dns-common/model"
)

type AddCertMsg struct {
	Name        string
	CertContent string
	KeyContent  string
}

type UploadCertMsg struct {
	CertContent string
	KeyContent  string
}

type QueryCertMsg struct {
	NamePattern string
	UserId      uint
	Limit       int
	Offset      int
}

//custom apply cert
type CustomDomainCertMsg struct {
	ApplyDomain  string //ex. example.customdomain.com  (cname to pz1.mesoncdn.com)
	PullZoneName string //ex. pz1
	HostedDomain string //ex. mesoncdn.com
}

type QueryCertResp struct {
	CertList []*model.Cert
	Count    int64
}

type CertContentResp struct {
	CertContent string
	KeyContent  string
}
