package schemes

// UserStruct : user information auth implicated
type UserStruct struct {
	Username string
	Mail     string
	Password string
	Token    string
}

// Verbose : verbosity params
type VerboseStruct struct {
	Verbose bool
	Level   int
}

var User = &UserStruct{}
var Verbose = &VerboseStruct{}
