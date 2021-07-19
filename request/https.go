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

type HttpsClient struct {
	client http.Client
}

func (this *HttpsClient) Get(url string, headers ...*Header) (resBody []byte, statusCode int, header http.Header, err error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", "go-mutils/1.0")

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

	// if res.TLS != nil {
	// 	for _, cert := range res.TLS.PeerCertificates {
	// 	}
	// }

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data, res.StatusCode, res.Header, err

}

func (this *HttpsClient) Post(url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPost, url, bodyReader)

	for _, head := range headers {
		if head == nil {
			continue
		}
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

func (this *HttpsClient) Patch(url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPatch, url, bodyReader)

	for _, head := range headers {
		if head == nil {
			continue
		}
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

func (this *HttpsClient) Put(url string, body []byte, headers ...*Header) (resBody []byte, statusCode int, err error) {
	bodyReader := bytes.NewReader(body)
	req, _ := http.NewRequest(http.MethodPut, url, bodyReader)

	for _, head := range headers {
		if head == nil {
			continue
		}
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

func (this *HttpsClient) Delete(url string, headers ...*Header) (resBody []byte, statusCode int, err error) {
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set("User-Agent", "go-mutils/1.0")

	for _, head := range headers {
		if head == nil {
			continue
		}
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

func NewHttpsClientWithByte(certBytes []byte, timeout time.Duration) (*HttpsClient, error) {
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
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: false,
					RootCAs:            clientCertPool,
				},
				Dial: func(network, addr string) (net.Conn, error) {
					return net.DialTimeout(network, addr, timeout)
				},
			},
		},
	}

	return httpClient, nil
}

func NewHttpsClient(caFile string, timeout time.Duration) (*HttpsClient, error) {
	certBytes, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, errors.New("Unable to read cert.pem: " + err.Error())
	}

	return NewHttpsClientWithByte(certBytes, timeout)
}
