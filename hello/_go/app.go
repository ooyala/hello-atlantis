package main

import (
	atlantis "atlantis/types"
	"flag"
	"fmt"
	"log"
	"net/http"
	"hello"
)

var (
	listenAddr string
)

var (
	talk = flag.Bool("talk", false, "Run the talk test for id-verify job.")
)

func init() {
	// read atlantis config
	cfg, err := atlantis.LoadAppConfig()
	if err != nil {
		// use defaults
		log.Printf("Error opening log: %s, using defaults", err.Error())
		listenAddr = ":9876"
		return
	}
	listenAddr = fmt.Sprintf(":%d", cfg.HTTPPort)
}

func main() {
	flag.Parse()
	if *talk {
		log.Println(hello.Talk())
		return
	}

	log.Println("Now listening on", listenAddr)
	hello.InitDataDog()
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", hello.HealthzHandler)
	mux.HandleFunc("/statusz", hello.StatuszHandler)
	mux.HandleFunc("/art", hello.MsgHandler(hello.Gopher, false))
	mux.HandleFunc("/", hello.MsgHandler("Hello from Go", true))
	s := &http.Server{
		Addr: listenAddr,
		Handler: mux,
	}
	log.Fatalln(s.ListenAndServe().Error())
}
