package main

import (
	"demo-gokit/Services"
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"net/http"
)

func main() {

	user := Services.UserService{}
	endp := Services.GenUserEndpoint(user)

	serverHandler := httptransport.NewServer(endp, Services.DecodeUserRequest, Services.EncodeUserResponse)

	r := mymux.NewRouter()
	r.Handle("/user/{uid:\\d+}", serverHandler)
	//r.Handle(`/user/{uid:\d+}`,serverHandler)

	r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(serverHandler)

	http.ListenAndServe(":8080", r)

}
