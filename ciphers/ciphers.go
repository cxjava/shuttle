package ciphers

import (
	"fmt"
	"github.com/cxjava/shuttle"
	"github.com/cxjava/shuttle/ciphers/ssaead"
	"github.com/cxjava/shuttle/ciphers/ssstream"
)

type ConnDecorate func(password string, conn shuttle.IConn) (shuttle.IConn, error)

//加密装饰
func CipherDecorate(password, method string, conn shuttle.IConn) (shuttle.IConn, error) {
	d := ssstream.GetStreamCiphers(method)
	if d != nil {
		return d(password, conn)
	}
	d = ssaead.GetAEADCiphers(method)
	if d != nil {
		return d(password, conn)
	}
	return nil, fmt.Errorf("[SS Cipher] not support : %s", method)
}
