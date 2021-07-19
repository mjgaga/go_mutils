package request

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type HttpsClientX509 struct {
	client http.Client
}

func (this *HttpsClientX509) Get(url string, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", "go-mutils/1.0")

	for _, head := range headers {
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

func (this *HttpsClientX509) Post(url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPost, url, bodyReader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for _, head := range headers {
		req.Header.Set(head.Key, head.Value)
	}

	res, err := this.client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, err
}

func (this *HttpsClientX509) Patch(url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPatch, url, bodyReader)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for _, head := range headers {
		req.Header.Set(head.Key, head.Value)
	}

	res, err := this.client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, err
}

func (this *HttpsClientX509) Put(url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPut, url, bodyReader)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for _, head := range headers {
		req.Header.Set(head.Key, head.Value)
	}

	res, err := this.client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, err
}

func (this *HttpsClientX509) Delete(url string, headers ...*Header) (resBody []byte, statusCode int, err error) {
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set("User-Agent", "go-mutils/1.0")

	for _, head := range headers {
		req.Header.Add(head.Key, head.Value)
	}

	res, err := this.client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, err
}

func NewHttpsClientX509WithBytes(caBytes, certBytes, keyData []byte, timeout time.Duration) (*HttpsClientX509, error) {
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
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: false,
					RootCAs:            clientCertPool,
					Certificates:       []tls.Certificate{cert},
				},
				Dial: func(network, addr string) (net.Conn, error) {
					return net.DialTimeout(network, addr, timeout)
				},
			},
		},
	}

	return httpClient, nil
}

func NewHttpsClientX509(caFile, certFile, keyFile string, timeout time.Duration) (*HttpsClientX509, error) {
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
	return NewHttpsClientX509WithBytes(certBytes, certPEMBlock, keyPEMBlock, timeout)
}
