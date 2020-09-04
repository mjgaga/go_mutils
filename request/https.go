package request

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type HttpsClient struct {
	client http.Client
}

func (this *HttpsClient) Get(url string, ch chan *Result, headers ...*Header) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", "go-mutils/1.0")

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

	// if res.TLS != nil {
	// 	for _, cert := range res.TLS.PeerCertificates {
	// 	}
	// }

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	r.Error = err
	r.Result = data
	r.StatusCode = res.StatusCode

	ch <- r
}

func (this *HttpsClient) Post(url string, body string, ch chan *Result, headers ...*Header) {
	bodyReader := strings.NewReader(body)
	req, _ := http.NewRequest(http.MethodPost, url, bodyReader)

	for _, head := range headers {
		req.Header.Set(head.Key, head.Value)
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

func (this *HttpsClient) Patch(url string, body string, ch chan *Result, headers ...*Header) {
	bodyReader := strings.NewReader(body)
	req, _ := http.NewRequest(http.MethodPatch, url, bodyReader)

	for _, head := range headers {
		req.Header.Set(head.Key, head.Value)
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

func (this *HttpsClient) Delete(url string, ch chan *Result, headers ...*Header) {
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set("User-Agent", "go-mutils/1.0")

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

	// if res.TLS != nil {
	// 	for _, cert := range res.TLS.PeerCertificates {
	// 	}
	// }

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	r.Error = err
	r.Result = data
	r.StatusCode = res.StatusCode

	ch <- r
}

func (this *HttpsClient) SyncGet(url string, headers ...*Header) (resBody []byte, statusCode int, err error) {
	ch := make(chan *Result, 1)
	go this.Get(url, ch, headers...)
	res := <-ch
	return res.Result, res.StatusCode, res.Error
}

func (this *HttpsClient) SyncPost(url string, body string, headers ...*Header) (resBody []byte, statusCode int, err error) {
	ch := make(chan *Result, 1)
	go this.Post(url, body, ch, headers...)
	res := <-ch
	return res.Result, res.StatusCode, res.Error
}

func (this *HttpsClient) SyncPatch(url string, body string, headers ...*Header) (resBody []byte, statusCode int, err error) {
	ch := make(chan *Result, 1)
	go this.Patch(url, body, ch, headers...)
	res := <-ch
	return res.Result, res.StatusCode, res.Error
}

func (this *HttpsClient) SyncDelete(url string, headers ...*Header) (resBody []byte, statusCode int, err error) {
	ch := make(chan *Result, 1)
	go this.Delete(url, ch, headers...)
	res := <-ch
	return res.Result, res.StatusCode, res.Error
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
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: false,
					RootCAs:            clientCertPool,
				},
				Dial: func(network, addr string) (net.Conn, error) {
					timeout := time.Second * 10
					return net.DialTimeout(network, addr, timeout)
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
