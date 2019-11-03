package cs104

import (
	"crypto/tls"
	"errors"
	"net"
	"net/url"
	"time"
)

// DefaultReconnectInterval defined default value
const DefaultReconnectInterval = 1 * time.Minute

type seqPending struct {
	seq      uint16
	sendTime time.Time
}

func openConnection(uri *url.URL, tlsc *tls.Config, timeout time.Duration) (net.Conn, error) {
	switch uri.Scheme {
	case "tcp":
		return net.DialTimeout("tcp", uri.Host, timeout)
	case "ssl":
		fallthrough
	case "tls":
		fallthrough
	case "tcps":
		return tls.DialWithDialer(&net.Dialer{Timeout: timeout}, "tcp", uri.Host, tlsc)
	}
	return nil, errors.New("Unknown protocol")
}
