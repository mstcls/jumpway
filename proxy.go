package jumpway

import (
	"context"

	"github.com/wzshiming/bridge/chain"
	"github.com/wzshiming/bridge/multiple/proxy"
	_ "github.com/wzshiming/bridge/protocols/command"
	_ "github.com/wzshiming/bridge/protocols/connect"
	"github.com/wzshiming/bridge/protocols/local"
	_ "github.com/wzshiming/bridge/protocols/netcat"
	_ "github.com/wzshiming/bridge/protocols/shadowsocks"
	_ "github.com/wzshiming/bridge/protocols/smux"
	_ "github.com/wzshiming/bridge/protocols/socks4"
	_ "github.com/wzshiming/bridge/protocols/socks5"
	_ "github.com/wzshiming/bridge/protocols/ssh"
	_ "github.com/wzshiming/bridge/protocols/tls"
	_ "github.com/wzshiming/bridge/protocols/ws"
)

func RunProxy(ctx context.Context, address string, ways []string) error {
	listener, err := local.LOCAL.Listen(ctx, "tcp", address)
	if err != nil {
		return err
	}

	dialer, err := chain.Default.BridgeChain(local.LOCAL, ways...)
	if err != nil {
		return err
	}

	proxies := []string{
		"http://" + address,
		"socks5://" + address,
		"socks4://" + address,
	}
	proxy, err := proxy.NewProxy(ctx, proxies, dialer)
	if err != nil {
		return err
	}

	host := proxy.Match(address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go host.ServeConn(conn)
	}

	return nil
}
