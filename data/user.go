package data

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
