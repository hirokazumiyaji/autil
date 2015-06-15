package main

import (
	"fmt"
	"log"
	"os"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/credentials"
	"github.com/awslabs/aws-sdk-go/service/ec2"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandConfig,
	commandEc2,
}

var commandConfig = cli.Command{
	Name:  "config",
	Usage: "",
	Description: `
`,
	Action: doConfig,
}

var commandEc2 = cli.Command{
	Name:  "ec2",
	Usage: "",
	Description: `
`,
	Action: doEc2,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doConfig(c *cli.Context) {
	profile := c.Args().Get(0)
	cred := credentials.NewSharedCredentials("", profile)
	value, _ := cred.Get()
	fmt.Println("Credentials")
	fmt.Println("  Profile: ", profile)
	fmt.Println("    AWS Access Key ID: ", value.AccessKeyID)
	fmt.Println("    AWS Secret Access Key: ", value.SecretAccessKey)
}

func doEc2(c *cli.Context) {
	region := c.Args().Get(0)
	profile := c.Args().Get(1)
	cred := credentials.NewSharedCredentials("", profile)

	svc := ec2.New(&aws.Config{Region: region, Credentials: cred})
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, res := range resp.Reservations {
		for _, instance := range res.Instances {
			var tagName string
			for _, tag := range instance.Tags {
				if *tag.Key == "Name" {
					tagName = *tag.Value
				}
			}
			fmt.Println(fmt.Sprintf("%s\t%s\t%s", tagName, *instance.PublicDNSName, *instance.State.Name))
		}
	}
}
