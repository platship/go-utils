package reqx

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Reqx struct {
	Retry      int
	Header     map[string]string
	Proxy      func(*http.Request) (*url.URL, error)
	Timeout    time.Duration
	retryCount int
}

func New() *Reqx {
	r := &Reqx{Header: make(map[string]string)}
	r.AddHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36 Edg/96.0.1054.62")
	return r
}

func (r *Reqx) SetRetry(val int) *Reqx {
	r.Retry = val
	return r
}

func (r *Reqx) SetProxyURL(val *url.URL) *Reqx {
	if val != nil {
		r.Proxy = http.ProxyURL(val)
	}
	return r
}

func (r *Reqx) SetProxyURLStr(val string) *Reqx {
	if val != "" {
		u, _ := url.Parse(val)
		r.SetProxyURL(u)
	}
	return r
}

func (r *Reqx) ResetRetryCount() *Reqx {
	r.retryCount = 0
	return r
}

func (r *Reqx) SetHeader(val map[string]string) *Reqx {
	r.Header = val
	return r
}

func (r *Reqx) AddHeader(key, val string) *Reqx {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header[key] = val
	return r
}

func (r *Reqx) SetUserAgentMust(ua string) *Reqx {
	if ua == "" {
		return r
	}
	r.AddHeader("User-Agent", ua)
	return r
}

func (r *Reqx) SetReferer(val string) *Reqx {
	r.AddHeader("Referer", val)
	return r
}

func (r *Reqx) SetTimeout(val time.Duration) *Reqx {
	r.Timeout = val
	return r
}

func (r *Reqx) SetTimeoutSeconds(val int) *Reqx {
	if val > 0 {
		r.Timeout = time.Duration(val) * time.Second
	}
	return r
}

func (r *Reqx) SetCookie(val string) *Reqx {
	r.AddHeader("Cookie", val)
	return r
}

func (r *Reqx) SetHost(val string) *Reqx {
	r.AddHeader("Host", val)
	return r
}

func (r *Reqx) GetBody(uri string) (_ []byte, err error) {
	resp, err := r.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (r *Reqx) PostReturnBody(uri string, body io.Reader) (_ []byte, err error) {
	resp, err := r.Post(uri, body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (r *Reqx) Post(uri string, body io.Reader) (*http.Response, error) {
	return r.Request("POST", uri, body)
}

func (r *Reqx) Get(uri string) (resp *http.Response, err error) {
	return r.Request("GET", uri, nil)
}

func (r *Reqx) Head(uri string) (resp *http.Response, err error) {
	return r.Request("HEAD", uri, nil)
}

func (r *Reqx) Request(method, uri string, body io.Reader) (resp *http.Response, err error) {
	if strings.HasPrefix(uri, "//") {
		uri = "http:" + uri
	}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return
	}

	if r.Header != nil {
		for k, v := range r.Header {
			req.Header.Set(k, v)
		}
	}
	resp, err = r.client().Do(req)
	if err != nil && r.retryCount < r.Retry {
		r.retryCount++
		return r.Request(method, uri, body)
	}
	return
}

func (r *Reqx) client() *http.Client {
	c := &http.Client{}
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	if r.Proxy != nil {
		transport.Proxy = r.Proxy
	}
	c.Transport = transport
	if r.Timeout > 0 {
		c.Timeout = r.Timeout
	}
	return c
}
