package operations

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/LQR471814/multicast/common"
)

var conn = &net.UDPConn{}
var group *net.UDPAddr

var ctx context.Context
var cancel context.CancelFunc

type MulticastPacket struct {
	Read     int
	Src      *net.UDPAddr
	Contents []byte
}

func init() {
	var err error

	group, err = net.ResolveUDPAddr("udp", common.MULTICAST_GROUP)
	if err != nil {
		log.Fatal(err)
	}

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
			if err != nil {
				cancel() //? Kind of unnecessary but it's here cause it looks nice anyway
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

func Listen(intf int, handler func(MulticastPacket)) error {
	var err error

	networkInterface, err := net.InterfaceByIndex(intf)
	if err != nil {
		return err
	}

	conn, err = net.ListenMulticastUDP("udp4", networkInterface, group)
	go listener(handler)

	return err
}

func Ping(buf []byte) error {
	_, err := conn.WriteToUDP(buf, group)
	if err != nil {
		return err
	}

	return nil
}
