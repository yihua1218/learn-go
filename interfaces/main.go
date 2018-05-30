package main

import (
	"fmt"
	"net"
)

func localAddresses() {
	ifaces, err := net.Interfaces()
	fmt.Print(ifaces)
	fmt.Print(err)
	addrs, err := net.InterfaceAddrs()
	fmt.Print(addrs)
	fmt.Print(err)
	if err != nil {
		fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
		return
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
			continue
		}
		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPAddr:
				fmt.Printf("%v : %s (%s)\n", i.Name, v, v.IP.DefaultMask())
			}

		}
	}
}

func main() {
	localAddresses()
}
