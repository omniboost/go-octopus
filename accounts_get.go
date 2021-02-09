package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewAccountsGetRequest() AccountsGetRequest {
	r := AccountsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountsGetRequest struct {
	client      *Client
	queryParams *AccountsGetQueryParams
	pathParams  *AccountsGetPathParams
	method      string
	headers     http.Header
	requestBody AccountsGetRequestBody
}

func (r AccountsGetRequest) NewQueryParams() *AccountsGetQueryParams {
	return &AccountsGetQueryParams{}
}

type AccountsGetQueryParams struct {
	BookYearID string `schema:"bookyearId"`
}

func (p AccountsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AccountsGetRequest) QueryParams() *AccountsGetQueryParams {
	return r.queryParams
}

func (r AccountsGetRequest) NewPathParams() *AccountsGetPathParams {
	return &AccountsGetPathParams{}
}

type AccountsGetPathParams struct {
	DossierID string `schema:"dossierId"`
}

func (p *AccountsGetPathParams) Params() map[string]string {
	return map[string]string{
		"dossierId": p.DossierID,
	}
}

func (r *AccountsGetRequest) PathParams() *AccountsGetPathParams {
	return r.pathParams
}

func (r *AccountsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountsGetRequest) Method() string {
	return r.method
}

func (r AccountsGetRequest) NewRequestBody() AccountsGetRequestBody {
	return AccountsGetRequestBody{}
}

type AccountsGetRequestBody struct {
	DossierID string `json:"dossierId"`
}

func (r *AccountsGetRequest) RequestBody() *AccountsGetRequestBody {
	return &r.requestBody
}

func (r *AccountsGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *AccountsGetRequest) SetRequestBody(body AccountsGetRequestBody) {
	r.requestBody = body
}

func (r *AccountsGetRequest) NewResponseBody() *AccountsGetResponseBody {
	return &AccountsGetResponseBody{}
}

type AccountsGetResponseBody []AccountServiceData

func (r *AccountsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("dossiers/{{.dossierId}}/accounts", r.PathParams())
	return &u
}

func (r *AccountsGetRequest) Do() (AccountsGetResponseBody, error) {
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
