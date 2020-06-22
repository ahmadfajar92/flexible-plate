package src

import (
	"fmt"
	
	"bytes"
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"unicode/utf8"
	"github.com/aws/aws-lambda-go/events"
	"github.com/labstack/echo"
)

//LambdaResponder managing lambda response
type LambdaResponder struct {
	Echo *echo.Echo
}

type body struct {
	Message string `json:"message"`
}

//Handler for lambda
func (a *LambdaResponder) Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	newReq, err := a.ProxyEventToHTTPRequest(req)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	respWriter := NewProxyResponseWriter()

	a.Echo.ServeHTTP(http.ResponseWriter(respWriter), newReq)

	proxyResp, err := respWriter.GetProxyResponse()

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return proxyResp, nil
}

func (a *LambdaResponder) stripBasePrefix(req events.APIGatewayProxyRequest) string {

	var prefix, currentPath, newRequestPath string
	currentPath = req.Path
	newRequestPath = currentPath

	if !strings.HasPrefix(newRequestPath, "/") {
		newRequestPath = "/" + newRequestPath
	}

	if strings.HasSuffix(newRequestPath, "/") {
		newRequestPath = newRequestPath[:len(newRequestPath)-1]
	}
	if prefix != "" && len(prefix) > 1 && strings.HasPrefix(currentPath, prefix) {
		newRequestPath = strings.Replace(currentPath, prefix, "", 1)
	}
	return newRequestPath
}

//ProxyEventToHTTPRequest proxy to event handler
func (a *LambdaResponder) ProxyEventToHTTPRequest(req events.APIGatewayProxyRequest) (*http.Request, error) {
	decodedBody := []byte(req.Body)
	if req.IsBase64Encoded {
		base64Body, err := base64.StdEncoding.DecodeString(req.Body)
		if err != nil {
			return nil, err
		}
		decodedBody = base64Body
	}

	queryString := ""
	if len(req.QueryStringParameters) > 0 {
		queryString = "?"
		queryCnt := 0
		for q := range req.QueryStringParameters {
			if queryCnt > 0 {
				queryString += "&"
			}
			queryString += url.QueryEscape(q) + "=" + url.QueryEscape(req.QueryStringParameters[q])
			queryCnt++
		}
	}

	path := a.stripBasePrefix(req)
	httpRequest, err := http.NewRequest(
		strings.ToUpper(req.HTTPMethod),
		path+queryString,
		bytes.NewReader(decodedBody),
	)

	if err != nil {
		fmt.Printf("Could not convert request %s:%s to http.Request\n", req.HTTPMethod, req.Path)
		return nil, err
	}

	for h := range req.Headers {
		httpRequest.Header.Add(h, req.Headers[h])
	}

	return httpRequest, nil
}

// ProxyResponseWriter implements http.ResponseWriter and adds the method
// necessary to return an events.APIGatewayProxyResponse object
type ProxyResponseWriter struct {
	Headers http.Header `json:"headers"`
	Body    []byte      `json:"body"`
	Status  int         `json:"statusCode"`
}

// NewProxyResponseWriter returns a new ProxyResponseWriter object.
// The object is initialized with an empty map of headers and a
// status code of -1
func NewProxyResponseWriter() *ProxyResponseWriter {
	return &ProxyResponseWriter{
		Headers: make(http.Header),
		Status:  http.StatusOK,
	}

}

// Header implementation from the http.ResponseWriter interface.
func (r *ProxyResponseWriter) Header() http.Header {
	return r.Headers
}

// Write sets the response body in the object. If no status code
// was set before with the WriteHeader method it sets the status
// for the response to 200 OK.
func (r *ProxyResponseWriter) Write(body []byte) (int, error) {
	r.Body = body
	if r.Status == -1 {
		r.Status = http.StatusOK
	}

	return len(body), nil
}

//WriteHeader sets a status code for the response. This method is used
// for error responses.
func (r *ProxyResponseWriter) WriteHeader(status int) {
	r.Status = status
}

//GetProxyResponse converts the data passed to the response writer into
// an events.APIGatewayProxyResponse object.
// Returns a populated proxy response object. If the reponse is invalid, for example
// has no headers or an invalid status code returns an error.
func (r *ProxyResponseWriter) GetProxyResponse() (events.APIGatewayProxyResponse, error) {
	if len(r.Headers) == 0 {
		return events.APIGatewayProxyResponse{}, errors.New("No headers generated for response")
	}

	var output string
	isBase64 := false

	if utf8.Valid(r.Body) {
		output = string(r.Body)
	} else {
		output = base64.StdEncoding.EncodeToString(r.Body)
		isBase64 = true
	}

	proxyHeaders := make(map[string]string)

	for h := range r.Headers {
		proxyHeaders[h] = r.Headers.Get(h)
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      r.Status,
		Headers:         proxyHeaders,
		Body:            output,
		IsBase64Encoded: isBase64,
	}, nil
}
