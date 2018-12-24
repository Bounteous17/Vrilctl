package helpers

import (
	"flag"

	UsersMod "../../transport/user"
	Schemes "../schemes"
	Utils "../utils"
)

// AssignArgs : get cli args
func AssignArgs() {
	vbose := flag.Bool("v", false, "Specify username to log in api.")
	login := flag.String("login", "", "Define user to login")
	signup := flag.String("signup", "", "Create new user")
	flag.Parse()

	if *vbose {
		Schemes.Verbose.Verbose = true
		Schemes.Verbose.Level = 10
	}

	switch {
	case *login != "":
		Schemes.User.Username = *login
		UsersMod.Login()
	case *signup != "":
		Schemes.User.Username = *signup
	default:
		Utils.Colorize(Utils.Printer{Color: 0, MesgStruct: "Vrilctl %s", MesgData: "v0.0.0"})
	}
}
