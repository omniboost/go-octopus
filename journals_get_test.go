package octopus_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/omniboost/go-octopus"
)

func TestJournalsGet(t *testing.T) {
	var err error

	req := client.NewJournalsGetRequest()
	req.PathParams().DossierID = os.Getenv("OCTO_DOSSIER_ID")
	req.PathParams().BookYearID, err = strconv.Atoi(os.Getenv("OCTO_BOOKYEAR_ID"))
	if err != nil {
		t.Error(err)
	}
	req.SetRequestBody(octopus.JournalsGetRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
