module "ecs_user_db" {
  source = "./modules/ecs"

  service = {
    name               = "sock-shop-UserDBService-bcNIEWFwyFiJ"
    cluster_id         = aws_ecs_cluster.sock_shop.id
    security_group_ids = [aws_security_group.ecs.id]
    subnet_ids         = [aws_subnet.public_subnet_1.id, aws_subnet.public_subnet_2.id]
    load_balancer      = null
  }

  task = {
    name               = "user-db"
    image              = "weaveworksdemos/user-db"
    family             = "sock-shop-UserDBTask-JH16IOxl33fR"
    execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
    port               = 27017
    environment        = []
    dockerLabels = {
      "agent.signalfx.com.port.27017" = "true",
    }
    task_role_arn = null
  }

  registry = {
    name         = "user-db"
    namespace_id = aws_service_discovery_private_dns_namespace.local.id
  }
}
