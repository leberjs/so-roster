package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"github.com/leberjs/so-roster/service"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	service.NewSORosterStack(app, "SORosterStack", &service.SORosterStackProps{
		awscdk.StackProps{
			Env: nil,
		},
	})

	app.Synth(nil)
}
