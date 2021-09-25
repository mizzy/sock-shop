package queuemaster

import (
	e "github.com/mizzy/sock-shop/pulumi/lib/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var taskDef *ecs.TaskDefinition

func newTaskDefinition(ctx *pulumi.Context) error {
	var err error
	taskDef, err = ecs.NewTaskDefinition(ctx, "queue_master", &ecs.TaskDefinitionArgs{
		ContainerDefinitions: pulumi.String("[{\"command\":[],\"cpu\":0,\"dnsSearchDomains\":[],\"dnsServers\":[],\"dockerLabels\":{},\"dockerSecurityOptions\":[],\"entryPoint\":[],\"environment\":[],\"environmentFiles\":[],\"essential\":true,\"extraHosts\":[],\"image\":\"weaveworksdemos/queue-master\",\"links\":[],\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"sock-shop\",\"awslogs-region\":\"ap-northeast-1\",\"awslogs-stream-prefix\":\"queue-master\"},\"secretOptions\":[]},\"mountPoints\":[{\"containerPath\":\"/var/run/docker.sock\",\"sourceVolume\":\"VarRunDocker_Sock\"}],\"name\":\"queue-master\",\"portMappings\":[{\"containerPort\":80,\"hostPort\":80,\"protocol\":\"tcp\"}],\"secrets\":[],\"systemControls\":[],\"ulimits\":[],\"volumesFrom\":[]}]"),
		Cpu:                  pulumi.String("256"),
		ExecutionRoleArn:     e.TaskExecutionRole.Arn,
		Family:               pulumi.String("sock-shop-QueueMasterTask-WrLvLSJ3b4Sz"),
		Memory:               pulumi.String("1024"),
		RequiresCompatibilities: pulumi.StringArray{
			pulumi.String("FARGATE"),
		},
		NetworkMode: pulumi.String("awsvpc"),
		Volumes: ecs.TaskDefinitionVolumeArray{
			&ecs.TaskDefinitionVolumeArgs{
				Name: pulumi.String("VarRunDocker_Sock"),
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}
