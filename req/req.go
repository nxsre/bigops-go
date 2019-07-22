package req

import (
	"crypto/sha1"
	"fmt"
	"github.com/soopsio/bigops-go/utils"
	"github.com/soopsio/req"
	"go.uber.org/zap"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type proxyReq struct {
	key    string
	secret string
}

func (p *proxyReq) reqProxy(req *http.Request) (*url.URL, error) {
	timestamp := fmt.Sprint(time.Now().UnixNano() / 1000000)
	h := sha1.New()
	io.WriteString(h, timestamp)
	sign := utils.GetSignature(timestamp, p.secret)

	q := req.URL.Query()
	q.Add("accesskey", p.key)
	q.Add("sign", sign)
	q.Add("timestamp", timestamp)
	req.URL.RawQuery = q.Encode()

	return http.ProxyFromEnvironment(req)
}

// create a default client
func newClient(key, secret string) *http.Client {
	preq := proxyReq{
		key:    key,
		secret: secret,
	}
	jar, _ := cookiejar.New(nil)
	transport := &http.Transport{
		Proxy: preq.reqProxy,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	return &http.Client{
		Jar:       jar,
		Transport: transport,
		Timeout:   2 * time.Minute,
	}
}

type Req struct {
	*req.Req
}

func New(accessKey, secretKey string, debug bool, logger *zap.Logger) *Req {
	req.Debug = debug
	r := req.New()
	r.SetClient(newClient(accessKey, secretKey))
	r.SetLogger(logger)
	return &Req{r}
}
