package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

var (
	Domain = "https://api.track.toggl.com%s"
)

type TogglContext struct {
	context  context.Context
	user     string
	password string
}

func NewToggleContext(ctx context.Context, user, password string) *TogglContext {
	return &TogglContext{context: ctx, user: user, password: password}
}

func doRequest[O interface{}](method, url string, ctx *TogglContext) (*O, error) {
	return doRequestWithBody[*any, O](method, url, ctx, nil)
}

func doRequestWithBody[I interface{}, O interface{}](method, url string, ctx *TogglContext, req *I) (*O, error) {
	var buffer = bytes.NewBuffer([]byte("{}"))
	if req != nil {
		data, err := json.Marshal(req)
		if err != nil {
			return nil, err
		}

		buffer = bytes.NewBuffer(data)
	}

	request, err := http.NewRequestWithContext(ctx.context, http.MethodPost, fmt.Sprintf(Domain, url), buffer)
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
		return nil, err
	}

	return resp, nil
}

func formatQuery(query any) string {
	formatted := ""
	v := reflect.ValueOf(query)
	tQuery := reflect.TypeOf(query)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		var fieldValue reflect.Value
		if field.Type().Kind() == reflect.Pointer {
			fieldValue = field.Elem()
		} else {
			fieldValue = field
		}

		if field.IsNil() {
			continue
		}

		value := fieldValue.Interface()

		if value == "" {
			continue
		}

		if formatted == "" {
			formatted = formatted + "?"
		} else {
			formatted = formatted + "&"
		}

		formatted = fmt.Sprintf("%s%s=%v", formatted, tQuery.Field(i).Tag.Get("json"), value)
	}

	return formatted
}

func checkClose(c io.Closer, err *error) {
	cErr := c.Close()
	if *err == nil {
		*err = cErr
	}
}
