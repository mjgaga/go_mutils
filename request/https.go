package request

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/http"
)

type HttpsClient struct {
	client http.Client
}

func (client *HttpsClient) Head(ctx context.Context, url string, headers ...*Header) (statusCode int, header http.Header, err error) {
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

func (client *HttpsClient) Get(ctx context.Context, url string, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
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

	// if res.TLS != nil {
	// 	for _, cert := range res.TLS.PeerCertificates {
	// 	}
	// }

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err

}

func (client *HttpsClient) Post(ctx context.Context, url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPost, url, bodyReader)
	req = req.WithContext(ctx)

	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Set(head.Key, head.Value)
	}

	res, err := client.client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err
}

func (client *HttpsClient) Patch(ctx context.Context, url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPatch, url, bodyReader)
	req = req.WithContext(ctx)

	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Set(head.Key, head.Value)
	}

	res, err := client.client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err
}

func (client *HttpsClient) Put(ctx context.Context, url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPut, url, bodyReader)
	req = req.WithContext(ctx)

	for _, head := range headers {
		if head == nil {
			continue
		}
		req.Header.Set(head.Key, head.Value)
	}

	res, err := client.client.Do(req)
	if err != nil {
		return nil, 0, nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err
}

func (client *HttpsClient) Delete(ctx context.Context, url string, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
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

func NewHttpsClientWithByte(certBytes []byte) (*HttpsClient, error) {
	clientCertPool := x509.NewCertPool()

	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		return nil, errors.New("failed to parse root certificate")
	}

	httpClient := &HttpsClient{
		client: http.Client{
			// CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 	return http.ErrUseLastResponse
			// },
			Transport: &http.Transport{
				DisableKeepAlives: true,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: false,
					RootCAs:            clientCertPool,
				},
			},
		},
	}

	return httpClient, nil
}

func NewHttpsClient(caFile string) (*HttpsClient, error) {
	certBytes, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, errors.New("Unable to read cert.pem: " + err.Error())
	}

	return NewHttpsClientWithByte(certBytes)
}
