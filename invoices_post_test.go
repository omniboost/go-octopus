package octopus_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/omniboost/go-octopus"
)

func TestInvoicesPost(t *testing.T) {
	req := client.NewInvoicesPostRequest()
	req.PathParams().DossierID = os.Getenv("OCTO_DOSSIER_ID")
	req.SetRequestBody(octopus.InvoicesPostRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
