package dns

import (
	"strconv"
	"testing"
)

func TestToggleBit(t *testing.T) {
    df := DNSFlags{
        QR:  true,
        OpCode: 4,
        AA: true,
        TC: false,
        RD: true,
        RA: false,
        Rcode: 2,
    }

    t.Log(df)
    t.Log(strconv.FormatInt(int64(df.Encode()), 2))
}
