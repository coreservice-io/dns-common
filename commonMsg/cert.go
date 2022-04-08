package commonMsg

type Cert struct {
	Id              int64
	User_id         int64
	Related_domain  []string
	Hash            string
	Cert_content    string
	Key_content     string
	Expiration_time int64

	Updated int64
	Created int64
}

//add api msg
// @Description Msg_Req_AddCert
type Msg_Req_AddCert struct {
	Related_domain *[]string //optional
	Cert_content   *string   //optional
	Key_content    *string   //optional
}

type Msg_Resp_AddCert struct {
	API_META_STATUS
	Cert *Cert
}

//query api msg
// @Description Msg_Req_QueryCert_Filter
type Msg_Req_QueryCert_Filter struct {
	Id                     *int64  //optional
	User_id                *int64  //optional
	Related_domain_pattern *string //optional
}

// @Description Msg_Req_QueryCert
type Msg_Req_QueryCert struct {
	Filter Msg_Req_QueryCert_Filter //required
	Limit  int                      //required
	Offset int                      //required
}

type Msg_Resp_QueryCert struct {
	API_META_STATUS
	Cert_list []*Cert
	Count     int64
}

//update api msg
// @Description Msg_Req_UpdateCert_Filter
type Msg_Req_UpdateCert_Filter struct {
	Id []int64 //required
}

// @Description Msg_Req_UpdateCert_To
type Msg_Req_UpdateCert_To struct {
	Cert_content string //required
	Key_content  string //required
}

// @Description Msg_Req_UpdateCert
type Msg_Req_UpdateCert struct {
	Filter Msg_Req_UpdateCert_Filter //required
	Update Msg_Req_UpdateCert_To     //required
}

//using API_META_STATUS as update response

//delete
// @Description Msg_Req_DeleteCert_Filter
type Msg_Req_DeleteCert_Filter struct {
	Id []int64 //required
}

// @Description Msg_Req_DeleteCert
type Msg_Req_DeleteCert struct {
	Filter Msg_Req_DeleteCert_Filter //required
}

//using API_META_STATUS as delete response

//custom apply cert

// @Description Msg_Req_ApplyCustomCert
type Msg_Req_ApplyCustomCert struct {
	Apply_domain  string //required //ex. example.customdomain.com  (cname to pz1.mesoncdn.com)
	Txt_name_tag  string //required //ex. pz1
	Hosted_domain string //required //ex. mesoncdn.com
}

type Msg_Resp_ApplyCustomCert struct {
	API_META_STATUS
	Cert_content string
	Key_content  string
}
