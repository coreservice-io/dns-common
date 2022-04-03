package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type User struct {
	Id          uint
	Email       string
	Password    string
	Token       string
	Forbidden   bool
	Roles       []string
	Permissions []string

	Updated int64
	Created int64
}

type Msg_Req_User_Filter struct {
	Id    *int
	Email *string
}

//query
type Msg_Req_QueryUser struct {
	Filter Msg_Req_User_Filter
	Limit  int
	Offset int
}

type Msg_Resp_QueryUser struct {
	api.API_META_STATUS
	User_list []*User
	Count     int64
}

//add
type Msg_Req_AddUser struct {
	Email       string
	Password    string
	Forbidden   bool
	Roles       []string
	Permissions []string
}

//resp for both 'add' and 'login' api
type Msg_Resp_UserInfo struct {
	api.API_META_STATUS
	User User
}

//update
type Msg_Req_UpdateUser_To struct {
	Forbidden   *bool
	Roles       *[]string
	Permissions *[]string
}

type Msg_Req_UpdateUser struct {
	Filter Msg_Req_User_Filter
	Update Msg_Req_UpdateUser_To
}

//login
type Msg_Req_UserLogin struct {
	Email      string
	Password   string
	Captcha_id string
	Captcha    string
}

//reset pass
type Msg_Req_UserResetPassword struct {
	Email    string
	Password string
	VCode    string
}

//req evcode
type Msg_Req_EmailVCode struct {
	Email string
}

//captcha
type Msg_Resp_Captcha struct {
	api.API_META_STATUS
	Id      string
	Content string
}
