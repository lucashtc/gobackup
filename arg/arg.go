package arg

import (
	"fmt"
	"log"
	"os"

	"github.com/lucashtc/gobackup/mysql"

	"github.com/urfave/cli"
)

// Arg statement param, define option backup
func Arg() {
	conf := mysql.New()

	app := cli.NewApp()

	app.Name = "gobackup"
	app.Usage = "Script para fazer backup da base de dados local"
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "local, l",
			Usage:       "Set location stored backups,default /'backup/'",
			Destination: &conf.Dir,
		},
		cli.StringFlag{
			Name:        "user, u",
			Usage:       "User",
			Destination: &conf.User,
		},
		cli.StringFlag{
			Name:        "password, p",
			Usage:       "Password",
			Destination: &conf.Password,
		},
		cli.StringSliceFlag{
			Name:  "param, pa",
			Usage: "Parameters optional",
			//Destination: &conf.Param,
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.String("user") == "" {
			return fmt.Errorf("Flag user/-u can not be empty.%s", c.String("user"))
		}
		if c.StringSlice("param") != nil {
			conf.Param = c.StringSlice("param")
		}
		err := conf.Dump()
		if err != nil {
			return err
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
