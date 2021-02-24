package octopus_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/omniboost/go-octopus"
)

func TestBuysellbookingsGet(t *testing.T) {
	req := client.NewBuysellbookingsGetRequest()
	req.PathParams().DossierID = os.Getenv("DOSSIER_ID")
	req.QueryParams().BookYearID = os.Getenv("BOOKYEAR_ID")
	req.QueryParams().JournalKey = os.Getenv("JOURNAL_KEY")
	req.QueryParams().DocumentSeqNr = os.Getenv("DOCUMENT_SEQ_NR")
	req.SetRequestBody(octopus.BuysellbookingsGetRequestBody{})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
