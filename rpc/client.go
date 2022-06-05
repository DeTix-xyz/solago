package rpc

import (
	"context"
	"net/http"
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
