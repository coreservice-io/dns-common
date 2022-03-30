package data

type Cert struct {
	ID             uint
	Name           string
	UserId         uint
	CertContent    string
	KeyContent     string
	ExpirationTime int64
	Hash           string
	RelatedDomain  []string

	Updated int64
	Created int64
}
