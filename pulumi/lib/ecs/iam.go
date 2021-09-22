package ecs

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func dynamoDbTaskRole(ctx *pulumi.Context) error {
	role, err := iam.NewRole(ctx, "sock-shop-DynamoDbTaskRole", &iam.RoleArgs{
		AssumeRolePolicy:    pulumi.Any("{\"Version\":\"2008-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"ecs-tasks.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}"),
		ForceDetachPolicies: pulumi.Bool(false),
		MaxSessionDuration:  pulumi.Int(3600),
		Name:                pulumi.String("sock-shop-DynamoDbTaskRole-13YK50YRC8S9F"),
		Path:                pulumi.String("/"),
	})
	if err != nil {
		return err
	}

	_, err = iam.NewRolePolicyAttachment(ctx, "dynamodb_task_role-amazon_dynamodb_full_access", &iam.RolePolicyAttachmentArgs{
		PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess"),
		Role:      pulumi.Any(role.Name),
	})
	if err != nil {
		return err
	}

	return nil
}

func ecsTaskExecutionRole(ctx *pulumi.Context) error {
	role, err := iam.NewRole(ctx, "sock-shop-EcsTaskExecutionRole", &iam.RoleArgs{
		AssumeRolePolicy:    pulumi.Any("{\"Version\":\"2008-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"ecs-tasks.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}"),
		ForceDetachPolicies: pulumi.Bool(false),
		MaxSessionDuration:  pulumi.Int(3600),
		Name:                pulumi.String("sock-shop-EcsTaskExecutionRole-4VE06B3BDH02"),
		Path:                pulumi.String("/"),
	})
	if err != nil {
		return err
	}

	_, err = iam.NewRolePolicyAttachment(ctx, "ecs_task_execution_role-amazon_ecs_task_execution_role_policy", &iam.RolePolicyAttachmentArgs{
		PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"),
		Role:      pulumi.Any(role.Name),
	})
	if err != nil {
		return err
	}

	return nil
}
