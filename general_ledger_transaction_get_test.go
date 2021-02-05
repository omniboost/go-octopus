package octopus_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/omniboost/go-octopus"
)

func TestGeneralLedgerTransactionGet(t *testing.T) {
	req := client.NewGeneralLedgerTransactionGetRequest()
	req.PathParams().Count = 10
	req.QueryParams().CreatedAfter = octopus.DateTime{time.Now().AddDate(0, 0, -7)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
