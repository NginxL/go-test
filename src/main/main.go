package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/juju/errors"
	"github.com/siddontang/go-log/log"
)

// test toml
type tomlConfig struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

// test toml end

// test log
var logLevel = flag.String("log_level", "info", "log level")

// test log end

// test errors
func error_test(arg int) (int, error) {
	if arg < 0 {
		return -1, errors.Errorf("arg is not right:%d", arg)
	} else {
		return arg * arg, nil
	}
}

func error_test2(arg int) (int, error) {
	arg, err := error_test(arg)
	if err != nil {
		//return -1, errors.Trace(err)
		return -1, errors.Annotate(err, "error in test2")
	}
	return arg, nil
}

// test errors end

func main() {
	flag.Parse()
	var config tomlConfig
	if _, err := toml.DecodeFile("../conf/example.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Title: %s\n", config.Title)
	fmt.Printf("Owner: %s (%s, %s), Born: %s\n",
		config.Owner.Name, config.Owner.Org, config.Owner.Bio,
		config.Owner.DOB)
	fmt.Printf("Database: %s %v (Max conn. %d), Enabled? %v\n",
		config.DB.Server, config.DB.Ports, config.DB.ConnMax,
		config.DB.Enabled)
	for serverName, server := range config.Servers {
		fmt.Printf("Server: %s (%s, %s)\n", serverName, server.IP, server.DC)
	}
	fmt.Printf("Client data: %v\n", config.Clients.Data)
	fmt.Printf("Client hosts: %v\n", config.Clients.Hosts)

	fmt.Printf("log level:%s\n", *logLevel)
	log.SetLevelByName(*logLevel)
	log.Info("info")
	log.Infof("info with %s", "some info")
	log.Error("error")
	log.Errorf("error with %s", "some string")

	_, err := error_test2(-1)
	if err != nil {
		log.Error(errors.ErrorStack(err))
	}
}
