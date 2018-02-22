package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"os/signal"
	"syscall"

	"github.com/xxlixin1993/CacheGo/logging"
	"github.com/xxlixin1993/CacheGo/configure"
)

const (
	KVersion = "0.0.1"
)

func main() {
	initFrame()
	waitSignal()
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

// Wait signal
func waitSignal() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan)

	sig := <-sigChan

	logging.Trace("signal: ", sig)

	switch sig {
	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
		// TODO Exit smoothly
		logging.Trace("exit...")
	case syscall.SIGUSR1:
		logging.Trace("catch the signal SIGUSR1")
	default:
		logging.Trace("signal do not know")
	}

	stop()
}

func stop(){
	logging.WaitLog()
}