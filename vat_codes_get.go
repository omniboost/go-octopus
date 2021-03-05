package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewVATCodesRequest() VATCodesRequest {
	r := VATCodesRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VATCodesRequest struct {
	client      *Client
	queryParams *VATCodesQueryParams
	pathParams  *VATCodesPathParams
	method      string
	headers     http.Header
	requestBody VATCodesRequestBody
}

func (r VATCodesRequest) NewQueryParams() *VATCodesQueryParams {
	return &VATCodesQueryParams{}
}

type VATCodesQueryParams struct {
}

func (p VATCodesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VATCodesRequest) QueryParams() *VATCodesQueryParams {
	return r.queryParams
}

func (r VATCodesRequest) NewPathParams() *VATCodesPathParams {
	return &VATCodesPathParams{}
}

type VATCodesPathParams struct {
	DossierID string `schema:"dossierId"`
}

func (p *VATCodesPathParams) Params() map[string]string {
	return map[string]string{
		"dossierId": p.DossierID,
	}
}

func (r *VATCodesRequest) PathParams() *VATCodesPathParams {
	return r.pathParams
}

func (r *VATCodesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VATCodesRequest) SetMethod(method string) {
	r.method = method
}

func (r *VATCodesRequest) Method() string {
	return r.method
}

func (r VATCodesRequest) NewRequestBody() VATCodesRequestBody {
	return VATCodesRequestBody{}
}

type VATCodesRequestBody struct{}

func (r *VATCodesRequest) RequestBody() *VATCodesRequestBody {
	return &r.requestBody
}

func (r *VATCodesRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *VATCodesRequest) SetRequestBody(body VATCodesRequestBody) {
	r.requestBody = body
}

func (r *VATCodesRequest) NewResponseBody() *VATCodesResponseBody {
	return &VATCodesResponseBody{}
}

type VATCodesResponseBody []VATCodeServiceData

func (r *VATCodesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("dossiers/{{.dossierId}}/vatcodes", r.PathParams())
	return &u
}

func (r *VATCodesRequest) Do() (VATCodesResponseBody, error) {
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
