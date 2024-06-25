package QueryType

import "errors"

type QueryType uint16

func New(qtype string) (QueryType, error) {
	switch qtype {
	case "A":
		return A, nil

	case "NS":
		return NS, nil

	case "CNAME":
		return CNAME, nil

	case "PTR":
		return PTR, nil

	case "MX":
		return MX, nil

	case "ANY":
		return ANY, nil

		// TODO: Add em all
	default:
		return 0, errors.New("unknown query type")
	}
}

func (qt *QueryType) Get() uint16 {
	return uint16(*qt)
}

const (
	A     QueryType = 1
	NS    QueryType = 2
	CNAME QueryType = 5
	PTR   QueryType = 12
	MX    QueryType = 15
	ANY   QueryType = 255
)
