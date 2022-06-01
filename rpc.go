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

type JSONRPCRequest struct {
	Version string `json:"jsonrpc"`
	ID      string `json:"id"`
	Method  string `json:"method"`
	Params  any    `json:"params,omitempty"`
}

type JSONRPCResponse struct {
	Version string        `json:"jsonrpc"`
	ID      string        `json:"id"`
	Error   *JSONRPCError `json:"error,omitempty"`
}

type JSONRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type JSONRPCClient struct {
	endpoint string
	client   *http.Client
	context  context.Context
	headers  map[string]string
}

func NewClient(endpoint string, headers map[string]string) *JSONRPCClient {
	if headers == nil {
		return &JSONRPCClient{
			endpoint: endpoint,
			client:   &http.Client{},
			context:  context.Background(),
			headers:  map[string]string{},
		}
	} else {
		return &JSONRPCClient{
			endpoint: endpoint,
			client:   &http.Client{},
			context:  context.Background(),
			headers:  headers,
		}
	}
}

func (jrpc *JSONRPCClient) newHTTPRequest(RPCRequest *JSONRPCRequest) (*http.Request, error) {
	body, err := json.Marshal(RPCRequest)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(jrpc.context, "POST", jrpc.endpoint, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	for header, value := range jrpc.headers {
		request.Header.Set(header, value)
	}

	return request, nil
}

func (jrpc *JSONRPCClient) Call(RPCRequest *JSONRPCRequest) ([]byte, error) {
	HTTPRequest, err := jrpc.newHTTPRequest(RPCRequest)

	if err != nil {
		return nil, err
	}

	HTTPResponse, err := jrpc.client.Do(HTTPRequest)

	if err != nil {
		return nil, err
	}

	defer HTTPResponse.Body.Close()

	return io.ReadAll(HTTPResponse.Body)
}

func (jrpc *JSONRPCClient) GetRecentBlockhash() RecentBlockhash {
	responseBytes, _ := jrpc.Call(
		&JSONRPCRequest{
			Version: JSON_RPC_VERSION,
			ID:      uuid.NewString(),
			Method:  "getLatestBlockhash",
			Params: []map[string]string{
				{"commitment": "processed"},
			},
		},
	)

	blockhashResponse := &struct {
		JSONRPCResponse
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

	json.Unmarshal(responseBytes, &blockhashResponse)

	return RecentBlockhashFromString(blockhashResponse.Result.Value.Blockhash)
}
