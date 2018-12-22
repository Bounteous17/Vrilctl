package app

import (
	"flag"
	"log"
	"os"
	"os/user"
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

// Verbose : verbosity params
type Verbose struct {
	Verbose bool
	Level   int
}

// ServerAPI : connect to api
type ServerAPI struct {
	server string
	port   int
}

// Tracking : logging events from client actions
type Tracking struct {
	path string
	cli  string
}

// Config : global config sections
type Config struct {
	serverAPI ServerAPI
	tracking  Tracking
}

// Logger : create log
type Logger struct {
	filename string
	*log.Logger
}

// UserSys : user info
type UserSys struct {
	username string
	uid      string
	homeDir  string
}

// Printer : info output parameters
type Printer struct {
	color      int
	mesgStruct string
	mesgData   string
	mesgErr    error
	log        bool
}

var verbose = Verbose{}
var logger *Logger
var once sync.Once

// GetLoggerInstance : create a loger instance object
func GetLoggerInstance() *Logger {
	once.Do(func() {
		logger = CreateLogger(ReadConf().tracking.path + ReadConf().tracking.cli)
	})
	return logger
}

// UserSysInfo : get user info
func UserSysInfo() UserSys {
	userSys := UserSys{}
	user, err := user.Current()
	if err != nil {
		Colorize(Printer{color: -1, mesgErr: err})
		os.Exit(1)
	}
	userSys.uid = user.Gid

	return userSys
}

// CreateLogger : instance log package
func CreateLogger(fname string) *Logger {
	t := time.Now()
	pid := os.Getpid()
	file, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		Colorize(Printer{color: -1, mesgErr: err})
		os.Exit(1)
	}

	return &Logger{
		filename: fname,
		Logger:   log.New(file, t.Format("2006-01-02 15:04:05 | PID:")+strconv.Itoa(pid), 0),
	}
}

// Colorize : manage cli logs
func Colorize(params Printer) {
	descStatus := "[+]"
	switch params.color {
	case -1:
		color.Red("%s", params.mesgErr)
		descStatus = "[-]"
	case 0:
		color.Cyan(params.mesgStruct, params.mesgData)
		descStatus = "[!]"
	case 1:
		color.Green(params.mesgStruct, params.mesgData)
	}

	if params.log {
		if params.color >= 0 {
			GetLoggerInstance().Printf(" "+descStatus+" "+params.mesgStruct+" ", params.mesgData)
		} else {
			GetLoggerInstance().Printf(" "+descStatus, " | ", params.mesgErr)
		}
	}
}

// ReadConf : read config toml
func ReadConf() Config {
	config := Config{}
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")

	err := viper.ReadInConfig()
	if err != nil {
		Colorize(Printer{color: -1, mesgErr: err})
		os.Exit(1)
	} else {
		config.serverAPI.server = viper.GetString("serverAPI.server")
		config.serverAPI.port = viper.GetInt("serverAPI.port")
		config.tracking.path = viper.GetString("tracking.path")
		config.tracking.cli = viper.GetString("tracking.cli")
	}

	return config
}

// AssignArgs : get cli args
func AssignArgs() {
	vbose := flag.Bool("v", false, "Specify username to log in api.")
	login := flag.String("login", "", "Define user to login")
	signup := flag.String("signup", "", "Create new user")
	flag.Parse()

	if *vbose {
		verbose.Verbose = true
		verbose.Level = 10
	}

	switch {
	case *login != "":
		Colorize(Printer{color: 1, mesgStruct: "Login user %s", mesgData: *login, log: true})
	case *signup != "":
		Colorize(Printer{color: 1, mesgStruct: "Signup user %s", mesgData: *signup, log: true})
	default:
		Colorize(Printer{color: 0, mesgStruct: "Vrilctl %s", mesgData: "v0.0.0"})
	}
}
