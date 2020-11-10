package Services

import (
	"context"
	"encoding/json"
	"errors"
	mymux "github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//判斷參數怎麼來的
	vars := mymux.Vars(r)
	if uid, ok := vars["uid"]; ok {
		uid, _ := strconv.Atoi(uid)
		return UserRequest{
			Uid:    uid,
			Method: r.Method,
		}, nil
	}

	return nil, errors.New("參數錯誤")

}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	//將資料包裝為json格式
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)

}
