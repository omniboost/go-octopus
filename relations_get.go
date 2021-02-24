package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewRelationsGetRequest() RelationsGetRequest {
	r := RelationsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type RelationsGetRequest struct {
	client      *Client
	queryParams *RelationsGetQueryParams
	pathParams  *RelationsGetPathParams
	method      string
	headers     http.Header
	requestBody RelationsGetRequestBody
}

func (r RelationsGetRequest) NewQueryParams() *RelationsGetQueryParams {
	return &RelationsGetQueryParams{}
}

type RelationsGetQueryParams struct {
	RelationID    int    `schema:"relationId,omitempty"`
	ExtRelationID int    `schema:"extRelationId,omitempty"`
	Name          string `schema:"name,omitempty"`
	VATNr         string `schema:"vatNr,omitempty"`
}

func (p RelationsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *RelationsGetRequest) QueryParams() *RelationsGetQueryParams {
	return r.queryParams
}

func (r RelationsGetRequest) NewPathParams() *RelationsGetPathParams {
	return &RelationsGetPathParams{}
}

type RelationsGetPathParams struct {
	DossierID string `schema:"dossierId"`
}

func (p *RelationsGetPathParams) Params() map[string]string {
	return map[string]string{
		"dossierId": p.DossierID,
	}
}

func (r *RelationsGetRequest) PathParams() *RelationsGetPathParams {
	return r.pathParams
}

func (r *RelationsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *RelationsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *RelationsGetRequest) Method() string {
	return r.method
}

func (r RelationsGetRequest) NewRequestBody() RelationsGetRequestBody {
	return RelationsGetRequestBody{}
}

type RelationsGetRequestBody struct {
}

func (r *RelationsGetRequest) RequestBody() *RelationsGetRequestBody {
	return &r.requestBody
}

func (r *RelationsGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *RelationsGetRequest) SetRequestBody(body RelationsGetRequestBody) {
	r.requestBody = body
}

func (r *RelationsGetRequest) NewResponseBody() *RelationsGetResponseBody {
	return &RelationsGetResponseBody{}
}

type RelationsGetResponseBody RelationServiceData

func (r *RelationsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("dossiers/{{.dossierId}}/relations", r.PathParams())
	return &u
}

func (r *RelationsGetRequest) Do() (RelationsGetResponseBody, error) {
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
