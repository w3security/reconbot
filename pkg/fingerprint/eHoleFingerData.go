package fingerprint

import (
	_ "embed"

	"github.com/w3security/reconbot/lib/util"
)

//go:embed dicts/eHoleFinger.json
var eHoleFinger string

func init() {
	util.RegInitFunc(func() {
		eHoleFinger = util.GetVal4File("eHoleFinger", eHoleFinger)
	})
}
