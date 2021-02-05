package octopus_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestProductsGet(t *testing.T) {
	req := client.NewProductsGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
