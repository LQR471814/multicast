## multicast

***A high-level Golang package about multicasting. As everyone seems to know that working with multicast is a pain because of the inconsistencies between platforms and systems.***

### Basic Usage

```go
package main

import (
    "fmt"

    "github.com/LQR471814/multicast"
    operations "github.com/LQR471814/multicast/operations"
)

func main() {
    //? Check if multicast works
    multicastWorks, err := multicast.Check()
    if err != nil {
        log.Fatal(err)
    }

    if !multicastWorks {
        log.Fatal("Setup should be run first to continue")
        return
    }

    //? Get group addr struct
    group, err := net.ResolveUDPAddr("udp", "224.0.0.1:5001")
    if err != nil {
        log.Fatal(err)
    }

    //? Listen for packet
    err = multicast.Listen(group, func(packet operations.MulticastPacket) {
        log.Println(string(packet.Contents))
    })
    if err != nil {
        log.Fatal(err.Error())
    }

    //? Send to group
    err = multicast.Ping(group, []byte("Hello world!"))
    if err != nil {
        log.Fatal(err.Error())
        return
    }

    //? Wait
    log.Println("Waiting...")
    <-operations.Context().Done()
}
```

### Setup Api

```go
package setup

import (
    "log"

    "github.com/LQR471814/multicast"
)

func main() {
    err := multicast.Setup(23)

    if err != nil {
        log.Fatal(err)
    }
}
```

### Methods

#### `Check() (bool, error)`

Checks the operating system and the necessary setup steps required for multicasting to function on it.

The boolean returned will be false if setup is required.

**Returns**
a boolean which will be `true` if the system is ready for multicasting, if the system isn't supported or if it isn't ready for multicasting it will return `false`

#### `Setup(intf int) error`

**Parameter**
`intf` indicates the index of the network interface one should multicast on. This usually requires the user to manually choose.
However it can also be found by trying every interface to find which one(s) are connected.

Sets up multicasting on the system, on most systems this will require elevated privileges

#### `Ping(group *net.UDPAddr, buf []byte) error`

**Parameter**
`buf` is a byte slice to be broadcasted to the multicast group

Sends a message to the default multicast group

#### `Listen(group *net.UDPAddr, handler func(operations.MulticastPacket)) error`

**Parameter**
`handler` is called for every packet received from the default multicast group

An error will be returned if setup hasn't been completed on the given system.
The network interface it will listen on is the interface you've passed to the `Setup()` function

### Support

- Windows: Full
- MacOSX: None
- Linux: None
