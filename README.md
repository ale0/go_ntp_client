Cliente usando github.com/beevik/ntp para mostrar info del servidor NTP


## Instrucciones

```bash
mkdir go_ntp
cd go_ntp

go mod init ntp
go get github.com/beevik/ntp
```


## Compilar estatico (para copiar el binario a otro eqipo=
```bash
CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' ntp.go
```


```bash
~/test$ ./ntp pool.ntp.org

Time: 2023-03-31 13:20:18.157318544 -0300 -03 m=+0.235162367
ClockOffset: 72.079µs
RTT: 8.244917ms
Precision: 29ns
Stratum: 3
ReferenceID: 170658284 (10.44.9.236)
ReferenceTime: 2023-03-31 16:18:21.110967043 +0000 UTC
RootDelay: 142.868042ms
RootDispersion: 2.044678ms
Leap: 0
RootDistance: 77.601157ms
MinError: 0s
KissCode: 
Poll: 1s

```

