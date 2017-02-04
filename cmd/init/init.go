package init

import (
	"fmt"
	// "reflect"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/likemindnetworks/rbeta/terraform"
	"github.com/likemindnetworks/rbeta/aws"

	"github.com/aws/aws-sdk-go/aws/session"
)

type Init struct {
	namespace string
	dryrun bool
}

func (init *Init) run(ctx *kingpin.ParseContext) (err error) {
	// 0) check if terraform is present
	terraform.MustExists()

	// 1) setup s3 bucket for remote state, fails if already exists, or cannot create
	_, err = aws.ResolveCred()

	if (err != nil) {
		return err
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	creds, err := sess.Config.Credentials.Get()
	fmt.Printf("%s \n", creds)
	fmt.Printf("%s \n", *sess.Config.Region)

	// 2) create terraform config

	// 3) plan terraform

	// 4) execute the terraform

	return err
}

func ConfigCommand(rbeta *kingpin.Application, cmdName string) *kingpin.CmdClause {
	init := &Init{}

	cmd := rbeta.Command(cmdName, "init rbeta").Action(init.run);

	cmd.
		Flag(
			"namespace", "Name space for the rbeta deployment.",
		).Short('n').Required().StringVar(&init.namespace)

	cmd.
		Flag(
			"dryrun", "dryrun the init process.",
		).Short('d').BoolVar(&init.dryrun)

	return cmd
}
