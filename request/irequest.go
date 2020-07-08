package request

import "net/http"

type Request interface {
	Get(url string, ch chan *Result, headers ...*Header)
	Post(url string, body string, ch chan *Result, headers ...*Header)
	//Put(url string, body string, ch chan *Result, headers ...*Header)
	Patch(url string, body string, ch chan *Result, headers ...*Header)
	Delete(url string, ch chan *Result, headers ...*Header)
}

type Result struct {
	Result     []byte
	Error      error
	StatusCode int
}

func (r *Result) Is2XX() bool {
	return r.StatusCode >= http.StatusOK && r.StatusCode <= http.StatusIMUsed
}

type Header struct {
	Key   string
	Value string
}
