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
	NumberOfQuestions     uint16   // The amount of questions asked by client, usually 1
	NumberOfAnswerRRs     uint16   // The amount of answer resource records
	NumberOfAuthorityRRs  uint16   // The amount of authority resource records
	NumberOfAdditionalRRs uint16   // The amount of additional resource records
	Questions             []DNSQuestion
	Answers               []ResourceRecord
	Authority             []ResourceRecord
	AdditionalInformation []ResourceRecord
}

func (dnsMsg *DNSMessage) Encode() []byte {
	// TODO: Object pool
	header := make([]byte, 0, 12)

	// Write Identification
	binary.BigEndian.PutUint16(header, uint16(rand.UintN(0xFFFF)))
	// Write Flags
	binary.BigEndian.PutUint16(header, dnsMsg.Flags.Encode())

	// Write varint fields
	binary.BigEndian.PutUint16(header, uint16(len(dnsMsg.Questions)))
	binary.BigEndian.PutUint16(header, uint16(len(dnsMsg.Authority)))
	binary.BigEndian.PutUint16(header, uint16(len(dnsMsg.AdditionalInformation)))


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

// Encode will take a DNSFlags and encode it as a single uint16. Here be dragons
func (F *DNSFlags) Encode() uint16 {
	return uint16((b2i(F.QR)<<7|int(F.OpCode)<<3|b2i(F.AA)<<2|b2i(F.RD))<<8 | b2i(F.RA)<<7 | 0x00<<4 | int(F.Rcode))
}

func (F DNSFlags) String() string {
	return fmt.Sprintf("\nQR:%t,\nOpCode:%v,\nAA:%t,\nTC:%t,\nRD:%t,\nRA:%t,\nRcode:%v\n", F.QR, F.OpCode, F.AA, F.TC, F.RD, F.RA, F.Rcode)
}

type DNSQuestion struct {
	QueryName  []Label
	QueryType  qt.QueryType
	QueryClass uint16
}

type Label struct {
	Count uint8 // Max 63
	Name  []byte
}

type ResourceRecord struct {
	DomainName         []Label
	Type               qt.QueryType
	Class              uint16
	TTL                uint16
	ResourceDataLength []byte
	ResourceData       []byte
}

func b2i(b bool) int {
	if b {
		return 1
	}

	return 0
}
