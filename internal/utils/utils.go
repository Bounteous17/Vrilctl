package utils

import (
	"log"
	"os"
	"os/user"
	"strconv"
	"sync"
	"time"

	Schemes "../schemes"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

// Logger : create log
type Logger struct {
	filename string
	*log.Logger
}

// Printer : info output parameters
type Printer struct {
	Color      int
	MesgStruct string
	MesgData   string
	MesgErr    error
	Log        bool
}

var logger *Logger
var once sync.Once

// GetLoggerInstance : create a loger instance object
func GetLoggerInstance() *Logger {
	Tracking := &Schemes.Config.Tracking
	once.Do(func() {
		logger = CreateLogger(Tracking.Path + Tracking.Cli)
	})
	return logger
}

// UserSysInfo : get user info
func UserSysInfo() {
	user, err := user.Current()
	if err != nil {
		Colorize(Printer{Color: -1, MesgErr: err})
		os.Exit(1)
	}

	Schemes.UserSys.Username = user.Username
	Schemes.UserSys.Uid = user.Uid
	Schemes.UserSys.HomeDir = user.HomeDir
}

// CreateLogger : instance log package
func CreateLogger(fname string) *Logger {
	t := time.Now()
	uid := os.Getuid()
	file, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		Colorize(Printer{Color: -1, MesgErr: err})
		os.Exit(1)
	}

	return &Logger{
		filename: fname,
		Logger:   log.New(file, t.Format("2006-01-02 15:04:05 | PID:")+strconv.Itoa(uid), 0),
	}
}

// Colorize : manage cli logs
func Colorize(params Printer) {
	descStatus := "[+]"
	switch params.Color {
	case -1:
		color.Red("%s", params.MesgErr)
		descStatus = "[-]"
	case 0:
		color.Cyan(params.MesgStruct, params.MesgData)
		descStatus = "[!]"
	case 1:
		color.Green(params.MesgStruct, params.MesgData)
	}

	if params.Log {
		if params.Color >= 0 {
			GetLoggerInstance().Printf(" "+descStatus+" "+params.MesgStruct+" ", params.MesgData)
		} else {
			GetLoggerInstance().Printf(" "+descStatus, " | ", params.MesgErr)
		}
	}
}

// ReadConf : read config toml
func ReadConf() {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")

	err := viper.ReadInConfig()
	if err != nil {
		Colorize(Printer{Color: -1, MesgErr: err})
		os.Exit(1)
	} else {
		Schemes.Config.ServerAPI.Server = viper.GetString("serverAPI.server")
		Schemes.Config.ServerAPI.Port = viper.GetInt("serverAPI.port")
		Schemes.Config.Tracking.Path = viper.GetString("tracking.path")
		Schemes.Config.Tracking.Cli = viper.GetString("tracking.cli")
	}
}
