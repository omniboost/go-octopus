package octopus

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-octopus/utils"
)

func (c *Client) NewAuthenticationPostRequest() AuthenticationPostRequest {
	r := AuthenticationPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AuthenticationPostRequest struct {
	client      *Client
	queryParams *AuthenticationPostQueryParams
	pathParams  *AuthenticationPostPathParams
	method      string
	headers     http.Header
	requestBody AuthenticationPostRequestBody
}

func (r AuthenticationPostRequest) NewQueryParams() *AuthenticationPostQueryParams {
	return &AuthenticationPostQueryParams{}
}

type AuthenticationPostQueryParams struct {
}

func (p AuthenticationPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AuthenticationPostRequest) QueryParams() *AuthenticationPostQueryParams {
	return r.queryParams
}

func (r AuthenticationPostRequest) NewPathParams() *AuthenticationPostPathParams {
	return &AuthenticationPostPathParams{}
}

type AuthenticationPostPathParams struct {
}

func (p *AuthenticationPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AuthenticationPostRequest) PathParams() *AuthenticationPostPathParams {
	return r.pathParams
}

func (r *AuthenticationPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AuthenticationPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *AuthenticationPostRequest) Method() string {
	return r.method
}

func (r AuthenticationPostRequest) NewRequestBody() AuthenticationPostRequestBody {
	return AuthenticationPostRequestBody{}
}

type AuthenticationPostRequestBody struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (r *AuthenticationPostRequest) RequestBody() *AuthenticationPostRequestBody {
	return &r.requestBody
}

func (r *AuthenticationPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *AuthenticationPostRequest) SetRequestBody(body AuthenticationPostRequestBody) {
	r.requestBody = body
}

func (r *AuthenticationPostRequest) NewResponseBody() *AuthenticationPostResponseBody {
	return &AuthenticationPostResponseBody{}
}

type AuthenticationPostResponseBody struct {
	Token string `json:"token"`
}

func (r *AuthenticationPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("authentication", r.PathParams())
	return &u
}

func (r *AuthenticationPostRequest) Do() (AuthenticationPostResponseBody, error) {
	r.RequestBody().User = r.client.Username()
	r.RequestBody().Password = r.client.Password()

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
