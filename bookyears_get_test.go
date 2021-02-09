package octopus_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/omniboost/go-octopus"
)

func TestBookyearsGet(t *testing.T) {
	req := client.NewBookyearsGetRequest()
	req.PathParams().DossierID = os.Getenv("DOSSIER_ID")
	req.SetRequestBody(octopus.BookyearsGetRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
