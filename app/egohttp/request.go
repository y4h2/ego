package egohttp

import "net/http"

type Request struct {
	*http.Request
}

func NewRequestByHTTP(r *http.Request) *Request {
	return &Request{
		Request: r,
	}
}
