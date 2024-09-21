package sid

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/helper/convert"
	"github.com/sony/sonyflake"
)

type Sid struct {
	sf *sonyflake.Sonyflake
}

func NewSid() *Sid {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		panic("sonyflake not created")
	}
	return &Sid{sf}
}

func (s Sid) GenString() (string, error) {
	id, err := s.sf.NextID()
	if err != nil {
		return "", err
	}
	return convert.IntToBase62(int(id)), nil
}

func (s Sid) GenUint64() (uint64, error) {
	return s.sf.NextID()
}

func (s Sid) GenRandomToken(baseLen int8) (string, error) {
	b := make([]byte, baseLen)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	nxtId, err := s.GenString()
	if err != nil {
		return "", err
	}
	b = append(b, []byte(nxtId)...)
	return base64.URLEncoding.EncodeToString(b), nil
}
