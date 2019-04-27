package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/domainr/whois"
	"github.com/genkiroid/cert"
	"github.com/urfave/cli"
)

var version string

func CmdConfigure() error {
	return nil
}

func CmdAddDomain(domain string) error {
	return nil
}

func CmdCheckDomain(domain string) error {
	request, _ := whois.NewRequest(domain)
	response, _ := whois.DefaultClient.Fetch(request)
	s := string(response.Body)
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Registry Expiry Date") {
			expiry_date := strings.TrimSpace(strings.Split(strings.TrimSpace(scanner.Text()), ": ")[1])
			fmt.Println(expiry_date)
		}
	}
	// Check SSL
	cert := cert.NewCert(domain)
	fmt.Println(cert.NotAfter)
	return nil
}

func main() {

	app := cli.NewApp()
	app.Version = version
	app.Name = "expiry"
	app.Usage = "Check domain/ssl expiration dates"
	app.Commands = []cli.Command{
		{
			Name:  "configure",
			Usage: "configure",
			Action: func(c *cli.Context) error {
				CmdConfigure()
				return nil
			},
		},
		{
			Name:  "add",
			Usage: "add [domain]",
			Action: func(c *cli.Context) error {
				CmdAddDomain(c.Args().First())
				return nil
			},
		},
		{
			Name:  "check",
			Usage: "check [domain]",
			Action: func(c *cli.Context) error {
				CmdCheckDomain(c.Args().First())
				return nil
			},
		},
	}
	app.Run(os.Args)
}
