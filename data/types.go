package data

type TransferStatus uint8

const NativeToken = "T9yD14Nj9j7xAB4dbGeiX9h8unkKHxuWwb"

const (
	Pending TransferStatus = iota
	Approved
	Rejected
	Waiting
	Canceled
)

func (d TransferStatus) ToInt() int {
	return int(d)
}

func (d TransferStatus) ToString() string {
	switch d {
	case Pending:
		return "pending"
	case Approved:
		return "approved"
	case Rejected:
		return "rejected"
	case Waiting:
		return "waiting"
	case Canceled:
		return "canceled"
	default:
		return ""
	}
}
