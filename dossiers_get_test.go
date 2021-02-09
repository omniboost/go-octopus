package octopus_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/omniboost/go-octopus"
)

func TestDossiersGet(t *testing.T) {
	req := client.NewDossiersGetRequest()
	req.SetRequestBody(octopus.DossiersGetRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
