package Services

import "github.com/go-kit/kit/endpoint"

type UserRequest struct {

	Uid int `json:"uid"`
}

type UserResponse struct {

	Result string `json:"result"`

}

func GenUserEnpoint() endpoint.Endpoint {

}
