package main

import (
    "fmt"
    "net/http"
        "net"
        "os"
)


func handler(w http.ResponseWriter, r *http.Request) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Fprintf(w, "%v\n", ipnet.IP.String())
			}
		}
	}
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8888", nil)
}
