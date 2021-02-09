package octopus_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/omniboost/go-octopus"
)

func TestVATCodes(t *testing.T) {
	req := client.NewVATCodesRequest()
	req.PathParams().DossierID = os.Getenv("DOSSIER_ID")
	req.SetRequestBody(octopus.VATCodesRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
