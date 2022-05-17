package egohttp

import "net/http"

type Response struct {
	*http.Response
}

func NewResponse(resp *http.Response) *Response {
	return &Response{
		Response: resp,
	}
}
