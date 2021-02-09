package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewDossiersGetRequest() DossiersGetRequest {
	r := DossiersGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DossiersGetRequest struct {
	client      *Client
	queryParams *DossiersGetQueryParams
	pathParams  *DossiersGetPathParams
	method      string
	headers     http.Header
	requestBody DossiersGetRequestBody
}

func (r DossiersGetRequest) NewQueryParams() *DossiersGetQueryParams {
	return &DossiersGetQueryParams{}
}

type DossiersGetQueryParams struct {
}

func (p DossiersGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DossiersGetRequest) QueryParams() *DossiersGetQueryParams {
	return r.queryParams
}

func (r DossiersGetRequest) NewPathParams() *DossiersGetPathParams {
	return &DossiersGetPathParams{}
}

type DossiersGetPathParams struct {
}

func (p *DossiersGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DossiersGetRequest) PathParams() *DossiersGetPathParams {
	return r.pathParams
}

func (r *DossiersGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DossiersGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *DossiersGetRequest) Method() string {
	return r.method
}

func (r DossiersGetRequest) NewRequestBody() DossiersGetRequestBody {
	return DossiersGetRequestBody{}
}

type DossiersGetRequestBody struct {
}

func (r *DossiersGetRequest) RequestBody() *DossiersGetRequestBody {
	return &r.requestBody
}

func (r *DossiersGetRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *DossiersGetRequest) SetRequestBody(body DossiersGetRequestBody) {
	r.requestBody = body
}

func (r *DossiersGetRequest) NewResponseBody() *DossiersGetResponseBody {
	return &DossiersGetResponseBody{}
}

type DossiersGetResponseBody []struct {
	City               string `json:"city"`
	Country            string `json:"country"`
	DossierDescription string `json:"dossierDescription"`
	DossierKey         struct {
		ID int `json:"id"`
	} `json:"dossierKey"`
	Email       string `json:"email"`
	PostalCode  string `json:"postalCode"`
	StreetAndNr string `json:"streetAndNr"`
	URL         string `json:"url"`
	VatNr       string `json:"vatNr"`
}

func (r *DossiersGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("dossiers", r.PathParams())
	return &u
}

func (r *DossiersGetRequest) Do() (DossiersGetResponseBody, error) {
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
