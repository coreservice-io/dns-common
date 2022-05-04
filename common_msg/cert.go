package common_msg

type Cert struct {
	Id              int64  `json:"id"`
	User_id         int64  `json:"user_id"`
	Cert_content    string `json:"cert_content"`
	Key_content     string `json:"key_content"`
	Expiration_time int64  `json:"expiration_time"`
	Hash            string `json:"hash"`
	Related_domain  string `json:"related_domain"`

	Updated_unixtime int64 `gorm:"autoUpdateTime" json:"updated_unixtime"`
	Created_unixtime int64 `gorm:"autoCreateTime" json:"created_unixtime"`
}

//add api msg
// @Description Msg_Req_AddCert
type Msg_Req_AddCert struct {
	Related_domain *[]string `json:"related_domain"` //optional
	Cert_content   *string   `json:"cert_content"`   //optional
	Key_content    *string   `json:"key_content"`    //optional
}

type Msg_Resp_AddCert struct {
	API_META_STATUS
	Cert *Cert `json:"cert"`
}

//query api msg
// @Description Msg_Req_QueryCert_Filter
type Msg_Req_QueryCert_Filter struct {
	Id                     *int64  `json:"id"`                     //optional
	User_id                *int64  `json:"user_id"`                //optional
	Related_domain_pattern *string `json:"related_domain_pattern"` //optional
}

// @Description Msg_Req_QueryCert
type Msg_Req_QueryCert struct {
	Filter Msg_Req_QueryCert_Filter `json:"filter"` //required
	Limit  int                      `json:"limit"`  //required
	Offset int                      `json:"offset"` //required
}

type Msg_Resp_QueryCert struct {
	API_META_STATUS
	Cert_list []*Cert `json:"cert_list"`
	Count     int64   `json:"count"`
}

//update api msg
// @Description Msg_Req_UpdateCert_Filter
type Msg_Req_UpdateCert_Filter struct {
	Id []int64 `json:"id"` //required
}

// @Description Msg_Req_UpdateCert_To
type Msg_Req_UpdateCert_To struct {
	Cert_content string `json:"cert_content"` //required
	Key_content  string `json:"key_content"`  //required
}

// @Description Msg_Req_UpdateCert
type Msg_Req_UpdateCert struct {
	Filter Msg_Req_UpdateCert_Filter `json:"filter"` //required
	Update Msg_Req_UpdateCert_To     `json:"update"` //required
}

//using API_META_STATUS as update response

//delete
// @Description Msg_Req_DeleteCert_Filter
type Msg_Req_DeleteCert_Filter struct {
	Id []int64 `json:"id"` //required
}

// @Description Msg_Req_DeleteCert
type Msg_Req_DeleteCert struct {
	Filter Msg_Req_DeleteCert_Filter `json:"filter"` //required
}

//using API_META_STATUS as delete response

//custom apply cert

// @Description Msg_Req_ApplyCustomCert
type Msg_Req_ApplyCustomCert struct {
	Apply_domain  string `json:"apply_domain"`  //required //ex. example.customdomain.com  (cname to pz1.mesoncdn.com)
	Txt_name_tag  string `json:"txt_name_tag"`  //required //ex. pz1
	Hosted_domain string `json:"hosted_domain"` //required //ex. mesoncdn.com
}

type Msg_Resp_ApplyCustomCert struct {
	API_META_STATUS
	Cert_content string `json:"cert_content"`
	Key_content  string `json:"key_content"`
}
