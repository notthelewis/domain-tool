package QueryType

type QueryType uint16

func New(qtype string) QueryType {
	switch qtype {
	case "A":
		return A

	case "NS":
		return NS

	case "CNAME":
		return CNAME

	case "PTR":
		return PTR 

	case "MX":
		return MX

	case "ANY":
		return ANY

		// TODO: Add em all
	default:
		panic("unknown query type")
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
