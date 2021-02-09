package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewFinancialDiversBookingsGetRequest() FinancialDiversBookingsGetRequest {
	r := FinancialDiversBookingsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type FinancialDiversBookingsGetRequest struct {
	client      *Client
	queryParams *FinancialDiversBookingsGetQueryParams
	pathParams  *FinancialDiversBookingsGetPathParams
	method      string
	headers     http.Header
	requestBody FinancialDiversBookingsGetRequestBody
}

func (r FinancialDiversBookingsGetRequest) NewQueryParams() *FinancialDiversBookingsGetQueryParams {
	return &FinancialDiversBookingsGetQueryParams{}
}

type FinancialDiversBookingsGetQueryParams struct {
	BookYearID    string `schema:"bookyearId"`
	JournalKey    string `schema:"journalKey"`
	DocumentSeqNr string `schema:"documentSeqNr"`
}

func (p FinancialDiversBookingsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *FinancialDiversBookingsGetRequest) QueryParams() *FinancialDiversBookingsGetQueryParams {
	return r.queryParams
}

func (r FinancialDiversBookingsGetRequest) NewPathParams() *FinancialDiversBookingsGetPathParams {
	return &FinancialDiversBookingsGetPathParams{}
}

type FinancialDiversBookingsGetPathParams struct {
	DossierID string `schema:"dossierId"`
}

func (p *FinancialDiversBookingsGetPathParams) Params() map[string]string {
	return map[string]string{
		"dossierId": p.DossierID,
	}
}

func (r *FinancialDiversBookingsGetRequest) PathParams() *FinancialDiversBookingsGetPathParams {
	return r.pathParams
}

func (r *FinancialDiversBookingsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *FinancialDiversBookingsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *FinancialDiversBookingsGetRequest) Method() string {
	return r.method
}

func (r FinancialDiversBookingsGetRequest) NewRequestBody() FinancialDiversBookingsGetRequestBody {
	return FinancialDiversBookingsGetRequestBody{}
}

type FinancialDiversBookingsGetRequestBody struct {
}

func (r *FinancialDiversBookingsGetRequest) RequestBody() *FinancialDiversBookingsGetRequestBody {
	return &r.requestBody
}

func (r *FinancialDiversBookingsGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *FinancialDiversBookingsGetRequest) SetRequestBody(body FinancialDiversBookingsGetRequestBody) {
	r.requestBody = body
}

func (r *FinancialDiversBookingsGetRequest) NewResponseBody() *FinancialDiversBookingsGetResponseBody {
	return &FinancialDiversBookingsGetResponseBody{}
}

type FinancialDiversBookingsGetResponseBody []struct{}

func (r *FinancialDiversBookingsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("dossiers/{{.dossierId}}/financialdiversbookings", r.PathParams())
	return &u
}

func (r *FinancialDiversBookingsGetRequest) Do() (FinancialDiversBookingsGetResponseBody, error) {
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
