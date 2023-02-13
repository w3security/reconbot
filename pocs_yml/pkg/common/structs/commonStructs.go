package structs

import (
	"net/http"
	"strings"

	"github.com/w3security/reconbot/pocs_yml/pkg/xray/structs"
)

var (
	CeyeApi                  string
	CeyeDomain               string
	ReversePlatformType      structs.ReverseType
	DnslogCNGetDomainRequest *http.Request
	DnslogCNGetRecordRequest *http.Request
)

func InitReversePlatform(api, domain string) {
	if api != "" && domain != "" && strings.HasSuffix(domain, ".ceye.io") {
		CeyeApi = api
		CeyeDomain = domain
		ReversePlatformType = structs.ReverseType_Ceye
	} else {
		ReversePlatformType = structs.ReverseType_DnslogCN

		// 设置请求相关参数
		DnslogCNGetDomainRequest, _ = http.NewRequest("GET", "http://dnslog.cn/getdomain.php", nil)
		DnslogCNGetRecordRequest, _ = http.NewRequest("GET", "http://dnslog.cn/getrecords.php", nil)

	}
}
