package apiclient

import (
    "net/url"
    "log"
    "io"
    "io/ioutil"
    "encoding/json"
    "net/http"
    "errors"
)

const defaultLogMessage = "myDiamond-Scraper:"

type ApiClient struct {
    BaseURL    *url.URL
    UserAgent  string
    HttpClient *http.Client
}

type DefaultClient struct {
    Client ApiClient
}

type Request struct {
    TimeDelay int
}

func (c *ApiClient) Get(path string, headers http.Header) (Response, error) {
    response := Response{}
    rel := &url.URL{Path: path}
    u := c.BaseURL.ResolveReference(rel)
    request, err := http.NewRequest(http.MethodGet, u.String(), nil)
    if err != nil {
        log.Printf("%s Failed to create GET : %v\n", defaultLogMessage, err)
        return response, err
    }
    headers.Set("Accept", "application/json")
    headers.Set("User-Agent", c.UserAgent)
    request.Header = headers
    return c.Do(request)
}

func (c *ApiClient) Post(path string, body io.Reader, headers http.Header) (Response, error) {
    response := Response{}
    rel := &url.URL{Path: path}
    u := c.BaseURL.ResolveReference(rel)
    request, err := http.NewRequest(http.MethodPost, u.String(), body)
    if err != nil {
        log.Printf("%s Failed to create POST : %v\n", defaultLogMessage, err)
        return response, err
    }
    request.Header.Set("Content-Type", "application/json")
    headers.Set("User-Agent", c.UserAgent)
    request.Header = headers
    return c.Do(request)
}

func (c *ApiClient) Put(path string, body io.Reader, headers http.Header) (Response, error) {
    response := Response{}
    request, err := http.NewRequest(http.MethodPut, path, body)
    if err != nil {
        log.Printf("%s Failed to create PUT : %v\n", defaultLogMessage, err)
        return response, err
    }
    headers.Set("User-Agent", c.UserAgent)
    request.Header = headers
    return c.Do(request)
}

func (c *ApiClient) Do(request *http.Request) (Response, error) {
    resp := Response{}
    request.Close = true
    response, err := c.HttpClient.Do(request)
    if err != nil {
        log.Printf("%s Error sending HTTP request to %s: %v\n", defaultLogMessage, request.URL, err)
        return resp, err
    }
    if response.Body != nil {
        resp.Body, err = ioutil.ReadAll(response.Body)
        if err != nil {
            log.Printf("%s Do(): Failed to Read the Body of the Response", defaultLogMessage)
            return resp, err
        }
    }
    response.Body.Close()
    resp.StatusCode = StatusCode(response.StatusCode)
    return resp, err
}

type Response struct {
    Body       []byte
    StatusCode StatusCode
}

type ResponseError struct {
    Error string `bson:"error" json:"error"`
}

// Error return error message
func (hr Response) Error() error {
    e := ResponseError{}
    if !hr.StatusCode.IsSuccessful() {
        json.Unmarshal(hr.Body, &e)
        return errors.New(e.Error)
    }
    return nil
}

type StatusCode int

func (s StatusCode) IsSuccessful() bool {
    return s/100 == 2
}

