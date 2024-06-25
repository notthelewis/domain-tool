package dns

import (
	qt "domain-tool/dns/QueryType"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
)

// DNSMessage is the fundamental interchange format between client and server.
type DNSMessage struct {
	Identification        uint16   // ID number decided by the client, which is repeated by the server upon response
	Flags                 DNSFlags // Message specific flags set dependent on use case
	Questions             []DNSQuestion
	Answers               []ResourceRecord
	Authority             []ResourceRecord
	AdditionalInformation []ResourceRecord
}

func (dnsMsg *DNSMessage) Encode() []byte {
	header := make([]byte, 0, 12)

    qLen := len(dnsMsg.Questions)
    anLen := len(dnsMsg.Answers)
    auLen := len(dnsMsg.Authority)
    adLen := len(dnsMsg.AdditionalInformation)
    // bodyLen := qLen + anLen + auLen + adLen
    // body := make([]byte, 0, bodyLen)

	// Write Identification
	binary.BigEndian.PutUint16(header, uint16(rand.UintN(0xFFFF)))

	// Write Flags
	binary.BigEndian.PutUint16(header, dnsMsg.Flags.Encode())

	// Write varint fields
	binary.BigEndian.PutUint16(header, uint16(qLen))
	binary.BigEndian.PutUint16(header, uint16(anLen))
	binary.BigEndian.PutUint16(header, uint16(auLen))
	binary.BigEndian.PutUint16(header, uint16(adLen))


    // for _, r := range dnsMsg.Questions {
    //
    // }

	return header
}

// DNSFlags are the flags which indicate properties about the DNSMessage.
type DNSFlags struct {
	QR     bool  // 0 == query, 1 == response
	OpCode uint8 // 4bit field. 0 == standard query, 1 == inverse query, 2 == server status request
	AA     bool  // Name server is authorative
	TC     bool  // Reply exceeded 512 bytes
	RD     bool  // Recursion desired
	RA     bool  // Recursion available
	Rcode  uint8 // 4bit field. 0 == no error, 3 == name error
}

// Encode will take a DNSFlags and encode it as a single uint16
func (F *DNSFlags) Encode() uint16 {
	return uint16((b2i(F.QR)<<7|int(F.OpCode)<<3|b2i(F.AA)<<2|b2i(F.RD))<<8 | b2i(F.RA)<<7 | 0x00<<4 | int(F.Rcode))
}

func (F DNSFlags) String() string {
	return fmt.Sprintf("{QR:%t, OpCode:%v, AA:%t,TC:%t, RD:%t, RA:%t, Rcode:%v}", F.QR, F.OpCode, F.AA, F.TC, F.RD, F.RA, F.Rcode)
}

// DNSQuestion is an individual request generated by the client for a specific piece of data
type DNSQuestion struct {
	QueryName  []Label
	QueryType  qt.QueryType
	QueryClass uint16
}

// Label is part of a domain, i.e. in `test.com`, `test` and `com` are labels
type Label struct {
	Name  []byte
}

// Encode will take in a label and encode it into a byte slice in the format required by DNS protocol.
// The format is a length-prepended, variable length byte array. 
// I.e. `test.com` will be encoded as two separate labels: `4test`, `3com`
func (L *Label) Encode() []byte {
    if len(L.Name) > 63 {
        panic("Max length of 63 for label exceeded")
    }

    buf := make([]byte, 0, len(L.Name)+1)
    buf = append(buf, uint8(len(L.Name)))
    buf = append(buf, L.Name...)

    return buf
}

// ResourceRecord is used for server responses to questions
type ResourceRecord struct {
	DomainName         []Label
	Type               qt.QueryType
	Class              uint16
	TTL                uint16
	ResourceDataLength []byte
	ResourceData       []byte
}

// b2i converts a bool to integer
func b2i(b bool) int {
	if b {
		return 1
	}

	return 0
}
