package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	Schemes "../../internal/schemes"
	Utils "../../internal/utils"
)

var ServerAPI = &Schemes.Config.ServerAPI
var Response = Schemes.ResponseStruct{}

func UserLogin() Schemes.ResponseStruct {
	RequestURL := "http://" + ServerAPI.Server + ":" + strconv.Itoa(ServerAPI.Port)
	UserBody := Schemes.User
	UserEncode := new(bytes.Buffer)
	json.NewEncoder(UserEncode).Encode(UserBody)
	Response.Res, Response.Err = http.Post(RequestURL+"/login", "application/json; charset=utf-8", UserEncode)

	return Response
}

func StoreToken(token string) {
	tokenRoute := Schemes.Auth.Path + Schemes.Auth.Token
	Utils.WriteFile(tokenRoute, token)
}
