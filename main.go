package main

import (
	"demo-gokit/Services"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func main() {

	user := Services.UserService{}
	endp := Services.GenUserEndpoint(user)

	serverHandler := httptransport.NewServer(endp, Services.DecodeUserRequest, Services.EncodeUserResponse)

	http.ListenAndServe(":8080", serverHandler)

}
