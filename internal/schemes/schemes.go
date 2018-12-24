package schemes

// UserStruct : user information auth implicated
type UserStruct struct {
	Username string
	Mail     string
	Password string
	Token    string
}

// VerboseStruct : verbosity params
type VerboseStruct struct {
	Verbose bool
	Level   int
}

// ServerAPIStruct : connect to api
type ServerAPIStruct struct {
	Server string
	Port   int
}

// TrackingStruct : logging events from client actions
type TrackingStruct struct {
	Path string
	Cli  string
}

// ConfigStruct : global config sections
type ConfigStruct struct {
	ServerAPI ServerAPIStruct
	Tracking  TrackingStruct
}

// UserSysStruct : user info
type UserSysStruct struct {
	Username string
	Uid      string
	GroupIds []string
	HomeDir  string
}

var User = &UserStruct{}
var Verbose = &VerboseStruct{}
var ServerApi = &ServerAPIStruct{}
var Config = &ConfigStruct{}
var UserSys = &UserSysStruct{}
