package main

import (
	"io"
	"log"
	"os"

	"github.com/lrstanley/girc"
	"golang.org/x/sync/errgroup"
)

func main() {
	ucdir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	cpaths := []string{"./pearbutter.toml", "/etc/pearbutter/pearbutter.toml", ucdir + "/pearbutter.toml"}

	c, err := GetConfig(cpaths)
	if err != nil {
		log.Fatal(err)
	}

	// logfile, err := os.Open(c.Config.Logfile)
	// logfile, err := os.Open(c.Config.Logfile)
	// if err != nil {
	// log.Fatal(err)
	// }

	logfile := os.Stderr

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
			// Debug:      logfile,
			Debug:      io.Discard,
			Out:        os.Stderr,
		})

		eg.Go(func() error {
			return HandleBot(bot, &bc)
		})
	}
	
	err = eg.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
