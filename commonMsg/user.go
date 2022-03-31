package commonMsg

import (
	"github.com/coreservice-io/dns-common/tools/http/api"
)

type Msg_Req_UserLogin struct {
	Email     string
	Password  string
	CaptchaId string
	Captcha   string
}

type Msg_Req_UserResetPassword struct {
	Email    string
	Password string
	VCode    string
}

type Msg_Req_EmailVCode struct {
	Email string
}

type Msg_Req_QueryUser struct {
	EmailPattern string
	UserId       uint
	Limit        int
	Offset       int
}

type Msg_Req_AddUser struct {
	Email       string
	Password    string
	Forbidden   bool
	Roles       []string
	Permissions []string
}

type Msg_Req_UpdateUser struct {
	Email       *string
	Forbidden   *bool
	Roles       *[]string
	Permissions *[]string
}

type Msg_Resp_RoleSetting struct {
	api.API_META_STATUS
	Roles       []string
	Permissions []string
}

type Msg_Resp_QueryUser struct {
	api.API_META_STATUS
	UserList []*User
	Count    int64
}

type Msg_Resp_Captcha struct {
	api.API_META_STATUS
	Id      string
	Content string
}

type Msg_Resp_UserInfo struct {
	api.API_META_STATUS
	User User
}

type User struct {
	ID          uint
	Email       string
	Password    string
	Token       string
	Forbidden   bool
	Roles       []string
	Permissions []string

	Updated int64
	Created int64
}
