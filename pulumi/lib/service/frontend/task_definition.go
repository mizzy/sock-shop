package frontend

import (
	e "github.com/mizzy/sock-shop/pulumi/lib/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var taskDef *ecs.TaskDefinition

func newTaskDefinition(ctx *pulumi.Context) error {
	var err error
	taskDef, err = ecs.NewTaskDefinition(ctx, "frontend", &ecs.TaskDefinitionArgs{
		ContainerDefinitions: pulumi.String("[{\"command\":[],\"cpu\":0,\"dnsSearchDomains\":[],\"dnsServers\":[],\"dockerLabels\":{},\"dockerSecurityOptions\":[],\"entryPoint\":[],\"environment\":[{\"name\":\"SESSION_REDIS\",\"value\":\"true\"}],\"environmentFiles\":[],\"essential\":true,\"extraHosts\":[],\"image\":\"weaveworksdemos/front-end\",\"links\":[],\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"sock-shop\",\"awslogs-region\":\"ap-northeast-1\",\"awslogs-stream-prefix\":\"front-end\"},\"secretOptions\":[]},\"mountPoints\":[],\"name\":\"front-end\",\"portMappings\":[{\"containerPort\":8079,\"hostPort\":8079,\"protocol\":\"tcp\"}],\"secrets\":[],\"systemControls\":[],\"ulimits\":[],\"volumesFrom\":[]}]"),
		Cpu:                  pulumi.String("256"),
		ExecutionRoleArn:     e.TaskExecutionRole.Arn,
		Family:               pulumi.String("sock-shop-FrontEndTask-HW7tO3wK3fXn"),
		Memory:               pulumi.String("512"),
		RequiresCompatibilities: pulumi.StringArray{
			pulumi.String("FARGATE"),
		},
		NetworkMode: pulumi.String("awsvpc"),
	})
	if err != nil {
		return err
	}

	return nil
}
