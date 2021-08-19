package gopwn

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

type HTTPClientOptions struct {
	Timeout         time.Duration
	ProxyURL        string
	TLSClientConfig *tls.Config
}

func HTTPGet(url string, optFns ...func(o *HTTPClientOptions)) ([]byte, error) {
	options := HTTPClientOptions{
		Timeout: 5 * time.Second,
	}
	for _, fn := range optFns {
		fn(&options)
	}
	client := newHTTPClient(options)

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func Download(url, filename string, optFns ...func(o *HTTPClientOptions)) error {
	options := HTTPClientOptions{
		Timeout: 10 * time.Second,
	}
	for _, fn := range optFns {
		fn(&options)
	}
	client := newHTTPClient(options)

	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}

func newHTTPClient(options HTTPClientOptions) *http.Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = options.TLSClientConfig

	if options.ProxyURL != "" {
		proxyURL, _ := url.Parse(options.ProxyURL)
		transport.Proxy = http.ProxyURL(proxyURL)
	}

	return &http.Client{
		Timeout:   options.Timeout,
		Transport: transport,
	}
}
