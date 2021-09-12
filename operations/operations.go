package operations

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/LQR471814/multicast/common"
)

var pingConn *net.UDPConn

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

func listener(conn *net.UDPConn, handler func(MulticastPacket)) {
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

	conn, err := net.ListenMulticastUDP("udp4", networkInterface, group)
	if err != nil {
		return err
	}

	conn.SetReadBuffer(common.BUFFER)
	go listener(conn, handler)

	return err
}

func Ping(group *net.UDPAddr, buf []byte) error {
	var err error

	if pingConn == nil {
		pingConn, err = net.DialUDP("udp", nil, group)
		if err != nil {
			return err
		}
	}

	_, err = pingConn.Write(buf)
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
