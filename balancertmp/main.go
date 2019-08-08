package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
)

type balancerConfig struct {
	NetworkInterface string   `json:"interface"`
	Upstream         string   `json:"upstream"`
	HttpPath         string   `json:"path"`
	HttpMethods      []string `json:"methods"`
	Backends         []string `json:"backends"`
	ProxyMethod      string   `json:"proxyMethod"`
}

var (
	version        = "0.1"
	host           string
	configFilePath string
	config         balancerConfig
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "balancer"
	app.Version = version
	app.Usage = "simple balancer and reverse proxy for dar2019Internship"

	app.Commands = []cli.Command{
		{
			Name:      "run",
			Usage:     "balancer run",
			UsageText: "balancer run [--config-file|-c]",
			Action:    run,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config-file, c",
					Usage:       "path to config file",
					Destination: &configFilePath,
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {

	if c.NumFlags() < 1 {
		configFilePath = "config.json"
	}

	configFile, err := os.Open(configFilePath)

	if err != nil {
		log.Fatal(err)
	}

	defer configFile.Close()

	configParser(configFile)

	err = runBalancer()

	return err
}

func configParser(configFile *os.File) error {

	b, err := ioutil.ReadAll(configFile)
	if err != nil {
		return err
	}

	json.Unmarshal([]byte(b), &config)
	fmt.Println(config)

	return nil
}
