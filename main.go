package main

import (
	"io"
	"net"
	"net/http"
	"os"
)

func check(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	IpAddress := getIPAddress(r)
	io.WriteString(w, "hostname: "+host)
	io.WriteString(w, ", RemoteAddr: "+IpAddress)
}

func getIPAddress(r *http.Request) string {
	address := r.Header.Get("X-Forwarded-For")

	if len(address) == 0 {
		address = r.Header.Get("X-Real-IP")
	}

	if len(address) == 0 {
		address, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	return address
}

func main() {
	http.HandleFunc("/", check)
	http.ListenAndServe(":8080", nil)
}
