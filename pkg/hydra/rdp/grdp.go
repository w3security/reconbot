package rdp

import (
	"fmt"

	"github.com/w3security/reconbot/pkg/kscan/lib/grdp"
)

func Check(ip, domain, login, password string, port int, protocol string) (bool, error) {
	var err error
	target := fmt.Sprintf("%s:%d", ip, port)
	if protocol == grdp.PROTOCOL_SSL {
		err = grdp.LoginForSSL(target, domain, login, password)
	} else {
		err = grdp.LoginForRDP(target, domain, login, password)
	}
	//err = grdp.Login(target, domain, login, password)
	//slog.Println(slog.INFO, target, domain, login, password)
	//slog.Println(slog.INFO, err)
	if err != nil {
		return false, err
	}
	return true, err
}
