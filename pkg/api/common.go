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
	Domain = "https://api.track.toggl.com/api/v9"
)

type TogglContext struct {
	context  context.Context
	user     string
	password string
}

func NewToggleContext(ctx context.Context, user, password string) *TogglContext {
	return &TogglContext{context: ctx, user: user, password: password}
}

func DoRequest[O interface{}](method, url string, ctx *TogglContext) (*O, error) {
	return DoRequestWithBody[*any, O](method, url, ctx, nil)
}

func DoRequestWithBody[I interface{}, O interface{}](method, url string, ctx *TogglContext, req *I) (*O, error) {
	var buffer *bytes.Buffer
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
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func FormatQuery(query any) string {
	formatted := ""
	v := reflect.ValueOf(query)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		t := field.Type()
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
		}

		var value any
		switch t.Kind() {
		case reflect.String:
			value = field.String()
		case reflect.Bool:
			value = field.Bool()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = field.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = field.Uint()
		case reflect.Float32, reflect.Float64:
			value = field.Float()
		default:
			value = ""
		}

		if value == "" {
			continue
		}

		if formatted == "" {
			formatted = formatted + "?"
		} else {
			formatted = formatted + "&"
		}

		formatted = fmt.Sprintf("%s%s=%s", formatted, t.Field(i).Tag.Get("json"), value)
	}

	return formatted
}

func checkClose(c io.Closer, err *error) {
	cErr := c.Close()
	if *err == nil {
		*err = cErr
	}
}
