package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"github.com/335is/config"
	"github.com/335is/log"
	"github.com/335is/server/internal/router"
	uuid "github.com/satori/go.uuid"
)

// top level config structure
type cfg struct {
	HTTP   http      `yaml:"http"`
	CPU    profiling `yaml:"cpu"`
	Memory profiling `yaml:"memory"`
}

// HTTP holds web server configuration
//
//	HTTP_ADDRESS
//	HTTP_PORT
//	HTTP_CONTENT
type http struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
	Content string `yaml:"content"`
}

// CPU and Memory profiling settings
//
//	CPU_ENABLED
//	CPU_FILENAME
//	MEMORY_ENABLED
//	MEMORY_FILENAME
type profiling struct {
	Enabled  bool   `yaml:"enabled"`
	Filename string `yaml:"filename"`
}

const (
	appName    = "server"
	appVersion = "0.0.2"
)

var (
	appInstance string
)

func init() {
	appInstance = fmt.Sprintf("%s", uuid.NewV4())
}

func main() {
	log.Infof("Starting %s %s %s LOG_LEVEL=%s", appName, appVersion, appInstance, log.GetLevel().String())

	c := cfg{}
	config.Load("", &c)

	// CPU profiling
	if c.CPU.Enabled {
		f, err := os.Create(c.CPU.Filename)
		if err != nil {
			log.Errorf(err.Error())
		}

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// Memory profiling
	if c.Memory.Enabled {
		f, err := os.Create(c.Memory.Filename)
		if err != nil {
			log.Errorf(err.Error())
		}

		defer func() {
			pprof.WriteHeapProfile(f)
			f.Close()
		}()
	}

	go router.ServeHTTP(c.HTTP.Port, c.HTTP.Content)
	waitForExit()

	log.Infof("Stopping %s %s %s", appName, appVersion, appInstance)
}

func waitForExit() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGKILL, syscall.SIGTERM)
	sig := <-sigs
	log.Infof("Received signal %s, exiting...", sig)
}
