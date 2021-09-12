## multicast

***A high-level Golang package about multicasting. As everyone seems to know that working with multicast is a pain because of the inconsistencies between platforms and systems.***

### Usage

<details>
    <summary>Setup</summary>

    ```go
    package main

    import (
        "log"
        "flag"

        "github.com/LQR471814/multicast"
    )

    func main() {
        reset := flag.Bool("Reset", false, "Should multicasting setup be reset")
        exec := flag.String("Path", "", "The path of the executable that should be allowed to multicast")
        intf := flag.Int("Interface", 0, "Pass the interface index to use during setup")

        flag.Parse()

        if *reset {
            err := multicast.Reset()
            if err != nil {
                log.Fatal(err)
            }

            return
        }

        err := multicast.Setup(*exec, *intf)
        if err != nil {
            log.Fatal(err)
        }
    }
    ```
</details>

<details>
    <summary>Methods</summary>

    ```go
    package main

    import (
        "fmt"

        "github.com/LQR471814/multicast"
    )

    func main() {
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
</details>

### Methods

#### `Setup(intf int) error`

`intf` indicates the index of the network interface one should multicast on. This usually requires the user to manually choose.
However it can also be found by trying every interface to find which one(s) are connected.

Sets up multicasting on the system, on most systems this will require elevated privileges

#### `Ping(group *net.UDPAddr, buf []byte) error`

`buf` is a byte slice to be broadcasted to the multicast group

Sends a message to the default multicast group

#### `Listen(group *net.UDPAddr, handler func(operations.MulticastPacket)) error`

`handler` is called for every packet received from the default multicast group

An error will be returned if setup hasn't been completed on the given system.
The network interface it will listen on is the interface you've passed to the `Setup()` function

### Support

- Windows: Unstable
- MacOSX: Untested
- Linux: Unstable
