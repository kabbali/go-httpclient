package examples

import (
	"github.com/kabbali/go-httpclient/gohttp"
	"github.com/kabbali/go-httpclient/gomime"
	"net/http"
	"time"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("agent-client").
		Build()
	return client
}
