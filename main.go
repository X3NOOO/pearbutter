package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/lrstanley/girc"
	"golang.org/x/sync/errgroup"
)

var flagConfig string
var flagVerbose bool

func init() {
	flag.StringVar(&flagConfig, "config", "", "Path to the configuration file")
	flag.BoolVar(&flagVerbose, "verbose", false, "Enable extra debug logs")
}

func main() {
	flag.Parse()

	ucdir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalln(err)
	}

	cpaths := []string{"./pearbutter.toml", "/etc/pearbutter/pearbutter.toml", ucdir + "/pearbutter.toml"}

	if flagConfig != "" {
		cpaths = []string{flagConfig}
	}

	c, err := GetConfig(cpaths)
	if err != nil {
		log.Fatalln(err)
	}

	logfile, err := os.OpenFile(c.Config.Logfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	var debugfile io.Writer
	if flagVerbose {
		debugfile = logfile
	} else {
		debugfile = io.Discard
	}

	log.SetOutput(logfile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var eg errgroup.Group

	for _, bc := range c.Servers {
		log.Println("Creating", bc.Name)

		bot := girc.New(girc.Config{
			Server:     bc.Address,
			ServerPass: bc.Password,
			Port:       bc.Port,
			SSL:        bc.Ssl,
			Nick:       bc.Nick,
			User:       bc.User,
			Name:       bc.Realname,
			Debug:      debugfile,
			Out:        logfile,
		})

		bc := bc
		eg.Go(func() error {
			return HandleBot(bot, &bc)
		})
	}

	err = eg.Wait()
	if err != nil {
		log.Fatalln(err)
	}
}
