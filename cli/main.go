package main

import (
	"fmt"
	"os"
	"gospeedy/cli"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "gospeedy",
		Usage: "A CLI tool to generate code and handle project setup for your Go applications",
		Commands: []*cli.Command{
			// Command to create a new app
			{
				Name:  "new",
				Usage: "Create a new application",
				ArgsUsage: "<app-name>",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return fmt.Errorf("please provide an app name")
					}
					appName := c.Args().Get(0)
					cli.CreateNewApp(appName)
					return nil
				},
			},

			// Command to generate various components (modules, services, controllers, etc.)
			{
				Name:  "generate",
				Usage: "Generate various components for your app",
				Subcommands: []*cli.Command{
					{
						Name:  "module",
						Usage: "Generate a new module with controller and service",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:    "with-tests",
								Aliases: []string{"t"},
								Usage:   "Include a test file with the generated module",
							},
						},
						Action: func(c *cli.Context) error {
							if c.NArg() < 1 {
								return fmt.Errorf("please provide a module name")
							}
							moduleName := c.Args().Get(0)
							withTests := c.Bool("with-tests") // Check if the --with-tests flag was set
							cli.Generate("module", moduleName, withTests)
							return nil
						},
					},
					{
						Name:  "controller",
						Usage: "Generate a new controller",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:    "with-tests",
								Aliases: []string{"t"},
								Usage:   "Include a test file with the generated controller",
							},
						},
						Action: func(c *cli.Context) error {
							if c.NArg() < 1 {
								return fmt.Errorf("please provide a controller name")
							}
							controllerName := c.Args().Get(0)
							withTests := c.Bool("with-tests")
							cli.Generate("controller", controllerName, withTests)
							return nil
						},
					},
					{
						Name:  "service",
						Usage: "Generate a new service",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:    "with-tests",
								Aliases: []string{"t"},
								Usage:   "Include a test file with the generated service",
							},
						},
						Action: func(c *cli.Context) error {
							if c.NArg() < 1 {
								return fmt.Errorf("please provide a service name")
							}
							serviceName := c.Args().Get(0)
							withTests := c.Bool("with-tests")
							cli.Generate("service", serviceName, withTests)
							return nil
						},
					},
					{
						Name:  "resolver",
						Usage: "Generate a new GraphQL resolver",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:    "with-tests",
								Aliases: []string{"t"},
								Usage:   "Include a test file with the generated resolver",
							},
						},
						Action: func(c *cli.Context) error {
							if c.NArg() < 1 {
								return fmt.Errorf("please provide a resolver name")
							}
							resolverName := c.Args().Get(0)
							withTests := c.Bool("with-tests")
							cli.Generate("resolver", resolverName, withTests)
							return nil
						},
					},
					{
						Name:  "mutation",
						Usage: "Generate a new GraphQL mutation",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:    "with-tests",
								Aliases: []string{"t"},
								Usage:   "Include a test file with the generated mutation",
							},
						},
						Action: func(c *cli.Context) error {
							if c.NArg() < 1 {
								return fmt.Errorf("please provide a mutation name")
							}
							mutationName := c.Args().Get(0)
							withTests := c.Bool("with-tests")
							cli.Generate("mutation", mutationName, withTests)
							return nil
						},
					},
				},
			},
		},
	}

	// Run the CLI app
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error running the app:", err)
	}
}
