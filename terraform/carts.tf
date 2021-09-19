resource "aws_ecs_service" "carts" {
  name            = "sock-shop-CartsService-IMy4x0jU4FX0"
  desired_count   = 1
  launch_type     = "FARGATE"
  cluster         = aws_ecs_cluster.sock_shop.id
  task_definition = module.ecs_carts.task_definition

  network_configuration {
    assign_public_ip = true
    security_groups  = [aws_security_group.ecs.id]
    subnets = [
      aws_subnet.public_subnet_1.id,
      aws_subnet.public_subnet_2.id,
    ]
  }

  service_registries {
    registry_arn = aws_service_discovery_service.carts.arn
  }
}

module "ecs_carts" {
  source = "./modules/ecs"
  task = {
    name               = "carts"
    image              = "weaveworksdemos/carts:0.4.8"
    family             = "sock-shop-CartsTask-eIq5v1xKpl13"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    port               = 80
    memory             = 1024
  }
}

resource "aws_service_discovery_service" "carts" {
  name = "carts"

  dns_config {
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
    dns_records {
      ttl  = 10
      type = "A"
    }
  }

  health_check_custom_config {
    failure_threshold = 1
  }
}
