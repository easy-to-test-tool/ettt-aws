package ettt_aws

import "github.com/aws/aws-sdk-go-v2/config"

const (
	ExtensionName string = "aws"
)

type AwsContext struct {
	Config config.Config
}

func (aws AwsContext) ExtensionKey() string {
	return ExtensionName
}
