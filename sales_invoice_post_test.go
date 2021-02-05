package octopus_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/omniboost/go-octopus"
)

func TestSalesInvoicePost(t *testing.T) {
	req := client.NewSalesInvoicePostRequest()
	req.SetRequestBody(octopus.SalesInvoicePostRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
