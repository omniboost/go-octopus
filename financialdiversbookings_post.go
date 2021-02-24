package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewFinancialDiversBookingsPostRequest() FinancialDiversBookingsPostRequest {
	r := FinancialDiversBookingsPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type FinancialDiversBookingsPostRequest struct {
	client      *Client
	queryParams *FinancialDiversBookingsPostQueryParams
	pathParams  *FinancialDiversBookingsPostPathParams
	method      string
	headers     http.Header
	requestBody FinancialDiversBookingsPostRequestBody
}

func (r FinancialDiversBookingsPostRequest) NewQueryParams() *FinancialDiversBookingsPostQueryParams {
	return &FinancialDiversBookingsPostQueryParams{}
}

type FinancialDiversBookingsPostQueryParams struct {
	BookYearID    string `schema:"bookyearId"`
	JournalKey    string `schema:"journalKey"`
	DocumentSeqNr string `schema:"documentSeqNr"`
}

func (p FinancialDiversBookingsPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *FinancialDiversBookingsPostRequest) QueryParams() *FinancialDiversBookingsPostQueryParams {
	return r.queryParams
}

func (r FinancialDiversBookingsPostRequest) NewPathParams() *FinancialDiversBookingsPostPathParams {
	return &FinancialDiversBookingsPostPathParams{}
}

type FinancialDiversBookingsPostPathParams struct {
	DossierID string `schema:"dossierId"`
}

func (p *FinancialDiversBookingsPostPathParams) Params() map[string]string {
	return map[string]string{
		"dossierId": p.DossierID,
	}
}

func (r *FinancialDiversBookingsPostRequest) PathParams() *FinancialDiversBookingsPostPathParams {
	return r.pathParams
}

func (r *FinancialDiversBookingsPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *FinancialDiversBookingsPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *FinancialDiversBookingsPostRequest) Method() string {
	return r.method
}

func (r FinancialDiversBookingsPostRequest) NewRequestBody() FinancialDiversBookingsPostRequestBody {
	return FinancialDiversBookingsPostRequestBody{}
}

type FinancialDiversBookingsPostRequestBody FinancialDiversBookingAndAttachmentRequest

func (r *FinancialDiversBookingsPostRequest) RequestBody() *FinancialDiversBookingsPostRequestBody {
	return &r.requestBody
}

func (r *FinancialDiversBookingsPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *FinancialDiversBookingsPostRequest) SetRequestBody(body FinancialDiversBookingsPostRequestBody) {
	r.requestBody = body
}

func (r *FinancialDiversBookingsPostRequest) NewResponseBody() *FinancialDiversBookingsPostResponseBody {
	return &FinancialDiversBookingsPostResponseBody{}
}

type FinancialDiversBookingsPostResponseBody []struct{}

func (r *FinancialDiversBookingsPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("dossiers/{{.dossierId}}/financialdiversbookings", r.PathParams())
	return &u
}

func (r *FinancialDiversBookingsPostRequest) Do() (FinancialDiversBookingsPostResponseBody, error) {
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
