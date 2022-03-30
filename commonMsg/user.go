package commonMsg

import (
	"github.com/coreservice-io/dns-common/data"
)

type UserLoginMsg struct {
	Email     string
	Password  string
	CaptchaId string
	Captcha   string
}

type UserResetPasswordMsg struct {
	Email    string
	Password string
	VCode    string
}

type EmailVCodeMsg struct {
	Email string
}

type QueryUserMsg struct {
	EmailPattern string
	UserId       uint
	Limit        int
	Offset       int
}

type AddUserMsg struct {
	Email       string
	Password    string
	Forbidden   bool
	Roles       []string
	Permissions []string
}

type UpdateUserMsg struct {
	Email       *string
	Forbidden   *bool
	Roles       *[]string
	Permissions *[]string
}

type RoleSettingResp struct {
	Roles       []string
	Permissions []string
}

type QueryUserResp struct {
	UserList []*data.User
	Count    int64
}

type CaptchaResp struct {
	Id      string
	Content string
}

type UserInfoResp struct {
	ID          uint
	Email       string
	Token       string
	Forbidden   bool
	Roles       []string
	Permissions []string
}
