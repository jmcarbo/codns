// Codns is a DNS proxy server
package main

import (
	"flag"
	"github.com/miekg/dns"
	"github.com/jmcarbo/codns"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var (
	configFile = flag.String("config", "codns.json", "Configuration file")
	listenAddr = flag.String("listen", "localhost:53", "Local bind address")
)

func serve(net string, address string) {
	err := dns.ListenAndServe(address, net, nil)
	if err != nil {
	  log.Printf("Can't listen on %s, are you root? Trying listening on %s\n", address, "localhost:8053")
    err = dns.ListenAndServe("localhost:8053", net, nil)
    if err != nil {
  		log.Fatalf("Failed to setup the "+net+" server: %s\n", err.Error())
    }
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)
	flag.Parse()
	log.Printf("Listening on %s\n", *listenAddr)

	configuration := codns.ReadConfig(*configFile)
	for _, filter := range configuration.Filters {
		// Ensure that each pattern is a FQDN name
		pattern := dns.Fqdn(filter.Pattern)

		log.Printf("Proxing %s on %v\n", pattern, filter.Addresses)
		dns.HandleFunc(pattern, codns.ServerHandler(filter.Addresses))
	}

	go serve("tcp", *listenAddr)
	go serve("udp", *listenAddr)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
forever:
	for {
		select {
		case s := <-sig:
			log.Printf("Signal (%d) received, stopping\n", s)
			break forever
		}
	}
}
