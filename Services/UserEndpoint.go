package Services

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	Uid    int `json:"uid"`
	Method string
}

type UserResponse struct {
	Result string `json:"result"`
}

func GenUserEndpoint(userService IUserService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		r := request.(UserRequest)
		result := "nothing"
		//通過判斷執行方法
		if r.Method == "GET" {
			result = userService.GetName(r.Uid)
		} else if r.Method == "DELETE" {
			err := userService.DelName(r.Uid)
			if err != nil {
				//代表有錯無法傳入
				result = err.Error()
			} else {
				result = fmt.Sprintf("userid為%d的使用者刪除成功", r.Uid)
			}

		}

		//result := userService.GetName(r.Uid) 強制執行Get方法

		return UserResponse{Result: result}, nil
	}
}
