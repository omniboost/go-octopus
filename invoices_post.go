package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewInvoicesPostRequest() InvoicesPostRequest {
	r := InvoicesPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InvoicesPostRequest struct {
	client      *Client
	queryParams *InvoicesPostQueryParams
	pathParams  *InvoicesPostPathParams
	method      string
	headers     http.Header
	requestBody InvoicesPostRequestBody
}

func (r InvoicesPostRequest) NewQueryParams() *InvoicesPostQueryParams {
	return &InvoicesPostQueryParams{}
}

type InvoicesPostQueryParams struct{}

func (p InvoicesPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *InvoicesPostRequest) QueryParams() *InvoicesPostQueryParams {
	return r.queryParams
}

func (r InvoicesPostRequest) NewPathParams() *InvoicesPostPathParams {
	return &InvoicesPostPathParams{}
}

type InvoicesPostPathParams struct {
	DossierID string `schema:"dossierId"`
}

func (p *InvoicesPostPathParams) Params() map[string]string {
	return map[string]string{
		"dossierId": p.DossierID,
	}
}

func (r *InvoicesPostRequest) PathParams() *InvoicesPostPathParams {
	return r.pathParams
}

func (r *InvoicesPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InvoicesPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoicesPostRequest) Method() string {
	return r.method
}

func (r InvoicesPostRequest) NewRequestBody() InvoicesPostRequestBody {
	return InvoicesPostRequestBody{}
}

type InvoicesPostRequestBody FinancialDiversBookingServiceData

func (r *InvoicesPostRequest) RequestBody() *InvoicesPostRequestBody {
	return &r.requestBody
}

func (r *InvoicesPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *InvoicesPostRequest) SetRequestBody(body InvoicesPostRequestBody) {
	r.requestBody = body
}

func (r *InvoicesPostRequest) NewResponseBody() *InvoicesPostResponseBody {
	return &InvoicesPostResponseBody{}
}

type InvoicesPostResponseBody []struct{}

func (r *InvoicesPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("dossiers/{{.dossierId}}/invoices", r.PathParams())
	return &u
}

func (r *InvoicesPostRequest) Do() (InvoicesPostResponseBody, error) {
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
