package elastic

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic"
	"github.com/w3security/reconbot/pkg/hydra/snmp"
)

// 9200/tcp open  wap-wsp
func ScanElastic(svs *snmp.Service) (err error, result *snmp.ScanResult) {
	result.Service = svs
	client, err := elastic.NewClient(elastic.SetURL(fmt.Sprintf("http://%v:%v", svs.Ip, svs.Port)),
		elastic.SetMaxRetries(3),
		elastic.SetBasicAuth(svs.Username, svs.Password),
	)
	if err == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, _, err = client.Ping(fmt.Sprintf("http://%v:%v", svs.Ip, svs.Port)).Do(ctx)
		if err == nil {
			result.Result = true
		}
	}
	return err, result
}
