package catalogue

import (
	e "github.com/mizzy/sock-shop/pulumi/lib/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var taskDef *ecs.TaskDefinition

func newTaskDefinition(ctx *pulumi.Context) error {
	var err error
	taskDef, err = ecs.NewTaskDefinition(ctx, "catalogue", &ecs.TaskDefinitionArgs{
		ContainerDefinitions: pulumi.Sprintf("[{\"command\":[\"/app\",\"-port=80\",\"-DSN=catalogue_user:default_password@tcp(%s:%d)/socksdb\"],\"cpu\":0,\"dnsSearchDomains\":[],\"dnsServers\":[],\"dockerLabels\":{},\"dockerSecurityOptions\":[],\"entryPoint\":[],\"environment\":[{\"name\":\"ZIPKIN\",\"value\":\"http://zipkin:9411/api/v1/spans\"}],\"environmentFiles\":[],\"essential\":true,\"extraHosts\":[],\"image\":\"weaveworksdemos/catalogue:0.3.5\",\"links\":[],\"logConfiguration\":{\"logDriver\":\"awslogs\",\"options\":{\"awslogs-group\":\"sock-shop\",\"awslogs-region\":\"ap-northeast-1\",\"awslogs-stream-prefix\":\"catalogue\"},\"secretOptions\":[]},\"mountPoints\":[],\"name\":\"catalogue\",\"portMappings\":[{\"containerPort\":80,\"hostPort\":80,\"protocol\":\"tcp\"}],\"secrets\":[],\"systemControls\":[],\"ulimits\":[],\"volumesFrom\":[]}]", dbInstance.Address, dbInstance.Port),
		Cpu:                  pulumi.String("256"),
		ExecutionRoleArn:     e.TaskExecutionRole.Arn,
		Family:               pulumi.String("sock-shop-CatalogueTask-DfaZsRAivWhG"),
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
