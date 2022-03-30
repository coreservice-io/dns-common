package data

type Record struct {
	ID        uint
	DomainId  uint
	Name      string
	Type      string //enum
	Forbidden bool
	TTL       uint32

	Updated int64
	Created int64
}
