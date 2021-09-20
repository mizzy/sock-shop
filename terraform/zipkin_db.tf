module "ecs_zipkin_db" {
  source = "./modules/ecs"

  service = {
    name               = "sock-shop-ZipkinDBService-WJxuKFFGBP8f"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
  }

  task = {
    name               = "zipkin-mysql"
    image              = "openzipkin/zipkin-mysql"
    family             = "sock-shop-ZipkinDBTask-25TrSccjpz8a"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    portMappings = [
      {
        containerPort = 3306,
        hostPort      = 3306,
        protocol      = "tcp",
      }
    ]
    memory = 1024
  }

  registry = {
    name         = "zipkin-mysql"
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
  }
}
