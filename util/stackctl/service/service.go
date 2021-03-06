package service

import (
	"sort"
	"strings"

	"github.com/stack-labs/stack"
	"github.com/stack-labs/stack/pkg/cli"
	"github.com/stack-labs/stack/service"
	"github.com/stack-labs/stack/util/log"
	"github.com/stack-labs/stack/util/stackctl/internal/util"
)

func client() service.Service {
	c := stack.NewService(stack.Name("stack.rpc.stackctl"))
	err := c.Init()
	if err != nil {
		log.Fatal("stackctl client init err: %s", err)
	}

	return c
}
func services(c *cli.Context, args []string) ([]byte, error) {
	cli := client()
	ss, err := cli.Client().Options().Registry.ListServices()
	if err != nil {
		log.Fatal("stackctl list services err: %s", err)
	}

	var services []string
	for _, service := range ss {
		services = append(services, service.Name)
	}

	sort.Strings(services)
	return []byte(strings.Join(services, "\n")), nil
}

func Commands() []cli.Command {
	return []cli.Command{
		{
			Name:   "services",
			Usage:  "List services in the registry",
			Action: util.Print(services),
		},
	}
}
