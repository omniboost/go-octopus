package octopus_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/omniboost/go-octopus"
)

func TestRelationsPut(t *testing.T) {
	req := client.NewRelationsPutRequest()
	req.PathParams().DossierID = os.Getenv("OCTO_DOSSIER_ID")
	req.SetRequestBody(octopus.RelationsPutRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
