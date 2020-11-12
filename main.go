package main

import (
	"demo-gokit/Services"
	"demo-gokit/util"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	user := Services.UserService{}
	endp := Services.GenUserEndpoint(user)

	serverHandler := httptransport.NewServer(endp, Services.DecodeUserRequest, Services.EncodeUserResponse)

	r := mymux.NewRouter()
	{
		r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(serverHandler)
		r.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-type", "application/json")
			writer.Write([]byte(`{"status":"ok"}`))

		})
	}

	//使用協程的概念運行http Server
	errChan := make(chan error)
	go (func() {

		util.RegService()
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			log.Println(err)
			errChan <- err
		}

	})()

	go (func() {

		sig_c := make(chan os.Signal)
		signal.Notify(sig_c, syscall.SIGINT, syscall.SIGALRM)
		errChan <- fmt.Errorf("%s", <-sig_c)

	})()

	getErr := <-errChan
	util.Unregservice()
	log.Println(getErr)

}
