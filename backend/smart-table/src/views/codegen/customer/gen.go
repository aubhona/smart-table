// Package viewsLUcustomer provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package viewsLUcustomer

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// CustomerV1OrderCreateRequest defines model for CustomerV1OrderCreateRequest.
type CustomerV1OrderCreateRequest struct {
	// CustomerUUID Логин пользователя в Telegram
	CustomerUUID openapi_types.UUID `json:"customer_uuid"`

	// RoomCode Код комнаты для группового заказа
	RoomCode *string `json:"room_code,omitempty"`

	// TableID Уникальный идентификатор стола
	TableID string `json:"table_id"`
}

// CustomerV1OrderCreateResponse defines model for CustomerV1OrderCreateResponse.
type CustomerV1OrderCreateResponse struct {
	// OrderUUID Уникальный идентификатор созданного заказа
	OrderUUID openapi_types.UUID `json:"order_uuid"`
}

// CustomerV1OrderCustomerSignInRequest defines model for CustomerV1OrderCustomerSignInRequest.
type CustomerV1OrderCustomerSignInRequest struct {
	ChatID  string `json:"chat_id"`
	TgID    string `json:"tg_id"`
	TgLogin string `json:"tg_login"`
}

// CustomerV1OrderCustomerSignInResponse defines model for CustomerV1OrderCustomerSignInResponse.
type CustomerV1OrderCustomerSignInResponse struct {
	CustomerUUID openapi_types.UUID `json:"customer_uuid"`
}

// CustomerV1OrderCustomerSignUpRequest defines model for CustomerV1OrderCustomerSignUpRequest.
type CustomerV1OrderCustomerSignUpRequest struct {
	Avatar  openapi_types.File `json:"avatar"`
	ChatID  string             `json:"chat_id"`
	TgID    string             `json:"tg_id"`
	TgLogin string             `json:"tg_login"`
}

// CustomerV1OrderCustomerSignUpResponse defines model for CustomerV1OrderCustomerSignUpResponse.
type CustomerV1OrderCustomerSignUpResponse struct {
	CustomerUUID openapi_types.UUID `json:"customer_uuid"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Code Код ошибки
	Code string `json:"code"`

	// Message Описание ошибки
	Message string `json:"message"`
}

// PostCustomerV1OrderCreateJSONRequestBody defines body for PostCustomerV1OrderCreate for application/json ContentType.
type PostCustomerV1OrderCreateJSONRequestBody = CustomerV1OrderCreateRequest

// PostCustomerV1OrderCustomerSignInJSONRequestBody defines body for PostCustomerV1OrderCustomerSignIn for application/json ContentType.
type PostCustomerV1OrderCustomerSignInJSONRequestBody = CustomerV1OrderCustomerSignInRequest

// PostCustomerV1OrderCustomerSignUpMultipartRequestBody defines body for PostCustomerV1OrderCustomerSignUp for multipart/form-data ContentType.
type PostCustomerV1OrderCustomerSignUpMultipartRequestBody = CustomerV1OrderCustomerSignUpRequest

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
	// PostCustomerV1OrderCreateWithBody request with any body
	PostCustomerV1OrderCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostCustomerV1OrderCreate(ctx context.Context, body PostCustomerV1OrderCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostCustomerV1OrderCustomerSignInWithBody request with any body
	PostCustomerV1OrderCustomerSignInWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostCustomerV1OrderCustomerSignIn(ctx context.Context, body PostCustomerV1OrderCustomerSignInJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostCustomerV1OrderCustomerSignUpWithBody request with any body
	PostCustomerV1OrderCustomerSignUpWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) PostCustomerV1OrderCreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostCustomerV1OrderCreateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostCustomerV1OrderCreate(ctx context.Context, body PostCustomerV1OrderCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostCustomerV1OrderCreateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostCustomerV1OrderCustomerSignInWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostCustomerV1OrderCustomerSignInRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostCustomerV1OrderCustomerSignIn(ctx context.Context, body PostCustomerV1OrderCustomerSignInJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostCustomerV1OrderCustomerSignInRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostCustomerV1OrderCustomerSignUpWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostCustomerV1OrderCustomerSignUpRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewPostCustomerV1OrderCreateRequest calls the generic PostCustomerV1OrderCreate builder with application/json body
func NewPostCustomerV1OrderCreateRequest(server string, body PostCustomerV1OrderCreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostCustomerV1OrderCreateRequestWithBody(server, "application/json", bodyReader)
}

// NewPostCustomerV1OrderCreateRequestWithBody generates requests for PostCustomerV1OrderCreate with any type of body
func NewPostCustomerV1OrderCreateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/customer/v1/order/create")
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

// NewPostCustomerV1OrderCustomerSignInRequest calls the generic PostCustomerV1OrderCustomerSignIn builder with application/json body
func NewPostCustomerV1OrderCustomerSignInRequest(server string, body PostCustomerV1OrderCustomerSignInJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostCustomerV1OrderCustomerSignInRequestWithBody(server, "application/json", bodyReader)
}

// NewPostCustomerV1OrderCustomerSignInRequestWithBody generates requests for PostCustomerV1OrderCustomerSignIn with any type of body
func NewPostCustomerV1OrderCustomerSignInRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/customer/v1/order/customer/sign-in")
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

// NewPostCustomerV1OrderCustomerSignUpRequestWithBody generates requests for PostCustomerV1OrderCustomerSignUp with any type of body
func NewPostCustomerV1OrderCustomerSignUpRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/customer/v1/order/customer/sign-up")
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
	// PostCustomerV1OrderCreateWithBodyWithResponse request with any body
	PostCustomerV1OrderCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostCustomerV1OrderCreateResponse, error)

	PostCustomerV1OrderCreateWithResponse(ctx context.Context, body PostCustomerV1OrderCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*PostCustomerV1OrderCreateResponse, error)

	// PostCustomerV1OrderCustomerSignInWithBodyWithResponse request with any body
	PostCustomerV1OrderCustomerSignInWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostCustomerV1OrderCustomerSignInResponse, error)

	PostCustomerV1OrderCustomerSignInWithResponse(ctx context.Context, body PostCustomerV1OrderCustomerSignInJSONRequestBody, reqEditors ...RequestEditorFn) (*PostCustomerV1OrderCustomerSignInResponse, error)

	// PostCustomerV1OrderCustomerSignUpWithBodyWithResponse request with any body
	PostCustomerV1OrderCustomerSignUpWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostCustomerV1OrderCustomerSignUpResponse, error)
}

type PostCustomerV1OrderCreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CustomerV1OrderCreateResponse
	JSON403      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r PostCustomerV1OrderCreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostCustomerV1OrderCreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostCustomerV1OrderCustomerSignInResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CustomerV1OrderCustomerSignInResponse
	JSON403      *ErrorResponse
	JSON404      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r PostCustomerV1OrderCustomerSignInResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostCustomerV1OrderCustomerSignInResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostCustomerV1OrderCustomerSignUpResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CustomerV1OrderCustomerSignUpResponse
	JSON409      *ErrorResponse
}

// Status returns HTTPResponse.Status
func (r PostCustomerV1OrderCustomerSignUpResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostCustomerV1OrderCustomerSignUpResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostCustomerV1OrderCreateWithBodyWithResponse request with arbitrary body returning *PostCustomerV1OrderCreateResponse
func (c *ClientWithResponses) PostCustomerV1OrderCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostCustomerV1OrderCreateResponse, error) {
	rsp, err := c.PostCustomerV1OrderCreateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostCustomerV1OrderCreateResponse(rsp)
}

func (c *ClientWithResponses) PostCustomerV1OrderCreateWithResponse(ctx context.Context, body PostCustomerV1OrderCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*PostCustomerV1OrderCreateResponse, error) {
	rsp, err := c.PostCustomerV1OrderCreate(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostCustomerV1OrderCreateResponse(rsp)
}

// PostCustomerV1OrderCustomerSignInWithBodyWithResponse request with arbitrary body returning *PostCustomerV1OrderCustomerSignInResponse
func (c *ClientWithResponses) PostCustomerV1OrderCustomerSignInWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostCustomerV1OrderCustomerSignInResponse, error) {
	rsp, err := c.PostCustomerV1OrderCustomerSignInWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostCustomerV1OrderCustomerSignInResponse(rsp)
}

func (c *ClientWithResponses) PostCustomerV1OrderCustomerSignInWithResponse(ctx context.Context, body PostCustomerV1OrderCustomerSignInJSONRequestBody, reqEditors ...RequestEditorFn) (*PostCustomerV1OrderCustomerSignInResponse, error) {
	rsp, err := c.PostCustomerV1OrderCustomerSignIn(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostCustomerV1OrderCustomerSignInResponse(rsp)
}

// PostCustomerV1OrderCustomerSignUpWithBodyWithResponse request with arbitrary body returning *PostCustomerV1OrderCustomerSignUpResponse
func (c *ClientWithResponses) PostCustomerV1OrderCustomerSignUpWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostCustomerV1OrderCustomerSignUpResponse, error) {
	rsp, err := c.PostCustomerV1OrderCustomerSignUpWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostCustomerV1OrderCustomerSignUpResponse(rsp)
}

// ParsePostCustomerV1OrderCreateResponse parses an HTTP response from a PostCustomerV1OrderCreateWithResponse call
func ParsePostCustomerV1OrderCreateResponse(rsp *http.Response) (*PostCustomerV1OrderCreateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostCustomerV1OrderCreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CustomerV1OrderCreateResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	}

	return response, nil
}

// ParsePostCustomerV1OrderCustomerSignInResponse parses an HTTP response from a PostCustomerV1OrderCustomerSignInWithResponse call
func ParsePostCustomerV1OrderCustomerSignInResponse(rsp *http.Response) (*PostCustomerV1OrderCustomerSignInResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostCustomerV1OrderCustomerSignInResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CustomerV1OrderCustomerSignInResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParsePostCustomerV1OrderCustomerSignUpResponse parses an HTTP response from a PostCustomerV1OrderCustomerSignUpWithResponse call
func ParsePostCustomerV1OrderCustomerSignUpResponse(rsp *http.Response) (*PostCustomerV1OrderCustomerSignUpResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostCustomerV1OrderCustomerSignUpResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CustomerV1OrderCustomerSignUpResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest ErrorResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Создание заказа, если первый пользователь в группе, иначе присоединение к заказу.
	// (POST /customer/v1/order/create)
	PostCustomerV1OrderCreate(c *gin.Context)
	// Авторизация пользователя в приложении
	// (POST /customer/v1/order/customer/sign-in)
	PostCustomerV1OrderCustomerSignIn(c *gin.Context)
	// Регистрация пользователя в приложении
	// (POST /customer/v1/order/customer/sign-up)
	PostCustomerV1OrderCustomerSignUp(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// PostCustomerV1OrderCreate operation middleware
func (siw *ServerInterfaceWrapper) PostCustomerV1OrderCreate(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostCustomerV1OrderCreate(c)
}

// PostCustomerV1OrderCustomerSignIn operation middleware
func (siw *ServerInterfaceWrapper) PostCustomerV1OrderCustomerSignIn(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostCustomerV1OrderCustomerSignIn(c)
}

// PostCustomerV1OrderCustomerSignUp operation middleware
func (siw *ServerInterfaceWrapper) PostCustomerV1OrderCustomerSignUp(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostCustomerV1OrderCustomerSignUp(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/customer/v1/order/create", wrapper.PostCustomerV1OrderCreate)
	router.POST(options.BaseURL+"/customer/v1/order/customer/sign-in", wrapper.PostCustomerV1OrderCustomerSignIn)
	router.POST(options.BaseURL+"/customer/v1/order/customer/sign-up", wrapper.PostCustomerV1OrderCustomerSignUp)
}

type PostCustomerV1OrderCreateRequestObject struct {
	Body *PostCustomerV1OrderCreateJSONRequestBody
}

type PostCustomerV1OrderCreateResponseObject interface {
	VisitPostCustomerV1OrderCreateResponse(w http.ResponseWriter) error
}

type PostCustomerV1OrderCreate200JSONResponse CustomerV1OrderCreateResponse

func (response PostCustomerV1OrderCreate200JSONResponse) VisitPostCustomerV1OrderCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostCustomerV1OrderCreate403JSONResponse ErrorResponse

func (response PostCustomerV1OrderCreate403JSONResponse) VisitPostCustomerV1OrderCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type PostCustomerV1OrderCustomerSignInRequestObject struct {
	Body *PostCustomerV1OrderCustomerSignInJSONRequestBody
}

type PostCustomerV1OrderCustomerSignInResponseObject interface {
	VisitPostCustomerV1OrderCustomerSignInResponse(w http.ResponseWriter) error
}

type PostCustomerV1OrderCustomerSignIn200JSONResponse CustomerV1OrderCustomerSignInResponse

func (response PostCustomerV1OrderCustomerSignIn200JSONResponse) VisitPostCustomerV1OrderCustomerSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostCustomerV1OrderCustomerSignIn403JSONResponse ErrorResponse

func (response PostCustomerV1OrderCustomerSignIn403JSONResponse) VisitPostCustomerV1OrderCustomerSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type PostCustomerV1OrderCustomerSignIn404JSONResponse ErrorResponse

func (response PostCustomerV1OrderCustomerSignIn404JSONResponse) VisitPostCustomerV1OrderCustomerSignInResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type PostCustomerV1OrderCustomerSignUpRequestObject struct {
	Body *multipart.Reader
}

type PostCustomerV1OrderCustomerSignUpResponseObject interface {
	VisitPostCustomerV1OrderCustomerSignUpResponse(w http.ResponseWriter) error
}

type PostCustomerV1OrderCustomerSignUp200JSONResponse CustomerV1OrderCustomerSignUpResponse

func (response PostCustomerV1OrderCustomerSignUp200JSONResponse) VisitPostCustomerV1OrderCustomerSignUpResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostCustomerV1OrderCustomerSignUp409JSONResponse ErrorResponse

func (response PostCustomerV1OrderCustomerSignUp409JSONResponse) VisitPostCustomerV1OrderCustomerSignUpResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(409)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Создание заказа, если первый пользователь в группе, иначе присоединение к заказу.
	// (POST /customer/v1/order/create)
	PostCustomerV1OrderCreate(ctx context.Context, request PostCustomerV1OrderCreateRequestObject) (PostCustomerV1OrderCreateResponseObject, error)
	// Авторизация пользователя в приложении
	// (POST /customer/v1/order/customer/sign-in)
	PostCustomerV1OrderCustomerSignIn(ctx context.Context, request PostCustomerV1OrderCustomerSignInRequestObject) (PostCustomerV1OrderCustomerSignInResponseObject, error)
	// Регистрация пользователя в приложении
	// (POST /customer/v1/order/customer/sign-up)
	PostCustomerV1OrderCustomerSignUp(ctx context.Context, request PostCustomerV1OrderCustomerSignUpRequestObject) (PostCustomerV1OrderCustomerSignUpResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// PostCustomerV1OrderCreate operation middleware
func (sh *strictHandler) PostCustomerV1OrderCreate(ctx *gin.Context) {
	var request PostCustomerV1OrderCreateRequestObject

	var body PostCustomerV1OrderCreateJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostCustomerV1OrderCreate(ctx, request.(PostCustomerV1OrderCreateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostCustomerV1OrderCreate")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostCustomerV1OrderCreateResponseObject); ok {
		if err := validResponse.VisitPostCustomerV1OrderCreateResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostCustomerV1OrderCustomerSignIn operation middleware
func (sh *strictHandler) PostCustomerV1OrderCustomerSignIn(ctx *gin.Context) {
	var request PostCustomerV1OrderCustomerSignInRequestObject

	var body PostCustomerV1OrderCustomerSignInJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostCustomerV1OrderCustomerSignIn(ctx, request.(PostCustomerV1OrderCustomerSignInRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostCustomerV1OrderCustomerSignIn")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostCustomerV1OrderCustomerSignInResponseObject); ok {
		if err := validResponse.VisitPostCustomerV1OrderCustomerSignInResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostCustomerV1OrderCustomerSignUp operation middleware
func (sh *strictHandler) PostCustomerV1OrderCustomerSignUp(ctx *gin.Context) {
	var request PostCustomerV1OrderCustomerSignUpRequestObject

	if reader, err := ctx.Request.MultipartReader(); err == nil {
		request.Body = reader
	} else {
		ctx.Error(err)
		return
	}

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostCustomerV1OrderCustomerSignUp(ctx, request.(PostCustomerV1OrderCustomerSignUpRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostCustomerV1OrderCustomerSignUp")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostCustomerV1OrderCustomerSignUpResponseObject); ok {
		if err := validResponse.VisitPostCustomerV1OrderCustomerSignUpResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RXzW7bRhB+FWLbIyPJdS7VLQ168KGo0SS9FIWxFtcyA5HLLFcGDMOAbBdtgARoeirQ",
	"okGKvoCqiDVjRfQrfPtGxSz1Q0mUYgtu7J5Ec70z38w338zwiDVkEMlQhDpm9SMWN/ZFwO3jw3asZSDU",
	"txtfK0+oh0pwLb4Rz9oi1nTOPc/Xvgx5a1vJSCjti5jV93grFi6LCq+OWGNkaqfd9j164Ym4ofyIrrM6",
	"w+/I8BYphg4ukWFgXuIcGXromlMkGJifHfScx6IlmooHzGV7UgVcszqzBl2mDyPB6izWyg+b7NhlSspg",
	"pyE9UeLtN2ToO7hAhvcYkg/zwkE/d/PWdMwZLi2OnoWVOThHFxfo0m+ZM813W2KnNLK/MERq71JQQ/MC",
	"7xyk6CPB0JwiNT/kx+YUmek45oQeMCjzQ1GJZ21fCY/Vv5s6defS+/3kptx9KhqaEC7hMo5kGItrkinJ",
	"wjImrx0vMpyjjy6GGJan+wNcz6WlgO4qiRj9+chvhlvhmsW9z/WI/MXKaK44acmmH5YczhPdzFme3HAn",
	"LteIcC3KF/R7PU6uX58F0E+i9WjhB1xzNQN21w+5OixT8C1y6I6RrpGXu0nml0pJtS64lS07M8+R4m9c",
	"IC1jMRBxzJtl91/jEqk5sW0mRfIBS/MRE6ap9cWQ6YIf7slFxw+2t8aThaaK6aCLHgbUDZHSuCm0OrxH",
	"WiEwvm6R9UcBV/oxdXnnK7nrt4TzYHuLuexAqDi3vlGpVWoUuIxEyCOf1dlmpVbZZC6LuN63Ca2O+aoe",
	"bFRtc6w2bPenw0jmsprL1p/jpmx+MaeO7cu9vJNP4DLrVHG6suWxOtuWsS6dMixPpoj1F9I7zDkOtQhz",
	"PUdRy29YK9WnsQynKwg9farEHquzT6rTHaU6WlCqK7eT41kKtWoL+yIvSZuXz2q1/xrLSAAWzFyGfx3n",
	"0TFn5gSXSMxzyvPMQCRm79c2bwzmrC7LYL2eqKLrUKXmUxop8W5+RIqUhvmAfvrIaFmxu1LXSiZuBwG1",
	"12IBjeVWKHPXQWJOciMUd2dcW+Wb30va/ApbWeISBNrafiLDJKnUJi1B3x4kY58XBbfmrGIxlqlh/Cb2",
	"m+G9vJkv0cWrYkbMGRJSx5KF9UoCmZnNH0kopRvPLQumfEcpq9AZCpCaTp71/5dYCOn9j4j0D5LDEF28",
	"yxfxpVqb1/GrxahoaK34RssFiQEy/DOSYno15bWjFcp7g4Q+EG0SSfCdG1bfk2il+oJ2S/sRV7pKu9I9",
	"j2t+I8U+3W3vjgALe+UaAvz8lgRoOjMVMhLgfD2/Wfyn9er5+PjfAAAA//9SPVBAOREAAA==",
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
