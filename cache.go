package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/xxlixin1993/CacheGo/configure"
	"github.com/xxlixin1993/CacheGo/logging"
	"github.com/xxlixin1993/CacheGo/server"
	"github.com/xxlixin1993/CacheGo/utils"
)

const (
	KVersion = "0.0.1"
)

func main() {
	initFrame()
	startServer()
	waitSignal()
	stop()
}

// Initialize framework
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

	// Initialize exitList
	utils.InitExitList()

	// Initialize configure
	configErr := configure.InitConfig(*configFile, *runMode)
	if configErr != nil {
		fmt.Printf("Initialize Configure error : %s", configErr)
		os.Exit(configure.KInitConfigError)
	}

	// Initialize log
	logErr := logging.InitLog()
	if logErr != nil {
		fmt.Printf("Initialize log error : %s", logErr)
		os.Exit(configure.KInitLogError)
	}

	logging.Trace("Initialized frame")
}

// Start server
func startServer() {
	serverErr := server.StartServer()
	if serverErr != nil {
		fmt.Printf("Start server error : %s", serverErr)
		os.Exit(configure.KInitSeverError)
	}
}

// Wait signal
func waitSignal() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan)

	sig := <-sigChan

	logging.TraceF("signal: %d", sig)

	switch sig {
	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
		// TODO Exit smoothly
		logging.Trace("exit...")
	case syscall.SIGUSR1:
		logging.Trace("catch the signal SIGUSR1")
	default:
		logging.Trace("signal do not know")
	}

}

// Stop the program
func stop() {
	utils.GetExitList().Stop()
}
