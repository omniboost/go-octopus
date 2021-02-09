package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewBookyearsGetRequest() BookyearsGetRequest {
	r := BookyearsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BookyearsGetRequest struct {
	client      *Client
	queryParams *BookyearsGetQueryParams
	pathParams  *BookyearsGetPathParams
	method      string
	headers     http.Header
	requestBody BookyearsGetRequestBody
}

func (r BookyearsGetRequest) NewQueryParams() *BookyearsGetQueryParams {
	return &BookyearsGetQueryParams{}
}

type BookyearsGetQueryParams struct {
}

func (p BookyearsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BookyearsGetRequest) QueryParams() *BookyearsGetQueryParams {
	return r.queryParams
}

func (r BookyearsGetRequest) NewPathParams() *BookyearsGetPathParams {
	return &BookyearsGetPathParams{}
}

type BookyearsGetPathParams struct {
	DossierID string `schema:"dossierId"`
}

func (p *BookyearsGetPathParams) Params() map[string]string {
	return map[string]string{
		"dossierId": p.DossierID,
	}
}

func (r *BookyearsGetRequest) PathParams() *BookyearsGetPathParams {
	return r.pathParams
}

func (r *BookyearsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BookyearsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *BookyearsGetRequest) Method() string {
	return r.method
}

func (r BookyearsGetRequest) NewRequestBody() BookyearsGetRequestBody {
	return BookyearsGetRequestBody{}
}

type BookyearsGetRequestBody struct {
	DossierID string `json:"dossierId"`
}

func (r *BookyearsGetRequest) RequestBody() *BookyearsGetRequestBody {
	return &r.requestBody
}

func (r *BookyearsGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *BookyearsGetRequest) SetRequestBody(body BookyearsGetRequestBody) {
	r.requestBody = body
}

func (r *BookyearsGetRequest) NewResponseBody() *BookyearsGetResponseBody {
	return &BookyearsGetResponseBody{}
}

type BookyearsGetResponseBody []BookYearServiceData

func (r *BookyearsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("dossiers/{{.dossierId}}/bookyears", r.PathParams())
	return &u
}

func (r *BookyearsGetRequest) Do() (BookyearsGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), errors.WithStack(err)
	}

	dossierToken, err := r.client.DossierToken(r.PathParams().DossierID)
	if err != nil {
		return *r.NewResponseBody(), errors.WithStack(err)
	}
	req.Header.Add("DossierToken", dossierToken)

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), errors.WithStack(err)
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, errors.WithStack(err)
}
