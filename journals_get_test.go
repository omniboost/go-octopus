package octopus_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/omniboost/go-octopus"
)

func TestJournalsGet(t *testing.T) {
	req := client.NewJournalsGetRequest()
	req.PathParams().DossierID = os.Getenv("DOSSIER_ID")
	req.PathParams().BookYearID = os.Getenv("BOOKYEAR_ID")
	req.SetRequestBody(octopus.JournalsGetRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
