package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/335is/config"
	"github.com/335is/log"
	"github.com/335is/server/internal/router"
	uuid "github.com/satori/go.uuid"
)

type cfg struct {
	HTTP http
}

// HTTP holds web server configuration
//	SERVER_HTTP_ADDRESS
//	SERVER_HTTP_PORT
//	SERVER_HTTP_CONTENT
type http struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
	Content string `yaml:"content"`
}

const (
	// AppName defines the prefix for any configuration environment variables, as in SERVER_HTTP_ADDRESS
	appName    = "server"
	appVersion = "0.0.1"
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
	config.Load(appName, "", &c)

	go router.ServeHTTP(c.HTTP.Port, c.HTTP.Content)

	waitForExit()
	log.Infof("Stopping %s %s %s", appName, appVersion, appInstance)
	log.Infof("Shutting down")
}

func waitForExit() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGKILL, syscall.SIGTERM)
	sig := <-sigs
	log.Infof("Received signal %s, exiting...", sig)
}
