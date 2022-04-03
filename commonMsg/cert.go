package commonMsg

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

//add api msg
type Msg_Req_AddCert struct {
	Related_domain []string
	Cert_content   string
	Key_content    string
}

type Msg_Resp_AddCert struct {
	API_META_STATUS
	Cert *Cert
}

//query api msg
type Msg_Req_QueryCert_Filter struct {
	Id             *uint
	User_id        *uint
	Related_domain *string
	Hash           *string
}

type Msg_Req_QueryCert struct {
	Filter Msg_Req_QueryCert_Filter
	Limit  int
	Offset int
}

type Msg_Resp_QueryCert struct {
	API_META_STATUS
	Cert_list []*Cert
	Count     int64
}

//update api msg
type Msg_Req_UpdateCert_Filter struct {
	Id []uint
}

type Msg_Req_UpdateCert_To struct {
	Cert_content string
	Key_content  string
}

type Msg_Req_UpdateCert struct {
	Filter Msg_Req_UpdateCert_Filter
	Update Msg_Req_UpdateCert_To
}

//using API_META_STATUS as update response

//delete
type Msg_Req_DeleteCert_Filter struct {
	Id []uint
}
type Msg_Req_DeleteCert struct {
	Filter Msg_Req_DeleteCert_Filter
}

//using API_META_STATUS as delete response

//custom apply cert
type Msg_Req_ApplyCustomCert struct {
	Apply_domain  string //ex. example.customdomain.com  (cname to pz1.mesoncdn.com)
	Txt_name_tag  string //ex. pz1
	Hosted_domain string //ex. mesoncdn.com
}
