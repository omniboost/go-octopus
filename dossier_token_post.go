package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewDossierTokenPostRequest() DossierTokenPostRequest {
	r := DossierTokenPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DossierTokenPostRequest struct {
	client      *Client
	queryParams *DossierTokenPostQueryParams
	pathParams  *DossierTokenPostPathParams
	method      string
	headers     http.Header
	requestBody DossierTokenPostRequestBody
}

func (r DossierTokenPostRequest) NewQueryParams() *DossierTokenPostQueryParams {
	return &DossierTokenPostQueryParams{}
}

type DossierTokenPostQueryParams struct {
	DossierID string `schema:"dossierId"`
}

func (p DossierTokenPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DossierTokenPostRequest) QueryParams() *DossierTokenPostQueryParams {
	return r.queryParams
}

func (r DossierTokenPostRequest) NewPathParams() *DossierTokenPostPathParams {
	return &DossierTokenPostPathParams{}
}

type DossierTokenPostPathParams struct {
}

func (p *DossierTokenPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DossierTokenPostRequest) PathParams() *DossierTokenPostPathParams {
	return r.pathParams
}

func (r *DossierTokenPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DossierTokenPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *DossierTokenPostRequest) Method() string {
	return r.method
}

func (r DossierTokenPostRequest) NewRequestBody() DossierTokenPostRequestBody {
	return DossierTokenPostRequestBody{}
}

type DossierTokenPostRequestBody struct {
	DossierID string `json:"dossierId"`
}

func (r *DossierTokenPostRequest) RequestBody() *DossierTokenPostRequestBody {
	return &r.requestBody
}

func (r *DossierTokenPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *DossierTokenPostRequest) SetRequestBody(body DossierTokenPostRequestBody) {
	r.requestBody = body
}

func (r *DossierTokenPostRequest) NewResponseBody() *DossierTokenPostResponseBody {
	return &DossierTokenPostResponseBody{}
}

type DossierTokenPostResponseBody struct {
	DossierToken string `json:"Dossiertoken"`
}

func (r *DossierTokenPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("dossiers", r.PathParams())
	return &u
}

func (r *DossierTokenPostRequest) Do() (DossierTokenPostResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), errors.WithStack(err)
	}

	token, err := r.client.Token()
	if err != nil {
		return *r.NewResponseBody(), errors.WithStack(err)
	}
	req.Header.Add("Token", token)

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), errors.WithStack(err)
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, errors.WithStack(err)
}
