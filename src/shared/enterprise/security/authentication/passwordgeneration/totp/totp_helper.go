package totp

import (
	"time"

	"github.com/pquerna/otp/totp"
)

type TotpHelper interface {
	GenerateCode(secret string, time time.Time) (string, error)
	GenerateCodeUsingCurrentTime(secret string) (string, error)
}

type totpHelper struct {
}

func NewTotpHelper() (TotpHelper, error) {
	return &totpHelper{}, nil
}

func (th *totpHelper) GenerateCode(secret string, time time.Time) (string, error) {
	return totp.GenerateCode(secret, time)

}

func (th *totpHelper) GenerateCodeUsingCurrentTime(secret string) (string, error) {
	return totp.GenerateCode(secret, time.Now())

}
