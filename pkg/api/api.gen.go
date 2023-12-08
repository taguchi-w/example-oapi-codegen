// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// Pet defines model for Pet.
type Pet struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Tag  *string `json:"tag,omitempty"`
}

// UpdatePetPartialJSONBody defines parameters for UpdatePetPartial.
type UpdatePetPartialJSONBody struct {
	Name *string `json:"name,omitempty"`
	Tag  *string `json:"tag,omitempty"`
}

// PostPetsJSONRequestBody defines body for PostPets for application/json ContentType.
type PostPetsJSONRequestBody = Pet

// UpdatePetPartialJSONRequestBody defines body for UpdatePetPartial for application/json ContentType.
type UpdatePetPartialJSONRequestBody UpdatePetPartialJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetPets request
	GetPets(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostPetsWithBody request with any body
	PostPetsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostPets(ctx context.Context, body PostPetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeletePet request
	DeletePet(ctx context.Context, petId int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdatePetPartialWithBody request with any body
	UpdatePetPartialWithBody(ctx context.Context, petId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdatePetPartial(ctx context.Context, petId int, body UpdatePetPartialJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetPets(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPetsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostPetsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostPetsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostPets(ctx context.Context, body PostPetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostPetsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeletePet(ctx context.Context, petId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeletePetRequest(c.Server, petId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdatePetPartialWithBody(ctx context.Context, petId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdatePetPartialRequestWithBody(c.Server, petId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdatePetPartial(ctx context.Context, petId int, body UpdatePetPartialJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdatePetPartialRequest(c.Server, petId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetPetsRequest generates requests for GetPets
func NewGetPetsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostPetsRequest calls the generic PostPets builder with application/json body
func NewPostPetsRequest(server string, body PostPetsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostPetsRequestWithBody(server, "application/json", bodyReader)
}

// NewPostPetsRequestWithBody generates requests for PostPets with any type of body
func NewPostPetsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeletePetRequest generates requests for DeletePet
func NewDeletePetRequest(server string, petId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "petId", runtime.ParamLocationPath, petId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdatePetPartialRequest calls the generic UpdatePetPartial builder with application/json body
func NewUpdatePetPartialRequest(server string, petId int, body UpdatePetPartialJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdatePetPartialRequestWithBody(server, petId, "application/json", bodyReader)
}

// NewUpdatePetPartialRequestWithBody generates requests for UpdatePetPartial with any type of body
func NewUpdatePetPartialRequestWithBody(server string, petId int, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "petId", runtime.ParamLocationPath, petId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetPetsWithResponse request
	GetPetsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetPetsResponse, error)

	// PostPetsWithBodyWithResponse request with any body
	PostPetsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostPetsResponse, error)

	PostPetsWithResponse(ctx context.Context, body PostPetsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostPetsResponse, error)

	// DeletePetWithResponse request
	DeletePetWithResponse(ctx context.Context, petId int, reqEditors ...RequestEditorFn) (*DeletePetResponse, error)

	// UpdatePetPartialWithBodyWithResponse request with any body
	UpdatePetPartialWithBodyWithResponse(ctx context.Context, petId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdatePetPartialResponse, error)

	UpdatePetPartialWithResponse(ctx context.Context, petId int, body UpdatePetPartialJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdatePetPartialResponse, error)
}

type GetPetsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Pet
}

// Status returns HTTPResponse.Status
func (r GetPetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostPetsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostPetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostPetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeletePetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeletePetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeletePetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdatePetPartialResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r UpdatePetPartialResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdatePetPartialResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetPetsWithResponse request returning *GetPetsResponse
func (c *ClientWithResponses) GetPetsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetPetsResponse, error) {
	rsp, err := c.GetPets(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPetsResponse(rsp)
}

// PostPetsWithBodyWithResponse request with arbitrary body returning *PostPetsResponse
func (c *ClientWithResponses) PostPetsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostPetsResponse, error) {
	rsp, err := c.PostPetsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostPetsResponse(rsp)
}

func (c *ClientWithResponses) PostPetsWithResponse(ctx context.Context, body PostPetsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostPetsResponse, error) {
	rsp, err := c.PostPets(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostPetsResponse(rsp)
}

// DeletePetWithResponse request returning *DeletePetResponse
func (c *ClientWithResponses) DeletePetWithResponse(ctx context.Context, petId int, reqEditors ...RequestEditorFn) (*DeletePetResponse, error) {
	rsp, err := c.DeletePet(ctx, petId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeletePetResponse(rsp)
}

// UpdatePetPartialWithBodyWithResponse request with arbitrary body returning *UpdatePetPartialResponse
func (c *ClientWithResponses) UpdatePetPartialWithBodyWithResponse(ctx context.Context, petId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdatePetPartialResponse, error) {
	rsp, err := c.UpdatePetPartialWithBody(ctx, petId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdatePetPartialResponse(rsp)
}

func (c *ClientWithResponses) UpdatePetPartialWithResponse(ctx context.Context, petId int, body UpdatePetPartialJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdatePetPartialResponse, error) {
	rsp, err := c.UpdatePetPartial(ctx, petId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdatePetPartialResponse(rsp)
}

// ParseGetPetsResponse parses an HTTP response from a GetPetsWithResponse call
func ParseGetPetsResponse(rsp *http.Response) (*GetPetsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePostPetsResponse parses an HTTP response from a PostPetsWithResponse call
func ParsePostPetsResponse(rsp *http.Response) (*PostPetsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostPetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseDeletePetResponse parses an HTTP response from a DeletePetWithResponse call
func ParseDeletePetResponse(rsp *http.Response) (*DeletePetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeletePetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseUpdatePetPartialResponse parses an HTTP response from a UpdatePetPartialWithResponse call
func ParseUpdatePetPartialResponse(rsp *http.Response) (*UpdatePetPartialResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdatePetPartialResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// ペットのリストを取得
	// (GET /pets)
	GetPets(ctx echo.Context) error
	// 新しいペットを追加
	// (POST /pets)
	PostPets(ctx echo.Context) error
	// 指定されたIDのペットを削除
	// (DELETE /pets/{petId})
	DeletePet(ctx echo.Context, petId int) error
	// 指定されたIDのペットの部分的な情報を更新
	// (PATCH /pets/{petId})
	UpdatePetPartial(ctx echo.Context, petId int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetPets converts echo context to params.
func (w *ServerInterfaceWrapper) GetPets(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPets(ctx)
	return err
}

// PostPets converts echo context to params.
func (w *ServerInterfaceWrapper) PostPets(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostPets(ctx)
	return err
}

// DeletePet converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "petId", runtime.ParamLocationPath, ctx.Param("petId"), &petId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeletePet(ctx, petId)
	return err
}

// UpdatePetPartial converts echo context to params.
func (w *ServerInterfaceWrapper) UpdatePetPartial(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "petId", runtime.ParamLocationPath, ctx.Param("petId"), &petId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdatePetPartial(ctx, petId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/pets", wrapper.GetPets)
	router.POST(baseURL+"/pets", wrapper.PostPets)
	router.DELETE(baseURL+"/pets/:petId", wrapper.DeletePet)
	router.PATCH(baseURL+"/pets/:petId", wrapper.UpdatePetPartial)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8SUT2sTTxjH30p5fr/j0qTa096UguQWEE+lh3H3aToluzOdeSKEsofdrVrTglKxpe2h",
	"BYtKxaDgoR6sL2bY2L4LmdnGpEloqCBedpZn5vnz/TzPzDoEIpIixpg0+OuggxWMmPutI9lFKiFREUdn",
	"5KH9Ulsi+MBjwgYqSDyIWYRDO5oUjxt2g1hjgj3xQOFaiysMwV+0Ua9CLHn9o+LxKgYEiT3L42XhonBq",
	"2r2HPJJNnKkjaRIKZ+7Va+DBE1Saixh8mJutzlZtdiExZpKDD3edyQPJaMUpqUgsJTdKnVYlIy7iWgg+",
	"PECy0cHWqaWIdSn/TrVql0DEhLFzY1I2eeAcK6vaZu9DdLgII+f4v8Jl8OG/ygB35Yp1xYJOfutmSrF2",
	"KTtEHSguqRTV23xVdI4cO92KIqba4IPJD0yem3zTpF2Tn5rsm/3PdoqXu8X5ng0rhZ6gry70QOBaCzXd",
	"F2H7VtqmSrreZVItTMZwztnlus6BpGzn4sf3onNs0j2TnrvvqP7e7mdn3xj3ciddlyvrEqkWJmWuJhKO",
	"A1lwdlu3nRHFIiRUGvzF0fKKF53L/ROT7ptsa5h+bQHsoILvRqw/zz641DBKwhsCOXqbkqUxTPM3YUq3",
	"e5/eFmdnJv3YL+6NybaHkHkwPy3Exbstk56YdMtkHed5aLLXJn0/mfr286J70E9zVFtw0zdoQFmGmz5G",
	"wco47UcyZI52nSnirDkNeu/wq+v034b+Zzfh+hN5+6dwwos37d5Ub+xmt5c/LY6/DE9Gn+A/nYy0e5l/",
	"KDaf/TzYMOnpVZHZTlmbFZ78CgAA//8hvB5RkAYAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
