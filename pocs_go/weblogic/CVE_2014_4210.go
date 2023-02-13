package weblogic

import (
	"github.com/w3security/reconbot/lib/util"
)

func CVE_2014_4210(url string) bool {
	if req, err := util.HttpRequset(url+"/uddiexplorer/SearchPublicRegistries.jsp", "GET", "", false, nil); err == nil {
		if req.StatusCode == 200 {
			util.SendLog(req.RequestUrl, "CVE-2014-4210", "Found vuln Weblogic", "")
			return true
		}
	}
	return false
}
