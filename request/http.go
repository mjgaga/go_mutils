package request

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type HttpClient struct {
	client http.Client
}

func (this *HttpClient) Head(url string, headers ...*Header) (statusCode int, header http.Header, err error) {
	req, _ := http.NewRequest(http.MethodHead, url, nil)
	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Add(head.Key, head.Value)
	}

	res, err := this.client.Do(req)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, res.Header, err
}

func (this *HttpClient) Get(url string, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Add(head.Key, head.Value)
	}

	res, err := this.client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err
}

func (this *HttpClient) Post(url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPost, url, bodyReader)
	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Add(head.Key, head.Value)
	}

	res, err := this.client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err
}

func (this *HttpClient) Patch(url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPatch, url, bodyReader)
	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Add(head.Key, head.Value)
	}

	res, err := this.client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err
}

func (this *HttpClient) Put(url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPut, url, bodyReader)
	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Add(head.Key, head.Value)
	}

	res, err := this.client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err
}

func (this *HttpClient) Delete(url string, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Add(head.Key, head.Value)
	}

	res, err := this.client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err
}

func NewHttpClient(timeout time.Duration) *HttpClient {
	return &HttpClient{
		client: http.Client{
			// CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 	return http.ErrUseLastResponse
			// },
			Transport: &http.Transport{
				DisableKeepAlives: true,
				Dial: func(network, addr string) (net.Conn, error) {
					return net.DialTimeout(network, addr, timeout)
				},
			},
		},
	}
}
