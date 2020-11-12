package main

import (
	"demo-gokit/Services"
	"demo-gokit/util"
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"net/http"
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

	util.RegService()
	http.ListenAndServe(":8080", r)

}
