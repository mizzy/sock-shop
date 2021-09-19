module "ecs_carts_db" {
  source = "./modules/ecs"

  service = {
    name               = "sock-shop-CartsDBService-mPiXfiIT9Aky"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
  }

  task = {
    name               = "carts-db"
    image              = "mongo"
    family             = "sock-shop-CartsDBTask-6LKqO3WQLkdN"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    port               = 27017
    memory             = 512
    command            = []
    environment        = []
  }

  registry = {
    name         = "carts-db"
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
  }
}
