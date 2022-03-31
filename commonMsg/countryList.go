package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type Msg_Resp_CountryList struct {
	api.API_META_STATUS
	CountryList map[string]CountryData
}

type CountryData struct {
	CountryCode   string
	CountryName   string
	ContinentCode string
	ContinentName string
}
