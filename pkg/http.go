package pkg

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Http struct {
	Timeout            time.Duration
	Method             string
	Header             map[string][]string
	URL                *url.URL
	StatusCodes        []int32
	SkipInsecureVerify bool
}

func (h *Http) Connect() error {

	req := &http.Request{Method: h.Method, Header: h.Header, URL: h.URL}
	client := client(h.Timeout, h.SkipInsecureVerify)
	resp, err := client.Do(req)
	if err != nil || len(h.StatusCodes) > 0 {
		return err
	}

	for _, c := range h.StatusCodes {
		if int(c) == resp.StatusCode {
			return nil
		}
	}

	return fmt.Errorf("response status code: %d, expected: %v", resp.StatusCode, h.StatusCodes)
}

func client(timeout time.Duration, skipInsecureVerify bool) *http.Client {
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipInsecureVerify,
		},
	}

	return &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}

}
