package smb

import (
	"context"
	"errors"
	"time"

	"github.com/stacktitan/smb/smb"
	"github.com/w3security/reconbot/lib/util"
)

func Check(Host, Username, Domain, Password string, Port int) (bool, error) {
	status := make(chan error)
	ctx, cancel := context.WithTimeout(util.Ctx_global, time.Second*5)
	defer cancel()
	options := smb.Options{
		Host:        Host,
		Port:        Port,
		User:        Username,
		Password:    Password,
		Domain:      Domain,
		Workstation: "",
	}
	//开始进行SMB连接
	go func() {
		session, err := smb.NewSession(options, false)
		if err != nil {
			status <- err
			return
		}
		defer session.Close()
		if session.IsAuthenticated == false {
			status <- err
			return
		}
		status <- nil
	}()

	select {
	case <-ctx.Done():
		return false, errors.New("timeout")
	case err := <-status:
		if status != nil {
			return false, err
		}
		return true, nil
	}

}