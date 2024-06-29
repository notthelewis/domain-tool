package dns

import (
	"bytes"
	"domain-tool/dns/QueryType"
	"testing"
)

func TestFlagEncode(t *testing.T) {
	df := DNSFlags{
		QR:     true,
		OpCode: 4,
		AA:     true,
		TC:     false,
		RD:     true,
		RA:     false,
		Rcode:  2,
	}

	if df.Encode() != 0b1010010100000010 {
		t.Fatal("bad encoding for flags")
	}
}

func TestLabelEncode(t *testing.T) {
	l := Label{
		Name: []byte("test"),
	}

	if !bytes.Equal(l.Encode(), []byte{4, 't', 'e', 's', 't'}) {
		t.Fatalf("bad encoding for label")
	}
}

func TestResourceRecordEncode(t *testing.T) {
	rr := ResourceRecord{
        DomainName: []Label{
            { Name: []byte("test") },
            { Name: []byte("com") },
        },
		Type:               QueryType.New("A"),
		Class:              1,
		TTL:                60,
		ResourceDataLength: 1,
		ResourceData:       []byte{0x00},
	}

    want := []byte{4, 't', 'e', 's', 't', 3, 'c', 'o', 'm', 0, 0, 1, 0, 1, 0, 60, 0, 1, 0}
    got := rr.Encode()

    if !bytes.Equal(want, got) {
        t.Fatalf("bad encoding for resource record.\nWanted: %v,\n Got: %v\n", want, got)
    }
}

func TestDNSQuestionEncode(t *testing.T) {
	dq := DNSQuestion{
        QueryName: []Label{
            { Name: []byte("test")},
            { Name: []byte("com") },
        },
        QueryType: QueryType.New("A"),
        QueryClass: 1,
    }

    want := []byte{4, 't', 'e', 's', 't', 3, 'c', 'o', 'm', 0, 0, 1, 0, 1}
    got := dq.Encode()

    if !bytes.Equal(want, got) {
        t.Fatalf("bad encoding for DNSQuestion.\nWanted: %v,\n Got: %v\n", want, got)
    }
}

// func TestDNSMessageEncode(t *testing.T) {
//     dmsg := DNSMessage{
//         Flags: DNSFlags{
//             QR:     true,
//             OpCode: 4,
//             AA:     true,
//             TC:     false,
//             RD:     true,
//             RA:     false,
//             Rcode:  2,
//         },
//         Questions: []DNSQuestion{
//             {
//                 QueryName: []Label{
//                     { Name: []byte("test") },
//                     { Name: []byte("com") },
//                 },
//                 QueryType: QueryType.New("A"),
//                 QueryClass: 1,
//             },
//         },
//     }
// }
