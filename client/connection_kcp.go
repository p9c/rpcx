package client

import (
	"net"

	"git.parallelcoin.io/dev/kcp9"
)

func newDirectKCPConn(c *Client, network, address string) (net.Conn, error) {
	return kcp.DialWithOptions(address, c.option.Block.(kcp.BlockCrypt), 10, 3)
}
