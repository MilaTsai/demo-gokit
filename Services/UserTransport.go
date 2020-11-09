package Services

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//判斷參數怎麼來的
	//http://localhost:xxx/?uid=101
	if r.URL.Query().Get("uid") != "" {
		uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))

		return UserRequest{
			Uid: uid,
		}, nil
	}
	return nil, errors.New("參數錯誤")

}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	//將資料包裝為json格式
	return json.NewEncoder(w).Encode(response)

}
