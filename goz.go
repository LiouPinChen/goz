package goz

import (
	"github.com/jarcoal/httpmock"
	"net/http"
)

// NewClient new request object
func NewClient(opts ...Options) *Request {
	req := &Request{}

	if len(opts) > 0 {
		req.opts = opts[0]
	}

	return req
}

type IsolateRule struct {
	Method   string
	Url      string
	Status   int
	Response string
}

// NewIsolateClient new request object
func NewIsolateClient(opts ...Options) *Request {
	req := &Request{}

	if len(opts) > 0 {
		req.opts = opts[0]
	}

	mockCli := http.Client{}

	// if len(rules) > 0 {
		// 截捷伺服器的回傳結果
		httpmock.Activate()
		httpmock.RegisterResponder("GET", "http://172.17.0.1:9000/server/msg1",
			httpmock.NewStringResponder(200, `custom`))
	// }

	req.ReplaceCli(&mockCli)

	return req
}

// Reset NewIsolateClient
func ResetIsolateRule() {
	httpmock.DeactivateAndReset()
}