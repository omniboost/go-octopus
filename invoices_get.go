package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewInvoicesGetRequest() InvoicesGetRequest {
	r := InvoicesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InvoicesGetRequest struct {
	client      *Client
	queryParams *InvoicesGetQueryParams
	pathParams  *InvoicesGetPathParams
	method      string
	headers     http.Header
	requestBody InvoicesGetRequestBody
}

func (r InvoicesGetRequest) NewQueryParams() *InvoicesGetQueryParams {
	return &InvoicesGetQueryParams{}
}

type InvoicesGetQueryParams struct {
	BookYearID    string `schema:"bookyearId"`
	JournalKey    string `schema:"journalKey"`
	DocumentSeqNr string `schema:"documentSeqNr"`
}

func (p InvoicesGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, errors.WithStack(err)
	}

	return params, nil
}

func (r *InvoicesGetRequest) QueryParams() *InvoicesGetQueryParams {
	return r.queryParams
}

func (r InvoicesGetRequest) NewPathParams() *InvoicesGetPathParams {
	return &InvoicesGetPathParams{}
}

type InvoicesGetPathParams struct {
	DossierID string `schema:"dossierId"`
}

func (p *InvoicesGetPathParams) Params() map[string]string {
	return map[string]string{
		"dossierId": p.DossierID,
	}
}

func (r *InvoicesGetRequest) PathParams() *InvoicesGetPathParams {
	return r.pathParams
}

func (r *InvoicesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InvoicesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoicesGetRequest) Method() string {
	return r.method
}

func (r InvoicesGetRequest) NewRequestBody() InvoicesGetRequestBody {
	return InvoicesGetRequestBody{}
}

type InvoicesGetRequestBody struct {
}

func (r *InvoicesGetRequest) RequestBody() *InvoicesGetRequestBody {
	return &r.requestBody
}

func (r *InvoicesGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *InvoicesGetRequest) SetRequestBody(body InvoicesGetRequestBody) {
	r.requestBody = body
}

func (r *InvoicesGetRequest) NewResponseBody() *InvoicesGetResponseBody {
	return &InvoicesGetResponseBody{}
}

type InvoicesGetResponseBody []struct{}

func (r *InvoicesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("dossiers/{{.dossierId}}/invoices", r.PathParams())
	return &u
}

func (r *InvoicesGetRequest) Do() (InvoicesGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), errors.WithStack(err)
	}

	token, err := r.client.DossierToken(r.PathParams().DossierID)
	if err != nil {
		return *r.NewResponseBody(), errors.WithStack(err)
	}
	req.Header.Add("DossierToken", token)

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), errors.WithStack(err)
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, errors.WithStack(err)
}
