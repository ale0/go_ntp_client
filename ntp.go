package main

import (
    "encoding/binary"
    "net"
    "fmt"
    "os"
    "time"
    "github.com/beevik/ntp"
)


func main() {
    if len(os.Args) < 2 {
        fmt.Println("Por favor, proporciona la dirección del servidor NTP como argumento al ejecutar el programa.")
        os.Exit(1)
    }

    ntpServer := os.Args[1]

    response, err := ntp.Query(ntpServer)
    if err != nil {
        fmt.Println(err)
        return
    }

    refID := make([]byte, 4)
    binary.BigEndian.PutUint32(refID, response.ReferenceID)

    fmt.Printf("Time: %v\n", time.Now().Add(response.ClockOffset))
    fmt.Printf("ClockOffset: %v\n", response.ClockOffset)
    fmt.Printf("RTT: %v\n", response.RTT)
    fmt.Printf("Precision: %v\n", response.Precision)
    fmt.Printf("Stratum: %v\n", response.Stratum)
    fmt.Printf("ReferenceID: %v (%s)\n", response.ReferenceID, net.IP(refID).String())
    fmt.Printf("ReferenceTime: %v\n", response.ReferenceTime)
    fmt.Printf("RootDelay: %v\n", response.RootDelay)
    fmt.Printf("RootDispersion: %v\n", response.RootDispersion)
    fmt.Printf("Leap: %v\n", response.Leap)
    fmt.Printf("RootDistance: %v\n", response.RootDistance)
    fmt.Printf("MinError: %v\n", response.MinError)
    fmt.Printf("KissCode: %v\n", response.KissCode)
    fmt.Printf("Poll: %v\n", response.Poll)
    

}
