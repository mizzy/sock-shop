module "ecs_user" {
  source = "../ecs"

  service = {
    name               = "sock-shop-UserService-jgQWShzp465f"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
  }

  task = {
    name               = "user"
    image              = "weaveworksdemos/user"
    family             = "sock-shop-UserTask-RdJdCNzDqhe2"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    portMappings = [
      {
        containerPort = 80
        hostPort      = 80
      }
    ]
    environment = [
      {
        name  = "ZIPKIN"
        value = "http://zipkin:9411/api/v1/spans"
      },
    ]
  }

  registry = {
    name         = "user"
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
  }
}
