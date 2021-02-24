package octopus_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/omniboost/go-octopus"
)

func TestRelationsGet(t *testing.T) {
	req := client.NewRelationsGetRequest()
	req.PathParams().DossierID = os.Getenv("DOSSIER_ID")
	req.SetRequestBody(octopus.RelationsGetRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
