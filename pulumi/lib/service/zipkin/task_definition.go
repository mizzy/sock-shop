package zipkin

import (
	e "github.com/mizzy/sock-shop/pulumi/lib/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var taskDef *ecs.TaskDefinition

func newTaskDefinition(ctx *pulumi.Context) error {
	var err error
	taskDef, err = ecs.NewTaskDefinition(ctx, "zipkin", &ecs.TaskDefinitionArgs{
		ContainerDefinitions: pulumi.String("[{\"command\":[],\"cpu\":0,\"dnsSearchDomains\":[],\"dnsServers\":[],\"dockerLabels\":{},\"dockerSecurityOptions\":[],\"entryPoint\":[],\"environment\":[{\"name\":\"MYSQL_HOST\",\"value\":\"zipkin-mysql\"},{\"name\":\"MYSQL_PASS\",\"value\":\"zipkin\"},{\"name\":\"MYSQL_USER\",\"value\":\"zipkin\"},{\"name\":\"STORAGE_TYPE\",\"value\":\"mysql\"}],\"environmentFiles\":[],\"essential\":true,\"extraHosts\":[],\"image\":\"openzipkin/zipkin\",\"links\":[],\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"sock-shop\",\"awslogs-region\":\"ap-northeast-1\",\"awslogs-stream-prefix\":\"zipkin\"},\"secretOptions\":[]},\"mountPoints\":[],\"name\":\"zipkin\",\"portMappings\":[{\"containerPort\":9411,\"hostPort\":9411,\"protocol\":\"tcp\"}],\"secrets\":[],\"systemControls\":[],\"ulimits\":[],\"volumesFrom\":[]}]"),
		Cpu:                  pulumi.String("512"),
		ExecutionRoleArn:     e.TaskExecutionRole.Arn,
		Family:               pulumi.String("sock-shop-ZipkinTask-xR5HsG5Dz7Qr"),
		Memory:               pulumi.String("1024"),
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
