package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/GoLive/configure"
	"github.com/GoLive/logging"
	"net"
)

const (
	KVersion = "0.0.1"
)

func main() {
	initFrame()
	// TODO 接受信号判断退出

	time.Sleep(time.Second * 2)

}

func initFrame() {
	// Parsing configuration environment
	runMode := flag.String("m", "local", "Use -m <config mode>")
	configFile := flag.String("c", "./conf/app.ini", "use -c <config file>")
	version := flag.Bool("v", false, "Use -v <current version>")
	flag.Parse()

	// Show version
	if *version {
		fmt.Println("Version", KVersion, runtime.GOOS+"/"+runtime.GOARCH)
		os.Exit(0)
	}

	// Initialize configure
	configErr := configure.InitConfig(*configFile, *runMode)
	if configErr != nil {
		fmt.Printf("Initialize Log error : %s", configErr)
		os.Exit(configure.KInitConfigError)
	}

	// Initialize Log
	LogErr := logging.InitLog()
	if LogErr != nil {
		fmt.Printf("Initialize Log error : %s", LogErr)
		os.Exit(configure.KInitLogError)
	}

	logging.Debug("test log")
}
