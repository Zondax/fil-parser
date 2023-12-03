package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zondax/fil-parser/types"
	"io"
	"net/http"
	"strings"
)

const (
	ContentTypeHeader = "Content-Type"
	ApplicationJSON   = "application/json"
	WSSProtocol       = "wss://"
	HTTPSProtocol     = "https://"
	getTipsetByHeight = "Filecoin.ChainGetTipSetByHeight"
	jsonRPC           = "2.0"
	requestID         = 1
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type HTTPFilClient interface {
	GetTipsetByHeight(height uint64) (*types.GetTipsetByHeightResponse, error)
}

type httpFilClient struct {
	Client  *http.Client
	BaseURL string
}

func NewHttpFilClient(client *http.Client, baseURL string) HTTPFilClient {
	return &httpFilClient{
		Client:  client,
		BaseURL: baseURL,
	}
}

func (fc *httpFilClient) doHTTPRequest(methodName string, params []interface{}, response interface{}) error {
	requestData := types.ApiRequest{
		Jsonrpc: jsonRPC,
		Method:  methodName,
		Params:  params,
		ID:      requestID,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return fmt.Errorf("error marshaling request data: %v", err)
	}

	baseURL := strings.Replace(fc.BaseURL, WSSProtocol, HTTPSProtocol, 1)
	req, err := http.NewRequest(http.MethodPost, baseURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %v", err)
	}
	req.Header.Set(ContentTypeHeader, ApplicationJSON)

	resp, err := fc.Client.Do(req)
	if err != nil {
		return fmt.Errorf("error performing HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		errorMessage, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("received error status code: %d, message: %s", resp.StatusCode, string(errorMessage))
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	if len(responseBody) == 0 {
		return fmt.Errorf("empty response body")
	}

	if err = json.Unmarshal(responseBody, response); err != nil {
		return fmt.Errorf("error unmarshaling response: %v", err)
	}

	return nil
}

func (fc *httpFilClient) GetTipsetByHeight(height uint64) (*types.GetTipsetByHeightResponse, error) {
	response := &types.GetTipsetByHeightResponse{}
	err := fc.doHTTPRequest(getTipsetByHeight, []interface{}{height, []interface{}{}}, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
