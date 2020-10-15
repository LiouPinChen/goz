package goz

import (
	"github.com/jarcoal/httpmock"
	"net/http"
)

func NewClient(opts ...Options) *Request {
	req := &Request{}

	// 截捷伺服器的回傳結果
	httpmock.RegisterResponder("GET", "http://172.17.0.1:9000/temporary",
		httpmock.NewStringResponder(200, `custom`))
	mockCli := &http.Client{}

	req.Cli = mockCli // 用有隔離能力的 client 取代原來的客戶端

	return req
}
