package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewJournalsGetRequest() JournalsGetRequest {
	r := JournalsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalsGetRequest struct {
	client      *Client
	queryParams *JournalsGetQueryParams
	pathParams  *JournalsGetPathParams
	method      string
	headers     http.Header
	requestBody JournalsGetRequestBody
}

func (r JournalsGetRequest) NewQueryParams() *JournalsGetQueryParams {
	return &JournalsGetQueryParams{}
}

type JournalsGetQueryParams struct {
}

func (p JournalsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalsGetRequest) QueryParams() *JournalsGetQueryParams {
	return r.queryParams
}

func (r JournalsGetRequest) NewPathParams() *JournalsGetPathParams {
	return &JournalsGetPathParams{}
}

type JournalsGetPathParams struct {
	DossierID  string `schema:"dossierId"`
	BookYearID string `schema:"bookyearId"`
}

func (p *JournalsGetPathParams) Params() map[string]string {
	return map[string]string{
		"dossierId":  p.DossierID,
		"bookyearId": p.BookYearID,
	}
}

func (r *JournalsGetRequest) PathParams() *JournalsGetPathParams {
	return r.pathParams
}

func (r *JournalsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalsGetRequest) Method() string {
	return r.method
}

func (r JournalsGetRequest) NewRequestBody() JournalsGetRequestBody {
	return JournalsGetRequestBody{}
}

type JournalsGetRequestBody struct{}

func (r *JournalsGetRequest) RequestBody() *JournalsGetRequestBody {
	return &r.requestBody
}

func (r *JournalsGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *JournalsGetRequest) SetRequestBody(body JournalsGetRequestBody) {
	r.requestBody = body
}

func (r *JournalsGetRequest) NewResponseBody() *JournalsGetResponseBody {
	return &JournalsGetResponseBody{}
}

type JournalsGetResponseBody []JournalServiceData

func (r *JournalsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("dossiers/{{.dossierId}}/bookyears/{{.bookyearId}}/journals", r.PathParams())
	return &u
}

func (r *JournalsGetRequest) Do() (JournalsGetResponseBody, error) {
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
