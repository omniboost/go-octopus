package octopus_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/omniboost/go-octopus"
)

func TestProductPost(t *testing.T) {
	req := client.NewProductPostRequest()
	req.SetRequestBody(octopus.ProductPostRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
