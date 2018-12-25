package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	Schemes "../../internal/schemes"
)

var ServerAPI = &Schemes.Config.ServerAPI
var Response = Schemes.ResponseStruct{}

func UserLogin() Schemes.ResponseStruct {
	RequestURL := "http://" + ServerAPI.Server + ":" + strconv.Itoa(ServerAPI.Port)
	UserBody := Schemes.UserStruct{
		Username: Schemes.User.Username,
		Mail:     Schemes.User.Mail,
		Password: Schemes.User.Password,
	}
	UserEncode := new(bytes.Buffer)
	json.NewEncoder(UserEncode).Encode(UserBody)
	Response.Res, Response.Err = http.Post(RequestURL+"/login", "application/json; charset=utf-8", UserEncode)

	return Response
}
