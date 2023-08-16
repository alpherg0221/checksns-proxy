package proxy

import (
	"net"
	"net/http"
	"net/url"

	"github.com/AdguardTeam/gomitmproxy"
)

func StartProxy() error {
	proxy := gomitmproxy.NewProxy(gomitmproxy.Config{
		ListenAddr: &net.TCPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: 8080,
		},
		OnRequest: func(session *gomitmproxy.Session) (request *http.Request, response *http.Response) {
			req := session.Request()

			if req.Method != "CONNECT" {
				req.URL, _ = url.Parse("https://github.com/AdguardTeam/gomitmproxy")
			}

			return req, nil
		},
	})
	err := proxy.Start()
	if err != nil {
		return err
	}

	return nil
}
