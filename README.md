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
    fmt.Println("Built")

    setup, err := multicast.Check()
    if err != nil {
        fmt.Println(err.Error())
    }

    if !setup {
        multicast.Setup(7)
    }

    if setup {
        multicast.Listen(func(packet operations.MulticastPacket) {
            fmt.Println(packet)
        })

        multicast.Ping([]byte("Hello world!"))
    }

    if setup {
        <-operations.Context().Done()
    }
}
```

### Methods

#### `Check() (bool, error)`

Checks the operating system and the necessary setup steps required for multicasting to function on it.

**Returns**
a boolean which will be `true` if the system is ready for multicasting, if the system isn't supported or if it isn't ready for multicasting it will return `false`

#### `Setup(intf int) error`

**Parameter**
`intf` indicates the index of the network interface one should multicast on. This usually requires the user to manually choose.
However it can also be found by trying every interface to find which one(s) are connected.

Sets up multicasting on the system, on most systems this will require elevated privileges

#### `Ping(buf []byte) error`

**Parameter**
`buf` is a byte slice to be broadcasted to the multicast group

Sends a message to the default multicast group

#### `Listen(handler func(operations.MulticastPacket)) error`

**Parameter**
`handler` is called for every packet received from the default multicast group

An error will be returned if setup hasn't been completed on the given system.
The network interface it will listen on is the interface you've passed to the `Setup()` function

### Support

- Windows: Full
- MacOSX: None
- Linux: None
