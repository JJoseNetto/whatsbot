package enums

type MessageStatus int

const (
	Pending = iota
	Approved
	Reproved
)

func (s MessageStatus) String() string {
	return [...]string{"PENDING", "APPROVED", "REPROVED"}[s]
}
