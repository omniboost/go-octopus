package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewBuysellbookingsGetRequest() BuysellbookingsGetRequest {
	r := BuysellbookingsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BuysellbookingsGetRequest struct {
	client      *Client
	queryParams *BuysellbookingsGetQueryParams
	pathParams  *BuysellbookingsGetPathParams
	method      string
	headers     http.Header
	requestBody BuysellbookingsGetRequestBody
}

func (r BuysellbookingsGetRequest) NewQueryParams() *BuysellbookingsGetQueryParams {
	return &BuysellbookingsGetQueryParams{}
}

type BuysellbookingsGetQueryParams struct {
	BookYearID    string `schema:"bookyearId"`
	JournalKey    string `schema:"journalKey"`
	DocumentSeqNr string `schema:"documentSeqNr"`
}

func (p BuysellbookingsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BuysellbookingsGetRequest) QueryParams() *BuysellbookingsGetQueryParams {
	return r.queryParams
}

func (r BuysellbookingsGetRequest) NewPathParams() *BuysellbookingsGetPathParams {
	return &BuysellbookingsGetPathParams{}
}

type BuysellbookingsGetPathParams struct {
	DossierID string `schema:"dossierId"`
}

func (p *BuysellbookingsGetPathParams) Params() map[string]string {
	return map[string]string{
		"dossierId": p.DossierID,
	}
}

func (r *BuysellbookingsGetRequest) PathParams() *BuysellbookingsGetPathParams {
	return r.pathParams
}

func (r *BuysellbookingsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BuysellbookingsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *BuysellbookingsGetRequest) Method() string {
	return r.method
}

func (r BuysellbookingsGetRequest) NewRequestBody() BuysellbookingsGetRequestBody {
	return BuysellbookingsGetRequestBody{}
}

type BuysellbookingsGetRequestBody struct {
}

func (r *BuysellbookingsGetRequest) RequestBody() *BuysellbookingsGetRequestBody {
	return &r.requestBody
}

func (r *BuysellbookingsGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *BuysellbookingsGetRequest) SetRequestBody(body BuysellbookingsGetRequestBody) {
	r.requestBody = body
}

func (r *BuysellbookingsGetRequest) NewResponseBody() *BuysellbookingsGetResponseBody {
	return &BuysellbookingsGetResponseBody{}
}

type BuysellbookingsGetResponseBody []struct{}

func (r *BuysellbookingsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("dossiers/{{.dossierId}}/buysellbookings", r.PathParams())
	return &u
}

func (r *BuysellbookingsGetRequest) Do() (BuysellbookingsGetResponseBody, error) {
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
