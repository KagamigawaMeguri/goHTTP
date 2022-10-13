package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

var port string
var hostList []string

func init() {
	if len(os.Args) >= 2 {
		port = os.Args[1]
	} else {
		port = "8000"
	}

	portCheck, _ := strconv.Atoi(port)
	if portCheck < 1 || portCheck > 65535 {
		log.Fatal("Port error. \nUsage: goHTTP 8000")
	}
	getAdapter()
}

func ServerHTTP() {
	for _, host := range hostList {
		fmt.Printf("Serving HTTP on http://%s:%s/\n", host, port)
	}
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	err := http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	log.Fatal(err)
}
func getAdapter() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				hostList = append(hostList, ipnet.IP.String())
			}
		}
	}
}
func main() {
	ServerHTTP()
}
