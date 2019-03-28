package common

import (
	"crypto/tls"
	"github.com/golang/glog"
	"net"
	"net/http"
	"net/url"

	"strings"

	"net/http/cookiejar"
	"time"

	"context"
	"log"

	"golang.org/x/net/proxy"
)

type HttpClient struct {
	ProxyAddr string
	Client    http.Client
}

const (
	DefaultIdleTimeout    = 60 * time.Second
	DefaultConnectTimeout = 60 * time.Second
)

// Conn wraps a net.Conn, and sets a deadline for every read
// and write operation.
type TimeoutConn struct {
	net.Conn
	IdleTimeout time.Duration
}

//使用自定义出口协议,注意,前缀要全部使用小写
//如果是代理,那么使用 http:// 或者 https:// 类型的地址,如果使用出口 IP, 那么直接使用 ip:// 作为前缀
//如果使用ipv6, 那么使用`[]`把地址包起来
//例如:
//		http://14845132.xgj.me:27035
//		socks5://14845132.xgj.me:27035
//		ip://192.168.1.12
//		ip://[2607:5300:60:6566::]

func MakeTransportX(proxyAddr string) (transport *http.Transport) {
	transport = new(http.Transport)
	transport.MaxIdleConnsPerHost = 16
	//disable verify ssl
	//transport.TLSClientConfig = &tls.Config{
	//	InsecureSkipVerify: true,
	//}
	//Disable http2.0
	transport.TLSNextProto = make(map[string]func(authority string, c *tls.Conn) http.RoundTripper)
	glog.V(5).Infoln("Disable http2.0")
	var (
		localAddr string
		dialer    proxy.Dialer
	)
	proxyAddr = strings.TrimSpace(proxyAddr)
	proxyUrl, err := url.Parse(proxyAddr)
	if err != nil {
		log.Printf(`MakeTransportX, proxyAddr ("%s") have wrong format, err %v`, proxyAddr, err)
	}
	switch proxyUrl.Scheme {
	case "ip":
		localAddr = proxyUrl.Host
	case "http", "socks5":
		transport.Proxy = http.ProxyURL(proxyUrl)
	case "https":
		// Disable HTTP/2.
		transport.Proxy = http.ProxyURL(proxyUrl)
		transport.TLSNextProto = make(map[string]func(authority string, c *tls.Conn) http.RoundTripper)
	default:
		log.Printf(`MakeTransportX, proxyAddr ("%s") use unsupport scheme %s`, proxyAddr, proxyUrl.Scheme)
	}

	transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		var (
			conn net.Conn
		)
		if dialer != nil {
			var err error
			conn, err = dialer.Dial(network, addr)
			if err != nil {
				return nil, err
			}
		} else {
			d := net.Dialer{Timeout: DefaultConnectTimeout}
			if localAddr != "" && localAddr[0] == '[' {
				//如果本地ip地址以"["开头, 那么是ipv6地址，强制使用 tcp6 拨号
				network = "tcp6"
			}
			lAddr, err := net.ResolveTCPAddr(network, localAddr+":0")
			if err != nil {
				return nil, err
			}
			d.LocalAddr = lAddr
			conn, err = d.DialContext(ctx, network, addr)
			if err != nil {
				return nil, err
			}
		}

		return NewTimeoutConn(conn, DefaultIdleTimeout)
	}
	return transport

}

func MakeNoProxyTransportX() (transport *http.Transport) {
	transport = new(http.Transport)
	transport.MaxIdleConnsPerHost = 16
	//disable verify ssl
	//transport.TLSClientConfig = &tls.Config{
	//	InsecureSkipVerify: true,
	//}
	//Disable http2.0
	transport.TLSNextProto = make(map[string]func(authority string, c *tls.Conn) http.RoundTripper)
	glog.V(5).Infoln("Disable http2.0")
	var (
		localAddr string
		dialer    proxy.Dialer
	)

	transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		var (
			conn net.Conn
		)
		if dialer != nil {
			var err error
			conn, err = dialer.Dial(network, addr)
			if err != nil {
				return nil, err
			}
		} else {
			d := net.Dialer{Timeout: DefaultConnectTimeout}
			if localAddr != "" && localAddr[0] == '[' {
				//如果本地ip地址以"["开头, 那么是ipv6地址，强制使用 tcp6 拨号
				network = "tcp6"
			}
			lAddr, err := net.ResolveTCPAddr(network, localAddr+":0")
			if err != nil {
				return nil, err
			}
			d.LocalAddr = lAddr
			conn, err = d.DialContext(ctx, network, addr)
			if err != nil {
				return nil, err
			}
		}

		return NewTimeoutConn(conn, DefaultIdleTimeout)
	}
	return transport

}

func NewTimeoutConn(conn net.Conn, idleTimeout time.Duration) (net.Conn, error) {
	c := &TimeoutConn{
		Conn:        conn,
		IdleTimeout: idleTimeout,
	}
	if c.IdleTimeout > 0 {
		deadline := time.Now().Add(idleTimeout)
		if e := c.Conn.SetDeadline(deadline); e != nil {
			return nil, e
		}
	}
	return c, nil
}
func (c *TimeoutConn) Read(b []byte) (int, error) {
	n, e := c.Conn.Read(b)
	if c.IdleTimeout > 0 && n > 0 && e == nil {
		err := c.Conn.SetDeadline(time.Now().Add(c.IdleTimeout))
		if err != nil {
			return 0, err
		}
	}
	return n, e
}

func (c *TimeoutConn) Write(b []byte) (int, error) {
	n, e := c.Conn.Write(b)
	if c.IdleTimeout > 0 && n > 0 && e == nil {
		err := c.Conn.SetDeadline(time.Now().Add(c.IdleTimeout))
		if err != nil {
			return 0, err
		}
	}
	return n, e
}

func NewHttpClient(proxyAddr string) *HttpClient {
	c := &HttpClient{ProxyAddr: proxyAddr}
	//Follow redirect 时复制 http header
	c.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for attr, val := range via[0].Header {
			if _, ok := req.Header[attr]; !ok {
				req.Header[attr] = val
			}
		}
		return nil
	}

	//	c.Client.Timeout = 75 * time.Second
	return c
}

func (hc *HttpClient) mkTransport() {
	if hc.Client.Transport != nil || hc.ProxyAddr == "" {

		hc.Client.Transport = MakeNoProxyTransportX()
		return
	}
	hc.Client.Transport = MakeTransportX(hc.ProxyAddr)
}

func (hc *HttpClient) Do(req *http.Request) (resp *http.Response, err error) {
	hc.mkTransport()
	return hc.Client.Do(req)
}

func (hc *HttpClient) EnableCookie() {
	if hc.Client.Jar == nil {
		cookieJar, _ := cookiejar.New(nil)
		hc.Client.Jar = cookieJar
	}
}

func (hc *HttpClient) DisableCookie() {
	hc.Client.Jar = nil
}

func (hc *HttpClient) IsCookieEnabled() bool {
	return hc.Client.Jar != nil
}

func (hc *HttpClient) GetCookies(u *url.URL) []*http.Cookie {
	if hc.Client.Jar == nil {
		return nil
	}
	return hc.Client.Jar.Cookies(u)
}

func (hc *HttpClient) GetCookie(u *url.URL, key string) *http.Cookie {
	if hc.Client.Jar == nil {
		return nil
	}
	//	u, e := url.Parse(rawUrl)
	//	if  e != nil {
	//		return nil
	//	}
	for _, c := range hc.Client.Jar.Cookies(u) {
		if c.Name == "key" {
			return c
		}
	}
	return nil
}
