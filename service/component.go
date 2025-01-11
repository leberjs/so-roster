package service

import (
  "github.com/aws/aws-cdk-go/awscdk/v2"
  s3 "github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/jsii-runtime-go"
	"github.com/aws/constructs-go/constructs/v10"
  
  soapi "github.com/leberjs/so-roster/service/api"
)

type SORosterStackProps struct {
	awscdk.StackProps
}

func NewSORosterStack(scope constructs.Construct, id string, props *SORosterStackProps) awscdk.Stack {
  var sprops awscdk.StackProps
  if props != nil {
    sprops = props.StackProps
  }
  stack := awscdk.NewStack(scope, &id, &sprops)

  s3.NewBucket(stack, jsii.String("MetroAdultSOBkt"), &s3.BucketProps{
    BlockPublicAccess: s3.BlockPublicAccess_BLOCK_ALL(),
    BucketName: jsii.String("metro-adult-so"),
    RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
  })

  soapi.NewApiConstruct(stack, "SORosterAPI")

  return stack
}
