package solago

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
)

const JSON_RPC_VERSION = "2.0"

type Request struct {
	Version string `json:"jsonrpc"`
	ID      string `json:"id"`
	Method  string `json:"method"`
	Params  any    `json:"params,omitempty"`
}

type Response struct {
	Version string `json:"jsonrpc"`
	ID      string `json:"id"`
	Error   *Error `json:"error,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type Client struct {
	endpoint string
	client   *http.Client
	context  context.Context
	headers  map[string]string
}

func NewClient(endpoint string, headers map[string]string) *Client {
	if headers == nil {
		return &Client{
			endpoint: endpoint,
			client:   &http.Client{},
			context:  context.Background(),
			headers:  map[string]string{},
		}
	} else {
		return &Client{
			endpoint: endpoint,
			client:   &http.Client{},
			context:  context.Background(),
			headers:  headers,
		}
	}
}

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

func (client *Client) GetRecentBlockhash() RecentBlockhash {
	responseBytes, _ := client.Call(
		&Request{
			Version: JSON_RPC_VERSION,
			ID:      uuid.NewString(),
			Method:  "getLatestBlockhash",
			Params: []map[string]string{
				{"commitment": "processed"},
			},
		},
	)

	blockhashResponse := &struct {
		Response
		Result struct {
			Context struct {
				Slot int `json:"slot"`
			} `json:"context"`
			Value struct {
				Blockhash            string `json:"blockhash"`
				LastValidBlockHeight int    `json:"lastValidBlockHeight"`
			} `json:"value"`
		} `json:"result"`
	}{}

	json.Unmarshal(responseBytes, blockhashResponse)

	return RecentBlockhashFromString(blockhashResponse.Result.Value.Blockhash)
}
