// Package ServiceOapi provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package ServiceOapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"strings"
)

// Domain defines model for domain.
type Domain string

// Record defines model for record.
type Record string

// RecordType defines model for recordType.
type RecordType string

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(req *http.Request, ctx context.Context) error

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example.
	Server string

	// HTTP client with any customized settings, such as certificate chains.
	Client http.Client

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetDomains request
	GetDomains(ctx context.Context) (*http.Response, error)

	// CreateDomain request
	CreateDomain(ctx context.Context) (*http.Response, error)

	// GetDomain request
	GetDomain(ctx context.Context, domain Domain) (*http.Response, error)

	// GetRecord request
	GetRecord(ctx context.Context, domain Domain, record Record, recordType RecordType) (*http.Response, error)

	// CreateRecord request
	CreateRecord(ctx context.Context, domain Domain, record Record, recordType RecordType) (*http.Response, error)
}

func (c *Client) GetDomains(ctx context.Context) (*http.Response, error) {
	req, err := NewGetDomainsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) CreateDomain(ctx context.Context) (*http.Response, error) {
	req, err := NewCreateDomainRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetDomain(ctx context.Context, domain Domain) (*http.Response, error) {
	req, err := NewGetDomainRequest(c.Server, domain)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetRecord(ctx context.Context, domain Domain, record Record, recordType RecordType) (*http.Response, error) {
	req, err := NewGetRecordRequest(c.Server, domain, record, recordType)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) CreateRecord(ctx context.Context, domain Domain, record Record, recordType RecordType) (*http.Response, error) {
	req, err := NewCreateRecordRequest(c.Server, domain, record, recordType)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(req, ctx)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewGetDomainsRequest generates requests for GetDomains
func NewGetDomainsRequest(server string) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/domains", server)

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateDomainRequest generates requests for CreateDomain
func NewCreateDomainRequest(server string) (*http.Request, error) {
	var err error

	queryUrl := fmt.Sprintf("%s/domains", server)

	req, err := http.NewRequest("PUT", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetDomainRequest generates requests for GetDomain
func NewGetDomainRequest(server string, domain Domain) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "domain", domain)
	if err != nil {
		return nil, err
	}

	queryUrl := fmt.Sprintf("%s/domains/%s", server, pathParam0)

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetRecordRequest generates requests for GetRecord
func NewGetRecordRequest(server string, domain Domain, record Record, recordType RecordType) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "domain", domain)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParam("simple", false, "record", record)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParam("simple", false, "recordType", recordType)
	if err != nil {
		return nil, err
	}

	queryUrl := fmt.Sprintf("%s/domains/%s/%s/%s", server, pathParam0, pathParam1, pathParam2)

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateRecordRequest generates requests for CreateRecord
func NewCreateRecordRequest(server string, domain Domain, record Record, recordType RecordType) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "domain", domain)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParam("simple", false, "record", record)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParam("simple", false, "recordType", recordType)
	if err != nil {
		return nil, err
	}

	queryUrl := fmt.Sprintf("%s/domains/%s/%s/%s", server, pathParam0, pathParam1, pathParam2)

	req, err := http.NewRequest("PUT", queryUrl, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses returns a ClientWithResponses with a default Client:
func NewClientWithResponses(server string) *ClientWithResponses {
	return &ClientWithResponses{
		ClientInterface: &Client{
			Client: http.Client{},
			Server: server,
		},
	}
}

// NewClientWithResponsesAndRequestEditorFunc takes in a RequestEditorFn callback function and returns a ClientWithResponses with a default Client:
func NewClientWithResponsesAndRequestEditorFunc(server string, reqEditorFn RequestEditorFn) *ClientWithResponses {
	return &ClientWithResponses{
		ClientInterface: &Client{
			Client:        http.Client{},
			Server:        server,
			RequestEditor: reqEditorFn,
		},
	}
}

type getDomainsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]string
}

// Status returns HTTPResponse.Status
func (r getDomainsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r getDomainsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type createDomainResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r createDomainResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r createDomainResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type getDomainResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r getDomainResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r getDomainResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type getRecordResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r getRecordResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r getRecordResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type createRecordResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r createRecordResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r createRecordResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetDomainsWithResponse request returning *GetDomainsResponse
func (c *ClientWithResponses) GetDomainsWithResponse(ctx context.Context) (*getDomainsResponse, error) {
	rsp, err := c.GetDomains(ctx)
	if err != nil {
		return nil, err
	}
	return ParsegetDomainsResponse(rsp)
}

// CreateDomainWithResponse request returning *CreateDomainResponse
func (c *ClientWithResponses) CreateDomainWithResponse(ctx context.Context) (*createDomainResponse, error) {
	rsp, err := c.CreateDomain(ctx)
	if err != nil {
		return nil, err
	}
	return ParsecreateDomainResponse(rsp)
}

// GetDomainWithResponse request returning *GetDomainResponse
func (c *ClientWithResponses) GetDomainWithResponse(ctx context.Context, domain Domain) (*getDomainResponse, error) {
	rsp, err := c.GetDomain(ctx, domain)
	if err != nil {
		return nil, err
	}
	return ParsegetDomainResponse(rsp)
}

// GetRecordWithResponse request returning *GetRecordResponse
func (c *ClientWithResponses) GetRecordWithResponse(ctx context.Context, domain Domain, record Record, recordType RecordType) (*getRecordResponse, error) {
	rsp, err := c.GetRecord(ctx, domain, record, recordType)
	if err != nil {
		return nil, err
	}
	return ParsegetRecordResponse(rsp)
}

// CreateRecordWithResponse request returning *CreateRecordResponse
func (c *ClientWithResponses) CreateRecordWithResponse(ctx context.Context, domain Domain, record Record, recordType RecordType) (*createRecordResponse, error) {
	rsp, err := c.CreateRecord(ctx, domain, record, recordType)
	if err != nil {
		return nil, err
	}
	return ParsecreateRecordResponse(rsp)
}

// ParsegetDomainsResponse parses an HTTP response from a GetDomainsWithResponse call
func ParsegetDomainsResponse(rsp *http.Response) (*getDomainsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &getDomainsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		response.JSON200 = &[]string{}
		if err := json.Unmarshal(bodyBytes, response.JSON200); err != nil {
			return nil, err
		}

	}

	return response, nil
}

// ParsecreateDomainResponse parses an HTTP response from a CreateDomainWithResponse call
func ParsecreateDomainResponse(rsp *http.Response) (*createDomainResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &createDomainResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParsegetDomainResponse parses an HTTP response from a GetDomainWithResponse call
func ParsegetDomainResponse(rsp *http.Response) (*getDomainResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &getDomainResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParsegetRecordResponse parses an HTTP response from a GetRecordWithResponse call
func ParsegetRecordResponse(rsp *http.Response) (*getRecordResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &getRecordResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParsecreateRecordResponse parses an HTTP response from a CreateRecordWithResponse call
func ParsecreateRecordResponse(rsp *http.Response) (*createRecordResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &createRecordResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// get a list of registered domains// (GET /domains)
	GetDomains(ctx echo.Context) error
	// create a new domain record// (PUT /domains)
	CreateDomain(ctx echo.Context) error
	// get full info about a domain// (GET /domains/{domain})
	GetDomain(ctx echo.Context, domain Domain) error
	// get full info about a domain record// (GET /domains/{domain}/{record}/{recordType})
	GetRecord(ctx echo.Context, domain Domain, record Record, recordType RecordType) error
	// create a new domain record// (PUT /domains/{domain}/{record}/{recordType})
	CreateRecord(ctx echo.Context, domain Domain, record Record, recordType RecordType) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetDomains converts echo context to params.
func (w *ServerInterfaceWrapper) GetDomains(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDomains(ctx)
	return err
}

// CreateDomain converts echo context to params.
func (w *ServerInterfaceWrapper) CreateDomain(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateDomain(ctx)
	return err
}

// GetDomain converts echo context to params.
func (w *ServerInterfaceWrapper) GetDomain(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "domain" -------------
	var domain Domain

	err = runtime.BindStyledParameter("simple", false, "domain", ctx.Param("domain"), &domain)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter domain: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDomain(ctx, domain)
	return err
}

// GetRecord converts echo context to params.
func (w *ServerInterfaceWrapper) GetRecord(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "domain" -------------
	var domain Domain

	err = runtime.BindStyledParameter("simple", false, "domain", ctx.Param("domain"), &domain)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter domain: %s", err))
	}

	// ------------- Path parameter "record" -------------
	var record Record

	err = runtime.BindStyledParameter("simple", false, "record", ctx.Param("record"), &record)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter record: %s", err))
	}

	// ------------- Path parameter "recordType" -------------
	var recordType RecordType

	err = runtime.BindStyledParameter("simple", false, "recordType", ctx.Param("recordType"), &recordType)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter recordType: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRecord(ctx, domain, record, recordType)
	return err
}

// CreateRecord converts echo context to params.
func (w *ServerInterfaceWrapper) CreateRecord(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "domain" -------------
	var domain Domain

	err = runtime.BindStyledParameter("simple", false, "domain", ctx.Param("domain"), &domain)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter domain: %s", err))
	}

	// ------------- Path parameter "record" -------------
	var record Record

	err = runtime.BindStyledParameter("simple", false, "record", ctx.Param("record"), &record)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter record: %s", err))
	}

	// ------------- Path parameter "recordType" -------------
	var recordType RecordType

	err = runtime.BindStyledParameter("simple", false, "recordType", ctx.Param("recordType"), &recordType)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter recordType: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateRecord(ctx, domain, record, recordType)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router runtime.EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/domains", wrapper.GetDomains)
	router.PUT("/domains", wrapper.CreateDomain)
	router.GET("/domains/:domain", wrapper.GetDomain)
	router.GET("/domains/:domain/:record/:recordType", wrapper.GetRecord)
	router.PUT("/domains/:domain/:record/:recordType", wrapper.CreateRecord)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RUwY7UMAz9lchw7E4LiEtuSMuBK+KG9hBadyYoTYLjDoyq/jty2kxBGrqwnPYyySTP",
	"znvPdidowxCDR88J9ATRkBmQkfK/LgzGetnJL0TDJ6jAmwFBl8sKCL+NlrADzTRiBak94WAkii9RkInJ",
	"+iPMs2DbQN0fMq6XT8n4KZ/vZM2Af8k8l8sE2o/OzRVY34cMtewE65Gd7S93Pr26S0hn28oTZ6Rkg/CA",
	"uYIQ0ZtoQcObQ3NooMrssrv14mDeH5FlCRHJsA3+QwdaDu9XiDBPMfiEGf66aWRpg2f0OdLE6GybY+uv",
	"KeSibeIs45BuqKzKgSEyl0V1h6klG3nR8E45m1iFXq1sD0q9HyJflO2VDx4V/rCJcynSOAyGLgtzZa6R",
	"hEebGAm7kkRejuMNyS2hYbzfWusqWmrw6xMLUhnl8fuaVq0NJFyKufW0bObHXc612dr/8wQvCXvQ8KLe",
	"hqTeIOsTMD/sEhUv+tE5Je2jzJcwijcl9hbVelqEXDfSvLsCPpbJeZqA6lFkcfZvkXnc/sOYayl3++RZ",
	"yt5tXJkipHMRMpIDDSfmqOvahda4U0is3zaNfEl+n9R8rTo8qyUDzA/zzwAAAP//u3KSUt8FAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
