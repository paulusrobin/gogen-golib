package httpclient

import (
	"context"
	"net/http"
)

type (
	Client interface {
		Do(ctx context.Context, request *http.Request) (*http.Response, error)
	}
)
