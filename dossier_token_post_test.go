package octopus_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDossierTokenPost(t *testing.T) {
	req := client.NewDossierTokenPostRequest()
	req.QueryParams().DossierID = os.Getenv("DOSSIER_ID")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
