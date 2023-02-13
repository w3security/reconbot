package fingerprint

import (
	_ "embed"

	"github.com/w3security/reconbot/lib/util"
)

//go:embed dicts/localFinger.json
var localFinger string

func init() {
	util.RegInitFunc(func() {
		localFinger = util.GetVal4File("localFinger", localFinger)
	})
}
