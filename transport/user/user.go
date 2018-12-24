package user

import (
	Schemes "../../internal/schemes"
	Utils "../../internal/utils"
)

// Login : obtain token if auth success
func Login() {
	Utils.Colorize(Utils.Printer{Color: 0, MesgStruct: "User trying to login -> %s", MesgData: Schemes.User.Username, Log: true})
}
