package data

type Domain struct {
	ID             uint
	Name           string
	ExpirationTime int64
	Forbidden      bool
	UserId         uint

	Updated int64
	Created int64
}
