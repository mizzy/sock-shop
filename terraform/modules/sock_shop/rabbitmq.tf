module "ecs_rabbitmq" {
  source = "../ecs"

  service = {
    name               = "sock-shop-RabbitMQService-TQCQxjCK5zIf"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
  }

  task = {
    name               = "rabbitmq"
    image              = "rabbitmq:3.6.8-management"
    family             = "sock-shop-RabbitMQTask-Reu77SsyR0Bg"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    portMappings = [
      {
        containerPort = 5672
        hostPort      = 5672
        protocol      = "tcp"
      },
      {
        containerPort = 15672
        hostPort      = 15672
        protocol      = "tcp"
      },
    ]
    dockerLabels = {
      "agent.signalfx.com.port.15672" = "true"
    }
  }

  registry = {
    name         = "rabbitmq"
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
  }
}
