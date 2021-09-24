package orders

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/dynamodb"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var taskRole *iam.Role

func newDynamoDb(ctx *pulumi.Context) error {
	var err error
	taskRole, err = iam.NewRole(ctx, "dynamodb_task_role", &iam.RoleArgs{
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
		Role:      pulumi.Any(taskRole.Name),
	})
	if err != nil {
		return err
	}

	_, err = dynamodb.NewTable(ctx, "orders", &dynamodb.TableArgs{
		Attributes: dynamodb.TableAttributeArray{
			&dynamodb.TableAttributeArgs{
				Name: pulumi.String("id"),
				Type: pulumi.String("S"),
			},
			&dynamodb.TableAttributeArgs{
				Name: pulumi.String("customerId"),
				Type: pulumi.String("S"),
			},
		},
		HashKey:       pulumi.String("id"),
		Name:          pulumi.String("orders"),
		RangeKey:      pulumi.String("customerId"),
		ReadCapacity:  pulumi.Int(5),
		WriteCapacity: pulumi.Int(5),
	})
	if err != nil {
		return err
	}

	return nil
}
