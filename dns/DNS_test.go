package dns

import (
	"bytes"
	"testing"
)

func TestFlagEncode(t *testing.T) {
    df := DNSFlags{
        QR:  true,
        OpCode: 4,
        AA: true,
        TC: false,
        RD: true,
        RA: false,
        Rcode: 2,
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
