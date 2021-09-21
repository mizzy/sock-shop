package ecs

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

func NewEcs() {
	pulumi.Run(dynamoDbTaskRole)
}
