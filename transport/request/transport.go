package transport

import (
	"io/ioutil"
	"os"

	Schemes "../../internal/schemes"
	Utils "../../internal/utils"
	UserModule "../../modules/user"
)

var ServerAPI = &Schemes.Config.ServerAPI

func UserLogin() {
	Response := UserModule.UserLogin()
	if Response.Err != nil {
		Utils.Colorize(Utils.Printer{Color: -1, MesgErr: Response.Err})
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(Response.Res.Body)
	if err != nil {
		Utils.Colorize(Utils.Printer{Color: -1, MesgErr: err})
		os.Exit(1)
	}
	Utils.ManageResponse(Response, body)
	UserModule.StoreToken(string(body))
}
