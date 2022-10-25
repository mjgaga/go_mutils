package request

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	client http.Client
}

func (client *HttpClient) Head(ctx context.Context, url string, headers ...*Header) (statusCode int, header http.Header, err error) {
	req, _ := http.NewRequest(http.MethodHead, url, nil)
	req = req.WithContext(ctx)
	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Add(head.Key, head.Value)
	}

	res, err := client.client.Do(req)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, res.Header, err
}

func (client *HttpClient) Get(ctx context.Context, url string, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req = req.WithContext(ctx)
	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Add(head.Key, head.Value)
	}

	res, err := client.client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err
}

func (client *HttpClient) Post(ctx context.Context, url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPost, url, bodyReader)

	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Add(head.Key, head.Value)
	}

	res, err := client.client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err
}

func (client *HttpClient) Patch(ctx context.Context, url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPatch, url, bodyReader)
	req = req.WithContext(ctx)
	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Add(head.Key, head.Value)
	}

	res, err := client.client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err
}

func (client *HttpClient) Put(ctx context.Context, url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPut, url, bodyReader)
	req = req.WithContext(ctx)

	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Add(head.Key, head.Value)
	}

	res, err := client.client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err
}

func (client *HttpClient) Delete(ctx context.Context, url string, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	req = req.WithContext(ctx)
	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Add(head.Key, head.Value)
	}

	res, err := client.client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: http.Client{
			// CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 	return http.ErrUseLastResponse
			// },
			Transport: &http.Transport{
				DisableKeepAlives: true,
			},
		},
	}
}
