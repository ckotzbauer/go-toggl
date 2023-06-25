package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

var (
	Domain = "https://api.track.toggl.com%s"
)

type TogglContext struct {
	context  context.Context
	user     string
	password string
}

func NewTogglContext(ctx context.Context, user, password string) *TogglContext {
	return &TogglContext{context: ctx, user: user, password: password}
}

func doRequest[O interface{}](method, url string, ctx *TogglContext) (*O, error) {
	return doRequestWithBody[*any, O](method, url, ctx, nil)
}

func doRequestWithBody[I interface{}, O interface{}](method, url string, ctx *TogglContext, req *I) (*O, error) {
	var buffer *bytes.Buffer
	if req != nil {
		data, err := json.Marshal(req)
		if err != nil {
			return nil, err
		}

		buffer = bytes.NewBuffer(data)
	}

	var request *http.Request
	var err error

	if buffer == nil {
		request, err = http.NewRequestWithContext(ctx.context, method, fmt.Sprintf(Domain, url), nil)
	} else {
		request, err = http.NewRequestWithContext(ctx.context, method, fmt.Sprintf(Domain, url), buffer)
	}

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.SetBasicAuth(ctx.user, ctx.password)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer checkClose(response.Body, &err)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		return nil, nil
	}

	var resp *O
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("error %w: %s", err, string(body))
	}

	return resp, nil
}

func formatQuery(opt any) string {
	v, _ := query.Values(opt)
	q := v.Encode()
	if q != "" {
		return "?" + q
	} else {
		return q
	}
}

func checkClose(c io.Closer, err *error) {
	cErr := c.Close()
	if *err == nil {
		*err = cErr
	}
}
