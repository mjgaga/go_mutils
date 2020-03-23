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

type HttpsClientX509 struct {
	client http.Client
}

func (this *HttpsClientX509) Get(url string, ch chan *Result, headers ...*Header) {
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

	if res.StatusCode == http.StatusOK {
		r.Error = err
		r.Result = data
	} else {
		r.Error = errors.New(string(data))
		r.Result = nil
	}

	ch <- r
}

func (this *HttpsClientX509) Post(url string, body string, ch chan *Result, headers ...*Header) {
	bodyReader := strings.NewReader(body)
	req, _ := http.NewRequest(http.MethodPost, url, bodyReader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

	ch <- r
}

func (this *HttpsClientX509) Patch(url string, body string, ch chan *Result, headers ...*Header) {
	bodyReader := strings.NewReader(body)
	req, _ := http.NewRequest(http.MethodPatch, url, bodyReader)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

	ch <- r
}

func (this *HttpsClientX509) Delete(url string, ch chan *Result, headers ...*Header) {
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

	if res.StatusCode == http.StatusOK {
		r.Error = err
		r.Result = data
	} else {
		r.Error = errors.New(string(data))
		r.Result = nil
	}

	ch <- r
}

func NewHttpsClientX509(caFile, certFile, keyFile string) (*HttpsClientX509, error) {
	certBytes, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, errors.New("Unable to read cert.pem: " + err.Error())
	}

	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)

	if !ok {
		return nil, errors.New("failed to parse root certificate")
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
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
					timeout := time.Second * 10
					return net.DialTimeout(network, addr, timeout)
				},
			},
		},
	}

	return httpClient, nil
}
