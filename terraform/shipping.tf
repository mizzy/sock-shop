module "ecs_shipping" {
  source = "./modules/ecs"

  service = {
    name               = "sock-shop-ShippingService-kBVbMfmEkDP1"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
  }

  task = {
    name               = "shipping"
    image              = "weaveworksdemos/shipping"
    family             = "sock-shop-ShippingTask-yXjccEf6mzTa"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    portMappings = [
      {
        containerPort = 80,
        hostPort      = 80,
      }
    ]
    memory = 1024
  }

  registry = {
    name         = "shipping"
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
  }
}
