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
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// Todo defines model for Todo.
type Todo struct {
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	Id        string    `json:"id"`
	Subject   string    `json:"subject"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// UpdateTodoPartialJSONBody defines parameters for UpdateTodoPartial.
type UpdateTodoPartialJSONBody struct {
	Body    *string `json:"body,omitempty"`
	Subject *string `json:"subject,omitempty"`
}

// PostTodosJSONRequestBody defines body for PostTodos for application/json ContentType.
type PostTodosJSONRequestBody = Todo

// UpdateTodoPartialJSONRequestBody defines body for UpdateTodoPartial for application/json ContentType.
type UpdateTodoPartialJSONRequestBody UpdateTodoPartialJSONBody

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
	// GetTodos request
	GetTodos(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostTodosWithBody request with any body
	PostTodosWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostTodos(ctx context.Context, body PostTodosJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteTodo request
	DeleteTodo(ctx context.Context, todoId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateTodoPartialWithBody request with any body
	UpdateTodoPartialWithBody(ctx context.Context, todoId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateTodoPartial(ctx context.Context, todoId string, body UpdateTodoPartialJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetTodos(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTodosRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTodosWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTodosRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTodos(ctx context.Context, body PostTodosJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTodosRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteTodo(ctx context.Context, todoId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteTodoRequest(c.Server, todoId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateTodoPartialWithBody(ctx context.Context, todoId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateTodoPartialRequestWithBody(c.Server, todoId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateTodoPartial(ctx context.Context, todoId string, body UpdateTodoPartialJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateTodoPartialRequest(c.Server, todoId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetTodosRequest generates requests for GetTodos
func NewGetTodosRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/todos")
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

// NewPostTodosRequest calls the generic PostTodos builder with application/json body
func NewPostTodosRequest(server string, body PostTodosJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostTodosRequestWithBody(server, "application/json", bodyReader)
}

// NewPostTodosRequestWithBody generates requests for PostTodos with any type of body
func NewPostTodosRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/todos")
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

// NewDeleteTodoRequest generates requests for DeleteTodo
func NewDeleteTodoRequest(server string, todoId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "todoId", runtime.ParamLocationPath, todoId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/todos/%s", pathParam0)
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

// NewUpdateTodoPartialRequest calls the generic UpdateTodoPartial builder with application/json body
func NewUpdateTodoPartialRequest(server string, todoId string, body UpdateTodoPartialJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateTodoPartialRequestWithBody(server, todoId, "application/json", bodyReader)
}

// NewUpdateTodoPartialRequestWithBody generates requests for UpdateTodoPartial with any type of body
func NewUpdateTodoPartialRequestWithBody(server string, todoId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "todoId", runtime.ParamLocationPath, todoId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/todos/%s", pathParam0)
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
	// GetTodosWithResponse request
	GetTodosWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetTodosResponse, error)

	// PostTodosWithBodyWithResponse request with any body
	PostTodosWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTodosResponse, error)

	PostTodosWithResponse(ctx context.Context, body PostTodosJSONRequestBody, reqEditors ...RequestEditorFn) (*PostTodosResponse, error)

	// DeleteTodoWithResponse request
	DeleteTodoWithResponse(ctx context.Context, todoId string, reqEditors ...RequestEditorFn) (*DeleteTodoResponse, error)

	// UpdateTodoPartialWithBodyWithResponse request with any body
	UpdateTodoPartialWithBodyWithResponse(ctx context.Context, todoId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateTodoPartialResponse, error)

	UpdateTodoPartialWithResponse(ctx context.Context, todoId string, body UpdateTodoPartialJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateTodoPartialResponse, error)
}

type GetTodosResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Todo
}

// Status returns HTTPResponse.Status
func (r GetTodosResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTodosResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostTodosResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostTodosResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostTodosResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteTodoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r DeleteTodoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteTodoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateTodoPartialResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r UpdateTodoPartialResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateTodoPartialResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetTodosWithResponse request returning *GetTodosResponse
func (c *ClientWithResponses) GetTodosWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetTodosResponse, error) {
	rsp, err := c.GetTodos(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTodosResponse(rsp)
}

// PostTodosWithBodyWithResponse request with arbitrary body returning *PostTodosResponse
func (c *ClientWithResponses) PostTodosWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTodosResponse, error) {
	rsp, err := c.PostTodosWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTodosResponse(rsp)
}

func (c *ClientWithResponses) PostTodosWithResponse(ctx context.Context, body PostTodosJSONRequestBody, reqEditors ...RequestEditorFn) (*PostTodosResponse, error) {
	rsp, err := c.PostTodos(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTodosResponse(rsp)
}

// DeleteTodoWithResponse request returning *DeleteTodoResponse
func (c *ClientWithResponses) DeleteTodoWithResponse(ctx context.Context, todoId string, reqEditors ...RequestEditorFn) (*DeleteTodoResponse, error) {
	rsp, err := c.DeleteTodo(ctx, todoId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteTodoResponse(rsp)
}

// UpdateTodoPartialWithBodyWithResponse request with arbitrary body returning *UpdateTodoPartialResponse
func (c *ClientWithResponses) UpdateTodoPartialWithBodyWithResponse(ctx context.Context, todoId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateTodoPartialResponse, error) {
	rsp, err := c.UpdateTodoPartialWithBody(ctx, todoId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateTodoPartialResponse(rsp)
}

func (c *ClientWithResponses) UpdateTodoPartialWithResponse(ctx context.Context, todoId string, body UpdateTodoPartialJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateTodoPartialResponse, error) {
	rsp, err := c.UpdateTodoPartial(ctx, todoId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateTodoPartialResponse(rsp)
}

// ParseGetTodosResponse parses an HTTP response from a GetTodosWithResponse call
func ParseGetTodosResponse(rsp *http.Response) (*GetTodosResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetTodosResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Todo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePostTodosResponse parses an HTTP response from a PostTodosWithResponse call
func ParsePostTodosResponse(rsp *http.Response) (*PostTodosResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostTodosResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseDeleteTodoResponse parses an HTTP response from a DeleteTodoWithResponse call
func ParseDeleteTodoResponse(rsp *http.Response) (*DeleteTodoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteTodoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseUpdateTodoPartialResponse parses an HTTP response from a UpdateTodoPartialWithResponse call
func ParseUpdateTodoPartialResponse(rsp *http.Response) (*UpdateTodoPartialResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateTodoPartialResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// TODOのリストを取得
	// (GET /todos)
	GetTodos(ctx echo.Context) error
	// 新しいTODOを追加
	// (POST /todos)
	PostTodos(ctx echo.Context) error
	// 指定されたIDのTODOを削除
	// (DELETE /todos/{todoId})
	DeleteTodo(ctx echo.Context, todoId string) error
	// 指定されたIDのTODOの部分的な情報を更新
	// (PATCH /todos/{todoId})
	UpdateTodoPartial(ctx echo.Context, todoId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTodos converts echo context to params.
func (w *ServerInterfaceWrapper) GetTodos(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTodos(ctx)
	return err
}

// PostTodos converts echo context to params.
func (w *ServerInterfaceWrapper) PostTodos(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTodos(ctx)
	return err
}

// DeleteTodo converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTodo(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "todoId" -------------
	var todoId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "todoId", runtime.ParamLocationPath, ctx.Param("todoId"), &todoId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter todoId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTodo(ctx, todoId)
	return err
}

// UpdateTodoPartial converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateTodoPartial(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "todoId" -------------
	var todoId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "todoId", runtime.ParamLocationPath, ctx.Param("todoId"), &todoId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter todoId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateTodoPartial(ctx, todoId)
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

	router.GET(baseURL+"/todos", wrapper.GetTodos)
	router.POST(baseURL+"/todos", wrapper.PostTodos)
	router.DELETE(baseURL+"/todos/:todoId", wrapper.DeleteTodo)
	router.PATCH(baseURL+"/todos/:todoId", wrapper.UpdateTodoPartial)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7yUQU8UMRTHvwp5ehzZRTnNDbOJ2RMk4olwKDMPKNmZlrZrsiFz6CwqLptoMEKAAyQS",
	"NRg3mnjAg/hhmlnhW5h2Flh2hyya6GU6afv63v/3f+0aBCziLMZYSfDXQAbLGBH3O8tCZkcuGEehKLrZ",
	"BRY27KgaHMEHqQSNlyDxIBBIFIZTyq4uMhERBT6EROE9RSMEbziEhoUnyfrCCgaqcK3Owz/LknggcLVO",
	"BYbgz9mUVwm8XE1/7f0Z5i8PY/n+xJ5G40WHRVFVs2uPacRrOGZpjU3NVMGDpygkZTH4MDFeHi/bshnH",
	"mHAKPjxwUx5wopYd0JJiIXN/S+hUWdpEURZXQ/DhEapZt8HqkJzFMvfhfrlsh4DFCmMXRziv0cBFllak",
	"zX9hp/2jCiMXeFfgIvhwp3RlfKnneslZnlzKJkKQRq46RBkIylWuq7vxOmsdOLiyHkVENMCH2enKtNEd",
	"0zw26XfT3DDpVvZqOzvdsSdyJgvEzTDZp261jlI97PXXrYWN1nO9B5SoYzIEc8IO10U6PenW2c8fWevQ",
	"6B2jT913UHd3+4ubX78W4Dbl3pbW7FANkzxHDRUOo6i4eVewbQ5BIlQoJPhzg3VlL1vnu0dG75p0s8e8",
	"WgHbmOC7tgIPYhJZA/O8MCjf68M3eF3mh9BMDqOxjk4WLeT1tM/ebxp9ZPSmSVsO275J3xj9oZhf+0XW",
	"2TP6rUnbRh9UK0Z3eihzqa5/iAqWh6k9cZfVUpshQlFSGwWvu//N+fUP4f1dG9/ylb35bUwKX6tRfV++",
	"ycNOt/ksO/xqdLv7+V12cmL0pwt2uVGXl+F/tILunDc/ZhvPf+2tG33cKy3dyiuySpPfAQAA//814hDU",
	"yQYAAA==",
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
