package hotp

import (
	"github.com/pquerna/otp/hotp"
)

type HotpHelper interface {
	GenerateCode(secret string, counter uint64) (string, error)
}

type hotpHelper struct {
}

func NewTotpHelper() (HotpHelper, error) {
	return &hotpHelper{}, nil
}

func (hh *hotpHelper) GenerateCode(secret string, counter uint64) (string, error) {
	return hotp.GenerateCode(secret, counter)

}
