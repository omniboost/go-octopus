package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewRelationsPutRequest() RelationsPutRequest {
	r := RelationsPutRequest{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type RelationsPutRequest struct {
	client      *Client
	queryParams *RelationsPutQueryParams
	pathParams  *RelationsPutPathParams
	method      string
	headers     http.Header
	requestBody RelationsPutRequestBody
}

func (r RelationsPutRequest) NewQueryParams() *RelationsPutQueryParams {
	return &RelationsPutQueryParams{}
}

type RelationsPutQueryParams struct{}

func (p RelationsPutQueryParams) ToURLValues() (url.Values, error) {
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

func (r *RelationsPutRequest) QueryParams() *RelationsPutQueryParams {
	return r.queryParams
}

func (r RelationsPutRequest) NewPathParams() *RelationsPutPathParams {
	return &RelationsPutPathParams{}
}

type RelationsPutPathParams struct {
	DossierID string `schema:"dossierId"`
}

func (p *RelationsPutPathParams) Params() map[string]string {
	return map[string]string{
		"dossierId": p.DossierID,
	}
}

func (r *RelationsPutRequest) PathParams() *RelationsPutPathParams {
	return r.pathParams
}

func (r *RelationsPutRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *RelationsPutRequest) SetMethod(method string) {
	r.method = method
}

func (r *RelationsPutRequest) Method() string {
	return r.method
}

func (r RelationsPutRequest) NewRequestBody() RelationsPutRequestBody {
	return RelationsPutRequestBody{}
}

type RelationsPutRequestBody RelationIdentificationServiceData

func (r *RelationsPutRequest) RequestBody() *RelationsPutRequestBody {
	return &r.requestBody
}

func (r *RelationsPutRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *RelationsPutRequest) SetRequestBody(body RelationsPutRequestBody) {
	r.requestBody = body
}

func (r *RelationsPutRequest) NewResponseBody() *RelationsPutResponseBody {
	return &RelationsPutResponseBody{}
}

type RelationsPutResponseBody []struct{}

func (r *RelationsPutRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("dossiers/{{.dossierId}}/relations", r.PathParams())
	return &u
}

func (r *RelationsPutRequest) Do() (RelationsPutResponseBody, error) {
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
