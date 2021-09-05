package operations

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/LQR471814/multicast/common"
)

var conn = &net.UDPConn{}

var ctx context.Context
var cancel context.CancelFunc

type MulticastPacket struct {
	Read     int
	Src      *net.UDPAddr
	Contents []byte
}

func init() {
	ctx, cancel = context.WithCancel(context.Background())
}

func listener(handler func(MulticastPacket)) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			buf := make([]byte, common.BUFFER)
			conn.SetReadDeadline(time.Now().Add(1 * time.Second))

			n, src, err := conn.ReadFromUDP(buf)
			if os.IsTimeout(err) {
				continue
			} else if err != nil {
				cancel()
				return
			}

			handler(MulticastPacket{
				Read:     n,
				Src:      src,
				Contents: buf,
			})
		}
	}
}

func Listen(group *net.UDPAddr, intf int, handler func(MulticastPacket)) error {
	networkInterface, err := net.InterfaceByIndex(intf)
	if err != nil {
		return err
	}

	conn, err = net.ListenMulticastUDP("udp4", networkInterface, group)
	if conn == nil {
		return common.BrokenInterfaceError{}
	}

	conn.SetReadBuffer(common.BUFFER)
	go listener(handler)

	return err
}

func Ping(group *net.UDPAddr, buf []byte) error {
	if conn == nil {
		return common.BrokenInterfaceError{}
	}

	_, err := conn.WriteToUDP(buf, group)
	if err != nil {
		return err
	}

	return nil
}

func Context() context.Context {
	return ctx
}

func Cancel() {
	cancel()
}
