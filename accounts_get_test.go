package octopus_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/omniboost/go-octopus"
)

func TestAccountsGet(t *testing.T) {
	req := client.NewAccountsGetRequest()
	req.PathParams().DossierID = os.Getenv("OCTO_DOSSIER_ID")
	req.QueryParams().BookYearID = os.Getenv("OCTO_BOOKYEAR_ID")
	req.SetRequestBody(octopus.AccountsGetRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
