package main

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

type ipaddress struct {
	IPAddress string
}

func (ip ipaddress) IsSameAs(target ipaddress) bool {
	return ip.IPAddress == target.IPAddress
}

func intToBin(n int) string  {
	return strconv.FormatUint(*(*uint64)(unsafe.Pointer(&n)), 2)
}

func (ip ipaddress) GetMask() string {
	mask := ""
	parts := strings.Split(ip.IPAddress, ".")
	cc, err:= strconv.Atoi(parts[0])
	if ip.Valid() {
		if err == nil {
			if cc > 0 && cc < 224 {
				if (cc < 128) {
					mask = "255.0.0.0"
				}
				if (cc > 127 && cc < 192) {
					mask = "255.255.0.0"
				}
				if (cc > 191) {
					mask = "255.255.255.0"
				}
			}
		}
	}

	return mask
}

func (ip ipaddress) GetClass() string {
	parts := strings.Split(ip.IPAddress, ".")
	for _, part := range parts {
		a, _ := strconv.Atoi(part)
		if (a >= 1 && a <= 126) {
			return "Class A"
		} else if (a >= 128 && a <= 191) {
			return "Class B"
		} else if (a >= 192 && a < 223) {
			return "Class C"
		} else {
			return "Class D"
		}
	}
    return "INVALID"
}

func (ip ipaddress) Valid() (isvalid bool) {
	isvalid = true
	parts := strings.Split(ip.IPAddress, ".")
	if len(parts) !=4 {
		isvalid = false
	}
	for _, part := range parts {
		a, err := strconv.Atoi(part)
		if err!= nil {
			isvalid = false
		}
		if a < 1  || a > 255 {
			isvalid = false
		}
	}
	return
}

func main() {
	ip1 := ipaddress{"10.20.30.40"}
	fmt.Println(ip1)
	if ip1.Valid() {fmt.Println("ValidIP")} else {fmt.Println("InValidIP")}
	gc := ip1.GetClass() //return "classA" or "classB" or ...
	fmt.Println("IpClass is:", gc)
	ip2 := ipaddress{"10.20.30.50"}
	if ip1.IsSameAs(ip2) {fmt.Println("Same")} else {fmt.Println("Not same")}
	ipmask := ip1.GetMask()
	fmt.Println("IPMask is:", ipmask)
	fmt.Println(intToBin(192))
}
