package commonMsg

type User struct {
	Id          int64
	Email       string
	Password    string
	Token       string
	Forbidden   bool
	Roles       []string
	Permissions []string

	Updated int64
	Created int64
}

//add
// @Description Msg_Req_AddUser
type Msg_Req_AddUser struct {
	Email       string   //required
	Password    string   //required
	Forbidden   bool     //required
	Roles       []string //required
	Permissions []string //required
}

//login
// @Description Msg_Req_UserLogin
type Msg_Req_UserLogin struct {
	Email      string //required
	Password   string //required
	Captcha_id string //required
	Captcha    string //required
}

//resp for both 'add' and 'login' api
type Msg_Resp_UserInfo struct {
	API_META_STATUS
	User *User
}

//query
// @Description Msg_Req_QueryUser_Filter
type Msg_Req_QueryUser_Filter struct {
	Id            *int64  //optional
	Email_pattern *string //optional
}

// @Description Msg_Req_QueryUser
type Msg_Req_QueryUser struct {
	Filter Msg_Req_QueryUser_Filter //required
	Limit  int64                    //required
	Offset int64                    //required
}

type Msg_Resp_QueryUser struct {
	API_META_STATUS
	User_list []*User
	Count     int64
}

//update
// @Description Msg_Req_UpdateUser_Filter
type Msg_Req_UpdateUser_Filter struct {
	Id []int64 //required
}

//delete
// @Description Msg_Req_DeleteUser_Filter
type Msg_Req_DeleteUser_Filter struct {
	Id []int64 //required
}

// @Description Msg_Req_DeleteUser
type Msg_Req_DeleteUser struct {
	Filter Msg_Req_DeleteUser_Filter //required
}

// @Description Msg_Req_UpdateUser_To
type Msg_Req_UpdateUser_To struct {
	Forbidden   *bool     //optional
	Roles       *[]string //optional
	Permissions *[]string //optional
}

// @Description Msg_Req_UpdateUser
type Msg_Req_UpdateUser struct {
	Filter Msg_Req_UpdateUser_Filter //required
	Update Msg_Req_UpdateUser_To     //required
}

//using API_META_STATUS as update response

//reset pass
// @Description Msg_Req_UserResetPassword
type Msg_Req_UserResetPassword struct {
	Email    string //required
	Password string //required
	VCode    string //required
}

//req evcode
// @Description Msg_Req_EmailVCode
type Msg_Req_EmailVCode struct {
	Email string //required
}

//captcha
type Msg_Resp_Captcha struct {
	API_META_STATUS
	Id      string
	Content string
}

type Msg_Resp_Auth_Config struct {
	API_META_STATUS
	Roles       []string
	Permissions []string
}
