package rpc

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) newHTTPRequest(RPCRequest *Request) (*http.Request, error) {
	body, err := json.Marshal(RPCRequest)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(client.context, "POST", client.endpoint, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	for header, value := range client.headers {
		request.Header.Set(header, value)
	}

	return request, nil
}

func (client *Client) Call(RPCRequest *Request) ([]byte, error) {
	HTTPRequest, err := client.newHTTPRequest(RPCRequest)

	if err != nil {
		return nil, err
	}

	HTTPResponse, err := client.client.Do(HTTPRequest)

	if err != nil {
		return nil, err
	}

	defer HTTPResponse.Body.Close()

	return io.ReadAll(HTTPResponse.Body)
}
