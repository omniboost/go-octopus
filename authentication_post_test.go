package octopus_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/omniboost/go-octopus"
)

func TestAuthenticationPost(t *testing.T) {
	req := client.NewAuthenticationPostRequest()
	req.SetRequestBody(octopus.AuthenticationPostRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
