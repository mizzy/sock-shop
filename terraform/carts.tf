module "ecs_carts" {
  source = "./modules/ecs"

  service = {
    name               = "sock-shop-CartsService-IMy4x0jU4FX0"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
  }

  task = {
    name               = "carts"
    image              = "weaveworksdemos/carts:0.4.8"
    family             = "sock-shop-CartsTask-eIq5v1xKpl13"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    port               = 80
    memory             = 1024
    command            = []
    environment        = []
  }

  registry = {
    name         = "carts"
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
  }
}
