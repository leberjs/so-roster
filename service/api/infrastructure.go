package api

import (
	// "github.com/aws/aws-cdk-go/awscdk/v2"
	apigwv2 "github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	integrations "github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewApiConstruct(scope constructs.Construct, id string) constructs.Construct {
	construct := constructs.NewConstruct(scope, &id)

	api := apigwv2.NewHttpApi(construct, jsii.String("SO"), &apigwv2.HttpApiProps{})

	// api.AddRoutes(&apigwv2.AddRoutesOptions{
	//   Path: jsii.String("/athlete"),
	//   Methods: &[]apigwv2.HttpMethod{
	//     apigwv2.HttpMethod_POST,
	//   },
	// })

	api.AddRoutes(&apigwv2.AddRoutesOptions{
		Path: jsii.String("/athletes"),
		Methods: &[]apigwv2.HttpMethod{
			apigwv2.HttpMethod_GET,
		},
		Integration: integrations.NewHttpUrlIntegration(
			jsii.String("GetAtheletes"),
			jsii.String("https://restful-booker.herokuapp.com/booking"),
			&integrations.HttpUrlIntegrationProps{},
		),
	})

	return construct
}
