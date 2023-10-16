package crockford

import (
	"encoding/base32"
	"bytes"

	"github.com/google/uuid"
)

func NewID() string {
	id := uuid.New()
	str := id.String()
	str = Encode([]byte(str[:16]))
	return str[:26]
}

func Encode(b []byte) string {
	var buf bytes.Buffer
	symbols := "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
	enc := base32.NewEncoder(base32.NewEncoding(symbols), &buf)
	enc.Write(b)
	enc.Close()
	return buf.String()
}
