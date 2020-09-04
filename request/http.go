package request

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type HttpClient struct {
	client http.Client
}

func (this *HttpClient) Get(url string, ch chan *Result, headers ...*Header) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	for _, head := range headers {
		req.Header.Add(head.Key, head.Value)
	}

	r := &Result{}
	res, err := this.client.Do(req)
	if err != nil {
		r.Error = err
		ch <- r
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	r.Error = err
	r.Result = data
	r.StatusCode = res.StatusCode

	ch <- r
}

func (this *HttpClient) Post(url string, body string, ch chan *Result, headers ...*Header) {
	bodyReader := strings.NewReader(body)
	req, _ := http.NewRequest(http.MethodPost, url, bodyReader)
	for _, head := range headers {
		req.Header.Add(head.Key, head.Value)
	}

	r := &Result{}
	res, err := this.client.Do(req)
	if err != nil {
		r.Error = err
		ch <- r
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	r.Error = err
	r.Result = data
	r.StatusCode = res.StatusCode

	ch <- r
}

func (this *HttpClient) Patch(url string, body string, ch chan *Result, headers ...*Header) {
	bodyReader := strings.NewReader(body)
	req, _ := http.NewRequest(http.MethodPatch, url, bodyReader)
	for _, head := range headers {
		req.Header.Add(head.Key, head.Value)
	}

	r := &Result{}
	res, err := this.client.Do(req)
	if err != nil {
		r.Error = err
		ch <- r
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	r.Error = err
	r.Result = data
	r.StatusCode = res.StatusCode

	ch <- r
}

func (this *HttpClient) Delete(url string, ch chan *Result, headers ...*Header) {
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	for _, head := range headers {
		req.Header.Add(head.Key, head.Value)
	}

	r := &Result{}
	res, err := this.client.Do(req)
	if err != nil {
		r.Error = err
		ch <- r
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	r.Error = err
	r.Result = data
	r.StatusCode = res.StatusCode

	ch <- r
}

func (this *HttpClient) SyncGet(url string, headers ...*Header) (resBody []byte, statusCode int, err error) {
	ch := make(chan *Result, 1)
	go this.Get(url, ch, headers...)
	res := <-ch
	return res.Result, res.StatusCode, res.Error
}

func (this *HttpClient) SyncPost(url string, body string, headers ...*Header) (resBody []byte, statusCode int, err error) {
	ch := make(chan *Result, 1)
	go this.Post(url, body, ch, headers...)
	res := <-ch
	return res.Result, res.StatusCode, res.Error
}

func (this *HttpClient) SyncPatch(url string, body string, headers ...*Header) (resBody []byte, statusCode int, err error) {
	ch := make(chan *Result, 1)
	go this.Patch(url, body, ch, headers...)
	res := <-ch
	return res.Result, res.StatusCode, res.Error
}

func (this *HttpClient) SyncDelete(url string, headers ...*Header) (resBody []byte, statusCode int, err error) {
	ch := make(chan *Result, 1)
	go this.Delete(url, ch, headers...)
	res := <-ch
	return res.Result, res.StatusCode, res.Error
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: http.Client{
			// CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 	return http.ErrUseLastResponse
			// },
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					timeout := time.Second * 10
					return net.DialTimeout(network, addr, timeout)
				},
			},
		},
	}
}
