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

type HttpsClientX509 struct {
	client http.Client
}

func (client *HttpsClientX509) Head(ctx context.Context, url string, headers ...*Header) (statusCode int, header http.Header, err error) {
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

func (client *HttpsClientX509) Get(ctx context.Context, url string, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
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

func (client *HttpsClientX509) Post(ctx context.Context, url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
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

func (client *HttpsClientX509) Patch(ctx context.Context, url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
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

func (client *HttpsClientX509) Put(ctx context.Context, url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
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

func (client *HttpsClientX509) Delete(ctx context.Context, url string, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
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

func NewHttpsClientX509WithBytes(caBytes, certBytes, keyData []byte, insecureSkipVerify bool) (*HttpsClientX509, error) {
	clientCertPool := x509.NewCertPool()
	if ok := clientCertPool.AppendCertsFromPEM(caBytes); !ok {
		return nil, errors.New("failed to parse root certificate")
	}
	cert, err := tls.X509KeyPair(certBytes, keyData)
	if err != nil {
		return nil, err
	}
	httpClient := &HttpsClientX509{
		client: http.Client{
			// CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 	return http.ErrUseLastResponse
			// },
			Transport: &http.Transport{
				DisableKeepAlives: true,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecureSkipVerify,
					RootCAs:            clientCertPool,
					Certificates:       []tls.Certificate{cert},
				},
			},
		},
	}

	return httpClient, nil
}

func NewHttpsClientX509(caFile, certFile, keyFile string, insecureSkipVerify bool) (*HttpsClientX509, error) {
	certBytes, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, errors.New("Unable to read cert.pem: " + err.Error())
	}

	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)

	if !ok {
		return nil, errors.New("failed to parse root certificate")
	}

	certPEMBlock, err := ioutil.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	keyPEMBlock, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	return NewHttpsClientX509WithBytes(certBytes, certPEMBlock, keyPEMBlock, insecureSkipVerify)
}
