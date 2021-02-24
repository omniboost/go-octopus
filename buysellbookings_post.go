package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewBuysellbookingsPostRequest() BuysellbookingsPostRequest {
	r := BuysellbookingsPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BuysellbookingsPostRequest struct {
	client      *Client
	queryParams *BuysellbookingsPostQueryParams
	pathParams  *BuysellbookingsPostPathParams
	method      string
	headers     http.Header
	requestBody BuysellbookingsPostRequestBody
}

func (r BuysellbookingsPostRequest) NewQueryParams() *BuysellbookingsPostQueryParams {
	return &BuysellbookingsPostQueryParams{}
}

type BuysellbookingsPostQueryParams struct{}

func (p BuysellbookingsPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BuysellbookingsPostRequest) QueryParams() *BuysellbookingsPostQueryParams {
	return r.queryParams
}

func (r BuysellbookingsPostRequest) NewPathParams() *BuysellbookingsPostPathParams {
	return &BuysellbookingsPostPathParams{}
}

type BuysellbookingsPostPathParams struct {
	DossierID string `schema:"dossierId"`
}

func (p *BuysellbookingsPostPathParams) Params() map[string]string {
	return map[string]string{
		"dossierId": p.DossierID,
	}
}

func (r *BuysellbookingsPostRequest) PathParams() *BuysellbookingsPostPathParams {
	return r.pathParams
}

func (r *BuysellbookingsPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BuysellbookingsPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *BuysellbookingsPostRequest) Method() string {
	return r.method
}

func (r BuysellbookingsPostRequest) NewRequestBody() BuysellbookingsPostRequestBody {
	return BuysellbookingsPostRequestBody{}
}

type BuysellbookingsPostRequestBody BuySellBookingAndAttachmentRequest

func (r *BuysellbookingsPostRequest) RequestBody() *BuysellbookingsPostRequestBody {
	return &r.requestBody
}

func (r *BuysellbookingsPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *BuysellbookingsPostRequest) SetRequestBody(body BuysellbookingsPostRequestBody) {
	r.requestBody = body
}

func (r *BuysellbookingsPostRequest) NewResponseBody() *BuysellbookingsPostResponseBody {
	return &BuysellbookingsPostResponseBody{}
}

type BuysellbookingsPostResponseBody []struct{}

func (r *BuysellbookingsPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("dossiers/{{.dossierId}}/buysellbookings", r.PathParams())
	return &u
}

func (r *BuysellbookingsPostRequest) Do() (BuysellbookingsPostResponseBody, error) {
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
