package orders

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var taskDef *ecs.TaskDefinition

func newTaskDefinition(ctx *pulumi.Context) error {
	/*
		var err error
		taskDef, err = ecs.NewTaskDefinition(ctx, "carts", &ecs.TaskDefinitionArgs{
			ContainerDefinitions: pulumi.String("[{\"command\":[],\"cpu\":0,\"dnsSearchDomains\":[],\"dnsServers\":[],\"dockerLabels\":{},\"dockerSecurityOptions\":[],\"entryPoint\":[],\"environment\":[],\"environmentFiles\":[],\"essential\":true,\"extraHosts\":[],\"image\":\"weaveworksdemos/carts:0.4.8\",\"links\":[],\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"sock-shop\",\"awslogs-region\":\"ap-northeast-1\",\"awslogs-stream-prefix\":\"carts\"},\"secretOptions\":[]},\"mountPoints\":[],\"name\":\"carts\",\"portMappings\":[{\"containerPort\":80,\"hostPort\":80,\"protocol\":\"tcp\"}],\"secrets\":[],\"systemControls\":[],\"ulimits\":[],\"volumesFrom\":[]}]"),
			Cpu:                  pulumi.String("256"),
			ExecutionRoleArn:     e.TaskExecutionRole.Arn,
			Family:               pulumi.String("sock-shop-CartsTask-eIq5v1xKpl13"),
			Memory:               pulumi.String("1024"),
			RequiresCompatibilities: pulumi.StringArray{
				pulumi.String("FARGATE"),
			},
			NetworkMode: pulumi.String("awsvpc"),
		})
		if err != nil {
			return err
		}

	*/

	return nil
}
