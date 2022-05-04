package common_msg

type User struct {
	Id          int64    `json:"id"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	Token       string   `json:"token"`
	Forbidden   bool     `json:"forbidden"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`

	Updated_unixtime int64 `gorm:"autoUpdateTime" json:"updated_unixtime"`
	Created_unixtime int64 `gorm:"autoCreateTime" json:"created_unixtime"`
}

//add
// @Description Msg_Req_AddUser
type Msg_Req_AddUser struct {
	Email       string   `json:"email"`       //required
	Password    string   `json:"password"`    //required
	Forbidden   bool     `json:"forbidden"`   //required
	Roles       []string `json:"roles"`       //required
	Permissions []string `json:"permissions"` //required
}

//login
// @Description Msg_Req_UserLogin
type Msg_Req_UserLogin struct {
	Email      string `json:"email"`      //required
	Password   string `json:"password"`   //required
	Captcha_id string `json:"captcha_id"` //required
	Captcha    string `json:"captcha"`    //required
}

//resp for both 'add' and 'login' api
type Msg_Resp_UserInfo struct {
	API_META_STATUS
	User *User `json:"user"`
}

//query
// @Description Msg_Req_QueryUser_Filter
type Msg_Req_QueryUser_Filter struct {
	Id            *int64  `json:"id"`            //optional
	Email_pattern *string `json:"email_pattern"` //optional
}

// @Description Msg_Req_QueryUser
type Msg_Req_QueryUser struct {
	Filter Msg_Req_QueryUser_Filter `json:"filter"` //required
	Limit  int                      `json:"limit"`  //required
	Offset int                      `json:"offset"` //required
}

type Msg_Resp_QueryUser struct {
	API_META_STATUS
	User_list []*User `json:"user_list"`
	Count     int64   `json:"count"`
}

//update
// @Description Msg_Req_UpdateUser_Filter
type Msg_Req_UpdateUser_Filter struct {
	Id []int64 `json:"id"` //required
}

//delete
// @Description Msg_Req_DeleteUser_Filter
type Msg_Req_DeleteUser_Filter struct {
	Id []int64 `json:"id"` //required
}

// @Description Msg_Req_DeleteUser
type Msg_Req_DeleteUser struct {
	Filter Msg_Req_DeleteUser_Filter `json:"filter"` //required
}

// @Description Msg_Req_UpdateUser_To
type Msg_Req_UpdateUser_To struct {
	Forbidden   *bool     `json:"forbidden"`   //optional
	Roles       *[]string `json:"roles"`       //optional
	Permissions *[]string `json:"permissions"` //optional
}

// @Description Msg_Req_UpdateUser
type Msg_Req_UpdateUser struct {
	Filter Msg_Req_UpdateUser_Filter `json:"filter"` //required
	Update Msg_Req_UpdateUser_To     `json:"update"` //required
}

//using API_META_STATUS as update response

//reset pass
// @Description Msg_Req_UserResetPassword
type Msg_Req_UserResetPassword struct {
	Email    string `json:"email"`    //required
	Password string `json:"password"` //required
	VCode    string `json:"vcode"`    //required
}

//req evcode
// @Description Msg_Req_EmailVCode
type Msg_Req_EmailVCode struct {
	Email string `json:"email"` //required
}

//captcha
type Msg_Resp_Captcha struct {
	API_META_STATUS
	Id      string `json:"id"`
	Content string `json:"content"`
}

type Msg_Resp_Auth_Config struct {
	API_META_STATUS
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}
